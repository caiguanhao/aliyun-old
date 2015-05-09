package ecs

import (
	"fmt"
	"math"
	"time"
)

type DescribeInstances struct {
	Instances struct {
		Instance []DescribeInstanceAttribute `json:"Instance"`
	} `json:"Instances"`
	PageNumber int64  `json:"PageNumber"`
	PageSize   int64  `json:"PageSize"`
	RequestId  string `json:"RequestId"`
	TotalCount int64  `json:"TotalCount"`
}

func (instances DescribeInstances) Do(ecs *ECS) (*DescribeInstances, error) {
	return &instances, ecs.Request(map[string]string{
		"Action":   "DescribeInstances",
		"RegionId": "cn-hangzhou",
	}, &instances)
}

func (instances DescribeInstances) Print() {
	for _, instance := range instances.Instances.Instance {
		fmt.Println(instance.InstanceId)
	}
}

func (instances DescribeInstances) PrintTable() {
	idMaxLength := 2
	nameMaxLength := 4
	statusMaxLength := 6
	for _, instance := range instances.Instances.Instance {
		idLength := len(instance.InstanceId)
		nameLength := len(instance.InstanceName)
		statusLength := len(instance.Status)
		if idLength > idMaxLength {
			idMaxLength = idLength
		}
		if nameLength > nameMaxLength {
			nameMaxLength = nameLength
		}
		if statusLength > statusMaxLength {
			statusMaxLength = statusLength
		}
	}
	format := fmt.Sprintf(
		"%%-%ds  %%-%ds  %%-%ds  %%-15s  %%-15s  %%-15s  %%-35s\n",
		idMaxLength,
		nameMaxLength,
		statusMaxLength,
	)
	fmt.Printf(format, "ID", "Name", "Status", "Public IP", "Private IP", "Type", "Created At")
	for _, instance := range instances.Instances.Instance {
		createdAt, _ := time.Parse(time.RFC3339, instance.CreationTime)
		duration := time.Since(createdAt)
		createdAtStr := fmt.Sprintf("%s (%.0f days ago)",
			createdAt.Local().Format("2006-01-02 15:04:05"),
			math.Floor(duration.Hours()/24))
		fmt.Printf(
			format,
			instance.InstanceId,
			instance.InstanceName,
			instance.Status,
			instance.PublicIpAddress.GetIPAddress(0),
			instance.InnerIpAddress.GetIPAddress(0),
			instance.InstanceType,
			createdAtStr,
		)
	}
}
