package ecs

import (
	"errors"
	"fmt"

	"github.com/caiguanhao/aliyun/misc/opts"
)

type ModifyInstanceAttribute struct {
	RequestId string `json:"RequestId"`
}

func (modify ModifyInstanceAttribute) Do(ecs *ECS) (*ModifyInstanceAttribute, error) {
	id, err := opts.GetInstanceId()
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"Action":     "ModifyInstanceAttribute",
		"InstanceId": id,
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
