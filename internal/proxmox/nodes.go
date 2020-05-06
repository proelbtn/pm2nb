package proxmox

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/yaml.v2"
)

type Node struct {
	Name string `json:"node"`
	Status string `json:"status"`
	Cpu float64 `json:"cpu"`
	Mem int64 `json:"mem"`
	MaxCpu int64 `json:"maxcpu"`
	MaxMem int64 `json:"maxmem"`
	Level string `json:"level"`
	FingerPrint string `json:"ssl_fingerprint"`
	Uptime int64 `json:"uptime"`
}

type VirtualMachine struct {
	Id string `json:"vmid"`
	Pid string `json:"pid"`
	Name string `json:"name"`
	Status string `json:"status"`
	Uptime int64 `json:"uptime"`
	Cpu float64 `json:"cpu"`
	Mem int64 `json:"mem"`
	MaxCpu int64 `json:"cpus"`
	MaxMem int64 `json:"maxmem"`
	DiskRead int64 `json:"diskread"`
	DiskWrite int64 `json:"diskwrite"`
	NetIn int64 `json:"netin"`
	NetOut int64 `json:"netout"`
	Tags string `json:"tags"`
	Interfaces []NetworkInterface
}

const (
	StatusRunning = "running"
	StatusStopped = "stopped"
)


type NetworkInterface struct {
	Name string
	MACAddress string
	IPAddr4 string
	Gateway4 string
	IPAddr6 string
	Gateway6 string
}

type CloudInitConfig struct {
	Version int64 `yaml:"version"`
	Config []struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		MACAddress string `yaml:"mac_address"`
		Subnets []struct {
			Address string `yaml:"address"`
			NetMask string `yaml:"netmask"`
			Gateway string `yaml:"gateway"`
		} `yaml:"subnets"`
	} `yaml:"config"`
}

type GetNodesResponse struct {
	Data []Node `json:"data"`
}

type GetVirtualMachinesResponse struct {
	Data []VirtualMachine`json:"data"`
}

type DumpCloudInitConfigResponse struct {
	Data string `json:"data"`
}

func (c *Client) GetNodes(ctx context.Context) ([]Node, error) {
	endpoint := "/api2/json/nodes"

	req, err := c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	data := GetNodesResponse{}
	err = c.decodeBody(res, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (c *Client) GetVirtualMachines(ctx context.Context, node string) ([]VirtualMachine, error) {
	endpoint := fmt.Sprintf("/api2/json/nodes/%s/qemu", node)
	data := GetVirtualMachinesResponse{}

	req, err := c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = c.decodeBody(res, &data)
	if err != nil {
		return nil, err
	}

	for i := range data.Data {
		vms := data.Data

		config, err := c.DumpCloudInitConfig(ctx, node, vms[i].Id)
		if err != nil {
			continue
		}

		data := CloudInitConfig{}
		err = yaml.Unmarshal([]byte(config), &data)
		if err != nil {
			continue
		}

		for _, conf := range data.Config {
			if conf.Type != "physical" {
				continue
			}

			nic := NetworkInterface{}
			nic.Name = conf.Name
			nic.MACAddress = conf.MACAddress
			for _, subnet := range conf.Subnets {
				if !strings.Contains(subnet.Address, ":") {
					// TODO: 決め打ちやめろ
					nic.IPAddr4 = subnet.Address + "/24"
					nic.Gateway4 = subnet.Gateway
				} else {
					nic.IPAddr6 = subnet.Address
					nic.Gateway6 = subnet.Gateway
				}
			}
			vms[i].Interfaces = append(vms[i].Interfaces, nic)
		}
	}

	return data.Data, nil
}

func (c *Client) DumpCloudInitConfig(ctx context.Context, node string, vmid string) (string, error) {
	endpoint := fmt.Sprintf("/api2/json/nodes/%s/qemu/%s/cloudinit/dump", node, vmid)
	values := url.Values{}
	values.Add("type", "network")
	data := DumpCloudInitConfigResponse{}

	req, err := c.newRequest(ctx, http.MethodGet, endpoint, values)
	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	err = c.decodeBody(res, &data)
	if err != nil {
		return "", err
	}

	return data.Data, nil
}
