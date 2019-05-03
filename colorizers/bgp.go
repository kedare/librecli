package colorizers

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

// ColorizeBGPPeerAdminStatus will set a green or red color depending if the BGP session admin state is enabled or not
func ColorizeBGPPeerAdminStatus(status string) string {
	if status == "start" || status == "running" {
		return fmt.Sprint(Bold(Green(status)))
	} else {
		return fmt.Sprint(Bold(Red(status)))
	}
}

// ColorizeBGPPeerStatus will set a green or red color depending if the effective status of the BGP session is healthy or not
func ColorizeBGPPeerStatus(status string) string {
	if status == "established" {
		return fmt.Sprint(Bold(Green(status)))
	} else {
		return fmt.Sprint(Bold(Red(status)))
	}
}
