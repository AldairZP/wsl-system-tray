package sytemtray

import (
	"fmt"

	"github.com/aldairzp/wsl-system-tray/assets"
	"github.com/energye/systray"
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
	mStartWSL.Click(func() {
		systray.SetIcon(assets.WhileIcono)
		fmt.Println("started")
		systray.SetIcon(assets.OnIcono)
	})
	systray.AddSeparator()

	mStopWSL := systray.AddMenuItem("Stop WSL", "stop WSL")
	mStopWSL.Click(func() {
		systray.SetIcon(assets.WhileIcono)
		fmt.Println("ended")
		systray.SetIcon(assets.OffIcono)
	})
	mStopWSL.SetIcon(assets.OnIcono)

	systray.AddSeparator()

	mSalir := systray.AddMenuItemCheckbox("Exit", "Exit", true)
	mSalir.Click(func() {
		systray.Quit()
	})
}

func onExit() {
	// Limpieza de memoria (opcional aquí)
	fmt.Println("La app se ha cerrado.")
}
