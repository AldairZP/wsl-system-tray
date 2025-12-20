package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows"
)

var initStatus = false

//go:embed icons/off.ico
var offIcono []byte

//go:embed icons/on.ico
var onIcono []byte

//go:embed icons/while.ico
var whileIcono []byte

func startDocker() error {
	cmd := exec.Command("cmd", "/C", "docker", "desktop", "start")
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow: true,
	}
	_, err := cmd.Output()
	return err
}

func stopDocker() error {
	cmd := exec.Command("cmd", "/C", "docker", "desktop", "stop")
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow: true,
	}
	_, err := cmd.Output()
	return err
}

func startWSL() error {
	cmd := exec.Command("wsl", "true")
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow: true,
	}
	_, err := cmd.Output()
	return err
}

func stopWSL() error {
	cmd := exec.Command("wsl", "--shutdown")
	cmd.SysProcAttr = &windows.SysProcAttr{
		HideWindow: true,
	}
	_, err := cmd.Output()
	return err
}

func onReady() {
	systray.SetTitle("WSL - Docker")
	systray.SetTooltip("Start and Stop WSL - Docker Services")

	iconPath := "off.ico"

	if initStatus {
		iconPath = "on.ico"
	}

	iconData, err := os.ReadFile(iconPath)
	if err == nil {
		systray.SetIcon(iconData)
	}

	mStartWSL := systray.AddMenuItem("Start WSL", "start WSL")
	systray.AddSeparator()

	mStopWSL := systray.AddMenuItem("Stop WSL", "stop WSL")
	systray.AddSeparator()

	mSalir := systray.AddMenuItem("Exit", "Exit")

	go func() {
		for {
			select {
			case <-mStartWSL.ClickedCh:
				systray.SetIcon(whileIcono)
				startWSL()
				startDocker()
				fmt.Println("started")
				systray.SetIcon(onIcono)

			case <-mStopWSL.ClickedCh:
				systray.SetIcon(whileIcono)
				stopDocker()
				stopWSL()
				fmt.Println("ended")
				systray.SetIcon(offIcono)

			case <-mSalir.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Limpieza de memoria (opcional aquí)
	fmt.Println("La app se ha cerrado.")
}

func processIsActive(processName string) bool {
	procs, _ := process.Processes()
	for _, p := range procs {
		name, _ := p.Name()
		if name == processName {
			fmt.Println(name)
			return true
		}
	}
	fmt.Println("No encontrado")
	return false
}

func main() {

	WSLStatus := processIsActive("vmmemWSL")
	dockerStatus := processIsActive("Docker Desktop.exe")

	if WSLStatus && dockerStatus {
		initStatus = true
	} else {

		stopDocker()
		stopWSL()
	}

	systray.Run(onReady, onExit)
}
