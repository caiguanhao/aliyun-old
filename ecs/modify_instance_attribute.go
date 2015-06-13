package ecs

import (
	"errors"
	"flag"
	"fmt"

	"github.com/caiguanhao/aliyun/misc/opts"
)

type ModifyInstanceAttribute struct {
	RequestId string `json:"RequestId"`
}

func (modify ModifyInstanceAttribute) Do(ecs *ECS) (*ModifyInstanceAttribute, error) {
	params := map[string]string{
		"Action":     "ModifyInstanceAttribute",
		"InstanceId": flag.Arg(1),
	}
	if opts.Description != "\x00" {
		params["Description"] = opts.Description
	}
	if opts.InstanceName != "" {
		params["InstanceName"] = opts.InstanceName
	}
	if len(params) > 2 {
		return &modify, ecs.Request(params, &modify)
	}
	return nil, errors.New("Please provide at least one: --name, --description.")
}

func (modify ModifyInstanceAttribute) Print() {
	fmt.Println(modify.RequestId)
}

func (modify ModifyInstanceAttribute) PrintTable() {
	fmt.Println(modify.RequestId)
}
