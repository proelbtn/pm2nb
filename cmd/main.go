package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/proelbtn/proxmox2netbox/internal/proxmox"
	"os"
)

type Parameter struct {
	BaseUrl string
	Username string
	Password string
}

func getParams() (*Parameter, error) {
	baseUrl := os.Getenv("PM2NB_BASE")
	if baseUrl == "" {
		return nil, errors.New("missing parameter: PM2NB_BASE")
	}

	username := os.Getenv("PM2NB_USER")
	if username == "" {
		return nil, errors.New("missing parameter: PM2NB_BASE")
	}

	password := os.Getenv("PM2NB_PASS")
	if password == "" {
		return nil, errors.New("missing parameter: PM2NB_BASE")
	}

	return &Parameter{
		BaseUrl: baseUrl,
		Username: username,
		Password: password,
	}, nil

}

func main() {
	params, err := getParams()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	pm, err := proxmox.NewClient(params.BaseUrl, true)

	err = pm.Login(ctx, params.Username, params.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes, err := pm.GetNodes(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, node := range nodes {
		machines, err := pm.GetVirtualMachines(ctx, node.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, machine := range machines {
			fmt.Println(machine)
		}
	}
}

