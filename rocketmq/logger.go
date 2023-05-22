package rocketmq

import (
	"os"

	"github.com/rs/zerolog"
)

func LogBooter() zerolog.Logger {
	log := zerolog.New(os.Stdout)
	return log
}
