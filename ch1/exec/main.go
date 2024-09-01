package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"os"
	"strings"
)

var Version = "unknow"

var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "v", false, "Show Version")
}

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
	regex := regexp.MustCompile(`\s+`)
	for _, ele := range raw[1 : len(raw)-1] {
		ret := regex.Split(ele, -1)
		fmt.Printf("%#v\n", ret)

	}

}

func main() {
	flag.Parse()
	if showVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	// runCmd()
	runStdout()
}
