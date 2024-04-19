package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Command to execute
	cmd := "brew install clamav"

	// Execute the command
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cmd2 := "freshclam"

	// Execute the command
	out2, err2 := exec.Command("bash", "-c", cmd2).CombinedOutput()
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// Print the command output
	fmt.Println("Output:", string(out), string(out2))
}
