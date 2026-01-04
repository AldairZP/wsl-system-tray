package monitoring

import (
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/sys/windows"
)

var ProcessesNames []string
var ActivesProcesses = make(chan []bool)

func ProcessIsActive(processName string) bool {
	cmd := exec.Command(
		"powershell",
		"-Command",
		fmt.Sprintf(`Get-Process -Name "%s" -ErrorAction Stop`, processName),
	)
	cmd.SysProcAttr = &windows.SysProcAttr{HideWindow: true}

	err := cmd.Run()
	return err == nil
}

func ProcessesAreActive() []bool {
	var activesProcesses []bool
	for _, value := range ProcessesNames {
		processStatus := ProcessIsActive(value)
		activesProcesses = append(activesProcesses, processStatus)
	}
	return activesProcesses
}

func RunMonitoring() {
	go func() {
		for {
			ActivesProcesses <- ProcessesAreActive()
			time.Sleep(5 * time.Second)
		}
	}()
}
