package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/schigh/go-proofs/signals_in_goroutines/common"
)

func generateRunners() []common.Runner {
	runners := make([]common.Runner, 0, 10)
	var lastRunner *common.TaskRunner
	for i := 0; i < 10; i++ {
		lastRunner = &common.TaskRunner{
			Name: fmt.Sprintf("task%d", i+1),
			Delay: time.Duration(i+1) * time.Second,
		}
		runners = append(runners, lastRunner)
	}
	if lastRunner != nil {
		lastRunner.Err = errors.New("im the last runner and I have an error")
	}

	return runners
}

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	runners := generateRunners()
	if len(runners) == 0 {
		return
	}

	outerCtx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(outerCtx)

	for i := range runners {
		runner := runners[i]
		group.Go(func() error {
			return runner.Run(ctx)
		})
	}

	errChan := make(chan error)
	go func() {
		errChan <- group.Wait()
	}()

	select {
	case sig := <-sigChan:
		cancel()
		_, _ = fmt.Fprintf(os.Stderr, "caught signal: %s\n", sig)
	case err := <-errChan:
		_, _ = fmt.Fprintf(os.Stderr, "caught error: %v", err)
	}
	close(errChan)
}
