package main

import "os/exec"

func goVer() error {
	_, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	goVer()
}
