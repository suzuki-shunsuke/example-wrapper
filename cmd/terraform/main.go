package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	log.Println("[INFO] Terraform wrapper")
	cmd := exec.Command("aqua", append([]string{"exec", "--", "terraform"}, os.Args[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
