package monitoring

import (
	"os/exec"

	"golang.org/x/sys/windows"
)

func ProcessIsActive(processName string) bool {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-Process ", processName, " -ErrorAction Stop",
	)
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}

	err := cmd.Run()

	if err != nil {
		return false
	}
	return true
}
