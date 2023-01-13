package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var execCommand = exec.Command

func RunDocker(container string) ([]byte, error) {
	cmd := execCommand("docker", "run", "-d", container)
	return cmd.CombinedOutput()
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

const dockerRunResult = "foo!"

func TestRunDocker(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := RunDocker("docker/whalesay")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != dockerRunResult {
		t.Errorf("Expected %q, got %q", dockerRunResult, out)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Printf("foo!")
	os.Exit(0)
}
