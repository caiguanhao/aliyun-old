package ecs

import (
	"flag"
	"fmt"
)

type StopInstance struct {
	RequestId string `json:"RequestId"`
}

func (stop StopInstance) Do(ecs *ECS) (*StopInstance, error) {
	return &stop, ecs.Request(map[string]string{
		"Action":     "StopInstance",
		"InstanceId": flag.Arg(1),
	}, &stop)
}

func (stop StopInstance) Print() {
	fmt.Println(stop.RequestId)
}

func (stop StopInstance) PrintTable() {
	fmt.Println(stop.RequestId)
}
