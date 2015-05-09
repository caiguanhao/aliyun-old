package ecs

import (
	"flag"
	"fmt"
)

type AllocatePublicIP struct {
	RequestId string `json:"RequestId"`
	IpAddress string `json:"IpAddress"`
}

func (alloc AllocatePublicIP) Do(ecs *ECS) (*AllocatePublicIP, error) {
	return &alloc, ecs.Request(map[string]string{
		"Action":     "AllocatePublicIpAddress",
		"InstanceId": flag.Arg(1),
	}, &alloc)
}

func (alloc AllocatePublicIP) Print() {
	fmt.Println(alloc.IpAddress)
}

func (alloc AllocatePublicIP) PrintTable() {
	fmt.Println(alloc.IpAddress)
}
