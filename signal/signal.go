package signal

import (
	"fmt"
	"os/exec"
	"runtime"
)

func checkAndInstallBeep() error {
	_, err := exec.LookPath("beep")
	if err == nil {
		return nil
	}

	installCmd := exec.Command("sudo", "apt-get", "install", "-y", "beep")
	installCmd.Stdout = nil
	installCmd.Stderr = nil

	err = installCmd.Run()
	if err != nil {
		return fmt.Errorf("`beep` utilitasini o'rnatishda xatolik: %v", err)
	}

	return nil
}

func PlayErrorSound() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-c", "[console]::beep(1000,5000)")
	case "linux":
		err := checkAndInstallBeep()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd = exec.Command("beep", "-f", "1000", "-l", "5000")
	case "darwin":
		cmd = exec.Command("afplay", "/System/Library/Sounds/Funk.aiff")
	default:
		fmt.Println("Platforma qo'llab-quvvatlanmaydi")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Ovozli signal yuborishda xato:", err)
	}
}