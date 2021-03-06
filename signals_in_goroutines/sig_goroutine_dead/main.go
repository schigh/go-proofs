package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/schigh/go-proofs/signals_in_goroutines/common"
)

func generateRunners() []common.Runner {
	runners := make([]common.Runner, 0, 11)
	runners = append(runners, &common.SigRunner{})

	for i := 0; i < 10; i++ {
		runners = append(runners, &common.TaskRunner{
			Name: fmt.Sprintf("task%d", i+1),
			Delay: 1 * time.Second,
		})
	}

	return runners
}

func main() {
	runners := generateRunners()
	if len(runners) == 0 {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	for i := range runners {
		runner := runners[i]
		group.Go(func() error {
			err := runner.Run(ctx)
			// this errcheck and cancel arent necessary because group.Go will
			// cancel the context implicitly if any of its child goroutines
			// return an error
			if err != nil {
				_,_ = fmt.Fprintf(os.Stderr, "errgroup received error: %s\n", err.Error())
				cancel()
			}
			return err
		})
	}

	err := group.Wait()

	var closeErr common.ShutdownError
	switch {
	case errors.As(err, &closeErr):
		fmt.Println("closed from signal")
	case err != nil:
		fmt.Println("closing because of error:", err.Error())
	default:
		panic("this should never happen")
	}

}
