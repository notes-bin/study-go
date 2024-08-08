package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func runCmd() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func runStdout() {
	cmd := exec.Command("df", "-h")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	raw := strings.Split(outStr, "\n")
	// fmt.Printf("%#v\n", raw)
	regex := regexp.MustCompile(`\s+`)
	for _, ele := range raw[1 : len(raw)-1] {
		ret := regex.Split(ele, -1)
		fmt.Printf("%#v\n", ret)

	}

}

func main() {
	// runCmd()
	runStdout()
}
