package ecs

import (
	"flag"
	"fmt"
)

type StartInstance struct {
	RequestId string `json:"RequestId"`
}

func (start StartInstance) Do(ecs *ECS) (*StartInstance, error) {
	return &start, ecs.Request(map[string]string{
		"Action":     "StartInstance",
		"InstanceId": flag.Arg(1),
	}, &start)
}

func (start StartInstance) Print() {
	fmt.Println(start.RequestId)
}

func (start StartInstance) PrintTable() {
	fmt.Println(start.RequestId)
}
