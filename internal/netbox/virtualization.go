package netbox

import (
	"context"
	"fmt"
)

type VirtualMachine struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Status struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"status"`
	Cluster struct {
		Id int64 `json:"id"`
		Url string `json:"url"`
		Name string `json:"name"`
		VirtualMachineCount int64 `json:"virtualmachine_count"`
	} `json:"cluster"`
	Cpu int64 `json:"vcpus"`
	Memory int64 `json:"memory"`
	Disk int64 `json:"disk"`
}

type GetVirtualMachinesResponse struct {
	Count int64
	Next string
	Previous string
	Results []VirtualMachine
}

const (
	StatusOffline = "offline"
	StatusActive = "active"
	StatusPlanned = "planned"
	StatusStaged = "staged"
	StatusFailed = "failed"
	StatusDecommissioning = "decommissioning"
)

func (c *Client) DeleteVirtualMachine(ctx context.Context, id int64) error {
	endpoint := fmt.Sprintf("/virtualization/virtual-machines/%d", id)

	return c.DoDelete(ctx, endpoint)
}

func (c *Client) CreateVirtualMachine(ctx context.Context, name string, cluster int64) (*VirtualMachine, error) {
	endpoint := "/virtualization/virtual-machines/"

	data := VirtualMachine{}
	err := c.DoPost(ctx, endpoint, map[string]interface{} {
		"name": name,
		"cluster": cluster,
	}, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Client) GetVirtualMachines(ctx context.Context) ([]VirtualMachine, error) {
	endpoint := "/virtualization/virtual-machines/"

	data := GetVirtualMachinesResponse{}

	err := c.DoGet(ctx, endpoint, nil, &data)
	if err != nil {
		return nil, err
	}

	return data.Results, nil
}

func (c *Client) PatchVirtualMachine(ctx context.Context, machine VirtualMachine) (*VirtualMachine, error) {
	endpoint := fmt.Sprintf("/virtualization/virtual-machines/%d/", machine.Id)

	data := VirtualMachine{}
	err := c.DoPatch(ctx, endpoint, map[string]interface{} {
		"name": machine.Name,
		"cluster": machine.Cluster.Id,
		"vcpus": machine.Cpu,
		"memory": machine.Memory,
	}, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
