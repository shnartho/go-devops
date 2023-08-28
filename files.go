package main

import (
	"fmt"
	"os"
)

func main() {
	directories := []string{
		"cli_tools",
		"api_integration",
		"file_and_configuration_handling",
		"concurrency_and_parallelism",
		"error_handling",
		"testing_and_continuous_integration",
		"data_serialization_and_parsing",
		"regular_expressions",
		"containerization_and_orchestration",
		"infrastructure_as_code_iac",
		"interfacing_with_databases",
		"web_services",
		"data_transformation_and_processing",
		"logging_and_metrics",
		"security_practices",
	}

	for _, dir := range directories {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
		} else {
			fmt.Printf("Created directory: %s\n", dir)
		}
	}
}
