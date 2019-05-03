package colorizers

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

// ShouldBeHigherThan will set the text to green if the value is higher than a threshold or red if not
func ShouldBeHigherThan(value int, low int) string {
	if value > low {
		return fmt.Sprint(Green(value))
	} else {
		return fmt.Sprint(Red(value))
	}
}

// ShouldBeLowerThan will set the text to green if the value is lower than a threshold or red if not
func ShouldBeLowerThan(value int, high int) string {
	if value < high {
		return fmt.Sprint(Green(value))
	} else {
		return fmt.Sprint(Red(value))
	}
}
