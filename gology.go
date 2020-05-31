package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := "goodluck | goatsay | goldog"
	out := exec.Command("bash", "-c", cmd)
	out.Stdout = os.Stdout
	out.Stderr = os.Stderr
	err := out.Run()
	if err != nil {
		log.Fatalf("cmd.Run failed with %s", err)
	}
}
