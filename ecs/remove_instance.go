package ecs

import (
	"flag"
	"fmt"
)

type RemoveInstance struct {
	RequestId string `json:"RequestId"`
}

func (remove RemoveInstance) Do(ecs *ECS) (*RemoveInstance, error) {
	return &remove, ecs.Request(map[string]string{
		"Action":     "DeleteInstance",
		"InstanceId": flag.Arg(1),
	}, &remove)
}

func (remove RemoveInstance) Print() {
	fmt.Println(remove.RequestId)
}

func (remove RemoveInstance) PrintTable() {
	fmt.Println(remove.RequestId)
}
