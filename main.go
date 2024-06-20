package main

import (
	"os/exec"
)

func main() {
	exec.Command("reg", "add", "HKEY_CURRENT_USER\\Software\\Microsoft\\Command Processor", "/v", "AutoRun", "/d", "start notepad.exe", "/f").Run()
}
