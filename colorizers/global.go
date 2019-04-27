package colorizers

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

func ShouldBeHigherThan(value int, low int) string {
	if value > low {
		return fmt.Sprint(Green(value))
	} else {
		return fmt.Sprint(Red(value))
	}
}

func ShouldBeLowerThan(value int, high int) string {
	if value < high {
		return fmt.Sprint(Green(value))
	} else {
		return fmt.Sprint(Red(value))
	}
}
