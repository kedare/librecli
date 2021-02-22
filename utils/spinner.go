package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

const SPINNER_TYPE = 42
const SPINNER_DELAY = 100

func NewSpinner() *spinner.Spinner {
	newSpinner := spinner.New(spinner.CharSets[SPINNER_TYPE], SPINNER_DELAY*time.Millisecond)
	newSpinner.Suffix = "    Loading..."
	return newSpinner
}
