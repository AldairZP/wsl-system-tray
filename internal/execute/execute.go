package execute

import (
	"fmt"
	"os/exec"

	"golang.org/x/sys/windows"
)

func executeCommand(args ...string) error {
	args = append([]string{"-Command"}, args...)

	cmd := exec.Command("powershell.exe", args...)
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow: true,
	}

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("si se ejecuto")
	return err
}

func StartDocker() error {
	err := executeCommand("docker", "desktop", "start")
	return err
}

func StopDocker() error {
	err := executeCommand("docker", "desktop", "stop")
	return err
}

func StartWSL() error {
	err := executeCommand("wsl", "true")
	return err
}

func StopWSL() error {
	err := executeCommand("wsl", "--shutdown")
	return err
}
