package colorizers

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func ColorizeBGPPeerAdminStatus(status string) string {
	if status == "start" || status == "running" {
		return fmt.Sprint(Bold(Green(status)))
	} else {
		return fmt.Sprint(Bold(Red(status)))
	}
}

func ColorizeBGPPeerStatus(status string) string {
	if status == "established" {
		return fmt.Sprint(Bold(Green(status)))
	} else {
		return fmt.Sprint(Bold(Red(status)))
	}
}
