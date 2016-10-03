package main

import "os/exec"

func main() {
	for i := 0; i < 100; i++ {
		go dos()
	}
	for {
	}
}

func dos() {
	for {
		_ = exec.Command("ping", "-c1", "-t1", "10.0.19.153")
	}
}
