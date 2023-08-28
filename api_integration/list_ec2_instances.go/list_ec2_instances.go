/*
A simple CLI tool that can help you list and display the status of instances on your cloud provider.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Usage: devops-cli <comand> [options]")
		fmt.Println("Available commands: list")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
		listInstances()
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func listInstances() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error describing instances:", err)
		return
	}

	fmt.Println("Listing all instances")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, Status: %s \n", *instance.InstanceId, *instance.State.Name)
		}
	}
}
