package ecs

import (
	"flag"
	"fmt"
)

type RestartInstance struct {
	RequestId string `json:"RequestId"`
}

func (restart RestartInstance) Do(ecs *ECS) (*RestartInstance, error) {
	return &restart, ecs.Request(map[string]string{
		"Action":     "RebootInstance",
		"InstanceId": flag.Arg(1),
	}, &restart)
}

func (restart RestartInstance) Print() {
	fmt.Println(restart.RequestId)
}

func (restart RestartInstance) PrintTable() {
	fmt.Println(restart.RequestId)
}
