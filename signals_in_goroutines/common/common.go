package common

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Runner interface {
	Run(context.Context) error
}

type ShutdownError struct{}

func (err ShutdownError) Error() string {
	return "shutting down from signal"
}

type SigRunner struct {
	sigChan chan os.Signal
}

func (runner *SigRunner) Run(ctx context.Context) error {
	runner.sigChan = make(chan os.Signal)
	signal.Notify(runner.sigChan, syscall.SIGTERM, syscall.SIGINT)

	_, _ = fmt.Fprintln(os.Stderr, "SigRunner awaiting signal")

	select {
	case sig := <-runner.sigChan:
		_, _ = fmt.Fprintf(os.Stderr, "SigRunner caught signal '%s', returning shutdown error\n", sig)
		return ShutdownError{}
	case <-ctx.Done():
		_, _ = fmt.Fprintln(os.Stderr, "stopping runner because context canceled")
		return ctx.Err()
	}
}

type TaskRunner struct {
	Name  string
	Delay time.Duration
	Err   error
}

func (runner *TaskRunner) Run(ctx context.Context) error {
	_, _ = fmt.Fprintf(os.Stderr, "%s will run for %s. error: %t\n", runner.Name, runner.Delay, runner.Err != nil)
	if runner.Delay > 0 {
		<-time.After(runner.Delay)
	}

	_, _ = fmt.Fprintf(os.Stderr, "%s completing after %s.  error: %v\n", runner.Name, runner.Delay, runner.Err)

	return runner.Err
}

