package main

import (
	"fmt"
	"os/exec"
	"log"
)

func runTerraformApply() error {
	cmd := exec.Command("terraform", "apply", "-auto-approve")
	cmd.Stdout = fmt.Println
	cmd.Stderr = fmt.Println
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to apply Terraform configuration: %v", err)
	}
	return nil
}

func main() {
	// Trigger the Terraform apply command
	err := runTerraformApply()
	if err != nil {
		log.Fatalf("Error running Terraform: %v", err)
	}
	fmt.Println("Terraform apply was successful!")
}
