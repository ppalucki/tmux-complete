package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {

	target := flag.String("t", "", "target-panel")
	flag.Parse()

	fmt.Printf("target = %+v\n", *target)

	// pipe-pane
	args := []string{"capture-pane"}
	if *target != "" {
		args = append(args, "-t")
		args = append(args, *target)
	}
	fmt.Printf("args = %+v\n", args)
	cmd := exec.Command("tmux", args...)
	fmt.Printf("cmd = %#v\n", cmd)
}
