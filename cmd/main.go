package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/proelbtn/proxmox2netbox/internal/netbox"
	"github.com/proelbtn/proxmox2netbox/internal/proxmox"
	"os"
)

type Parameter struct {
	Proxmox struct {
		BaseUrl  string
		Username string
		Password string
	}
	NetBox struct {
		BaseUrl string
		Token string
	}
}

func getParams() (*Parameter, error) {
	params := Parameter{}

	params.Proxmox.BaseUrl = os.Getenv("PM2NB_PM_BASE")
	if params.Proxmox.BaseUrl == "" {
		return nil, errors.New("missing parameter: PM2NB_PM_BASE")
	}

	params.Proxmox.Username = os.Getenv("PM2NB_PM_USER")
	if params.Proxmox.Username == "" {
		return nil, errors.New("missing parameter: PM2NB_PM_USER")
	}

	params.Proxmox.Password = os.Getenv("PM2NB_PM_PASS")
	if params.Proxmox.Password == "" {
		return nil, errors.New("missing parameter: PM2NB_PM_PASS")
	}

	params.NetBox.BaseUrl = os.Getenv("PM2NB_NB_BASE")
	if params.NetBox.BaseUrl == "" {
		return nil, errors.New("missing parameter: PM2NB_NB_BASE")
	}

	params.NetBox.Token = os.Getenv("PM2NB_NB_TOKEN")
	if params.NetBox.Token == "" {
		return nil, errors.New("missing parameter: PM2NB_NB_TOKEN")
	}

	return &params, nil
}

func initClients(params *Parameter) (*proxmox.Client, *netbox.Client, error) {
	ctx := context.Background()
	pm, err := proxmox.NewClient(params.Proxmox.BaseUrl, true)
	if err != nil {
		return nil, nil, err
	}

	err = pm.Login(ctx, params.Proxmox.Username, params.Proxmox.Password)
	if err != nil {
		return nil, nil, err
	}

	nb, err := netbox.NewClient(params.NetBox.BaseUrl, params.NetBox.Token, false)
	if err != nil {
		return nil, nil, err
	}

	return pm, nb, nil
}

func die(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func fetchVirtualMachinesFromProxmox(pm *proxmox.Client) ([]proxmox.VirtualMachine, error){
	ctx := context.Background()

	nodes, err := pm.GetNodes(ctx)
	if err != nil {
		return nil, err
	}

	machines := make([]proxmox.VirtualMachine, 0)
	for _, node := range nodes {
		ms, err := pm.GetVirtualMachines(ctx, node.Name)
		if err != nil {
			return nil, err
		}

		machines = append(machines, ms...)
	}

	return machines, nil
}

func fetchVirtualMachinesFromNetBox(nb *netbox.Client) ([]netbox.VirtualMachine, error){
	ctx := context.TODO()
	return nb.GetVirtualMachines(ctx)
}

func syncVirtualMachines(pm *proxmox.Client, nb *netbox.Client, pvms []proxmox.VirtualMachine, nvms []netbox.VirtualMachine) error {
	ctx := context.TODO()

	name2pvm := make(map[string]proxmox.VirtualMachine)
	for _, pvm := range pvms {
		name2pvm[pvm.Name] = pvm
	}

	name2nvm := make(map[string]netbox.VirtualMachine)
	for _, nvm := range nvms {
		name2nvm[nvm.Name] = nvm
	}

	// 1. Create Virtual Machine found in Proxmox, but not found in NetBox
	for _, pvm := range pvms {
		if _, ok := name2nvm[pvm.Name]; !ok {
			vm, err := nb.CreateVirtualMachine(ctx, pvm.Name, 1)
			if err != nil {
				fmt.Printf("error occured while creating virtual machine: %s\n", err)
				continue
			}
			fmt.Printf("created virtual machine: %v\n", vm)
			name2nvm[pvm.Name] = *vm
		}
	}

	// 2. Delete Virtual Machine found in NetBox, but not found in Proxmox
	for _, nvm := range nvms {
		if _, ok := name2pvm[nvm.Name]; !ok {
			err := nb.DeleteVirtualMachine(ctx, nvm.Id)
			if err != nil {
				fmt.Printf("error occured while deleting virtual machine: %s\n", err)
				continue
			}
			fmt.Printf("deleted virtual machine: %v\n", nvm)
			delete(name2nvm, nvm.Name)
		}
	}

	// 3. Sync parameters
	for _, nvm := range name2nvm {
		pvm := name2pvm[nvm.Name]
		nvm.Cpu = pvm.MaxCpu
		nvm.Memory = pvm.MaxMem / 1024 / 1024
		nvm.Status.Value = netbox.StatusOffline
		if pvm.Status == proxmox.StatusRunning {
			nvm.Status.Value = netbox.StatusActive
		}

		vm, err := nb.PatchVirtualMachine(ctx, nvm)
		if err != nil {
			fmt.Printf("error occured while updating virtual macihne: %s\n", err)
			continue
		}
		fmt.Printf("updated virtual machine: %v\n", vm)
	}

	return nil
}

func main() {
	params, err := getParams()
	if err != nil {
		die(err)
	}

	pm, nb, err := initClients(params)

	pvms, err := fetchVirtualMachinesFromProxmox(pm)
	if err != nil {
		die(err)
	}

	nvms, err := fetchVirtualMachinesFromNetBox(nb)
	if err != nil {
		die(err)
	}

	syncVirtualMachines(pm, nb, pvms, nvms)
}
