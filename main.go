package main

import (
	"github.com/kedare/librecli/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

var version string

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	cmd.Setup(version)
}
