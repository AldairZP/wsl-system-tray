package sytemtray

import (
	"fmt"

	"github.com/aldairzp/wsl-system-tray/assets"
	"github.com/getlantern/systray"
)

func SystemTrayRun() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("WSL - Docker")
	systray.SetTooltip("Start and Stop WSL - Docker Services")

	initIcon := assets.OffIcono

	systray.SetIcon(initIcon)

	mStartWSL := systray.AddMenuItem("Start WSL", "start WSL")
	systray.AddSeparator()

	mStopWSL := systray.AddMenuItem("Stop WSL", "stop WSL")
	systray.AddSeparator()

	mSalir := systray.AddMenuItem("Exit", "Exit")

	go func() {
		for {
			select {
			case <-mStartWSL.ClickedCh:
				systray.SetIcon(assets.WhileIcono)
				fmt.Println("started")
				systray.SetIcon(assets.OnIcono)

			case <-mStopWSL.ClickedCh:
				systray.SetIcon(assets.WhileIcono)
				fmt.Println("ended")
				systray.SetIcon(assets.OffIcono)

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
