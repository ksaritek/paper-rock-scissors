package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func AppContext() (context.Context, context.CancelFunc) {

	return signal.NotifyContext(context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
		os.Interrupt,
		os.Kill,
	)
}
