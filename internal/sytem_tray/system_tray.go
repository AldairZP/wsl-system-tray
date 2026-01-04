package sytemtray

import (
	"fmt"
	"time"

	"github.com/aldairzp/wsl-system-tray/assets"
	"github.com/aldairzp/wsl-system-tray/internal/execute"
	"github.com/aldairzp/wsl-system-tray/internal/monitoring"
	"github.com/energye/systray"
)

func SystemTrayRun() {
	systray.Run(onReady, onExit)
}

var activesProcesses = []bool{false, false}
var isInProcess = make(chan bool)

func onReady() {
	systray.SetTitle("WSL - Docker")
	systray.SetTooltip("Start and Stop WSL - Docker Services")
	systray.SetIcon(assets.WhileIcono)

	mToggleWSL := systray.AddMenuItem("WSL", "Toggle WSL")
	mToggleWSL.Click(func() {
		systray.SetIcon(assets.WhileIcono)
		mToggleWSL.SetIcon(assets.CircleWhile)
		select {
		case ch := <-isInProcess:
			if ch {
				return
			}
		default:

		}
		go func() {
			isInProcess <- true
			execute.ToggleWSL(activesProcesses[0])
			isInProcess <- false
		}()
	})
	systray.AddSeparator()

	mToggleDocker := systray.AddMenuItem("Docker", "Toggle Docker")
	mToggleDocker.Click(func() {
		systray.SetIcon(assets.WhileIcono)
		mToggleDocker.SetIcon(assets.CircleWhile)
		select {
		case ch := <-isInProcess:
			if ch {
				return
			}
		default:

		}
		go func() {
			isInProcess <- true
			execute.ToggleDocker(activesProcesses[1])
			isInProcess <- false

		}()
	})
	systray.AddSeparator()

	mSalir := systray.AddMenuItem("Exit", "Exit")
	mSalir.Click(func() {
		systray.Quit()
	})

	go func() {
		monitoring.ProcessesNames = append(monitoring.ProcessesNames, "VmmemWSL")
		monitoring.ProcessesNames = append(monitoring.ProcessesNames, "Docker Desktop")
		monitoring.RunMonitoring()
		for {
			activesProcesses = <-monitoring.ActivesProcesses

			select {
			case ch := <-isInProcess:
				if ch {
					time.Sleep(5 * time.Second)
					continue
				}
			default:

			}

			if activesProcesses[0] {
				systray.SetIcon(assets.OnIcono)
				mToggleWSL.SetIcon(assets.CircleRunning)
			} else {
				systray.SetIcon(assets.OffIcono)
				mToggleWSL.SetIcon(assets.CircleStop)
			}
			if activesProcesses[1] {
				mToggleDocker.SetIcon(assets.CircleRunning)
			} else {
				mToggleDocker.SetIcon(assets.CircleStop)
			}

			time.Sleep(5 * time.Second)
		}

	}()

}

func onExit() {
	fmt.Println("end")
}
