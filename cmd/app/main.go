package main

import (
	sytemtray "github.com/aldairzp/wsl-system-tray/internal/sytem_tray"
)

func main() {

	// monitoring.ProcessesNames = append(monitoring.ProcessesNames, "VmmemWSL")
	// monitoring.RunMonitoring()
	// for {

	// 	time.Sleep(5 * time.Second)
	// 	for index := range monitoring.ActivesProcesses {

	// 		fmt.Println(index)
	// 	}

	// }

	sytemtray.SystemTrayRun()

}
