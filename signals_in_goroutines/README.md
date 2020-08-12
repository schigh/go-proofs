# signals in goroutines

This is a demonstration of why exit signals (`SIGINT`, `SIGTERM`, etc) should only
be trapped from inside the main goroutine.

## main

Catching term signals in the main goroutine

Run `./run.sh main`.  You will see tasks spin up in goroutines.  If 10 seconds 
elapses, the last goroutine returns an error and terminates the program.  If you 
pass `SIGINT` or `SIGTERM` before that happens, the main goroutine exits, causing 
the scheduler to stop any child goroutines.  The application exits cleanly as you 
would expect.

## dead

Catching term signals in the only active child goroutine

Run `./run.sh dead`.  You will see tasks spin up in goroutines.  Additionally, a 
signal task is dispatched.  This signal task listens for SIGINT and `SIGTERM` and 
returns an error to the caller (errgroup) when the signal is caught.  The task 
goroutines will finish their work in one second, leaving only the signal task 
blocking the errgroup.  Sending `SIGINT` or `SIGTERM` will terminate the program.

## alive

Catching term signals in a child goroutine with other running child goroutines

Run `./run.sh alive`.  You will see tasks spin up in goroutines.  Additionally, a 
signal task is dispatched.  This signal task listens for `SIGINT` and `SIGTERM` and 
returns an error to the caller (errgroup) when the signal is caught.  However, if 
you try to terminate the program before the last task runner finishes, the program 
won't exit.  It will wait for the last task runner to return an error before exiting.
Also note that even though you sent a termination/interrupt signal and the program 
didnt exit on it, it still registers as the cause for shutdown.

## none

No explicit catching of term signal

Run `./run.sh none`.  You will see tasks spin up in goroutines. No one is actively 
listening for `SIGTERM` or `SIGINT`. Before the last goroutine finishes, send `SIGTERM`
or `SIGINT` to the program.  Note that the program terminates immediately.

## why is this happening?

When termination signals are delegated to a child goroutine, they only 
immediately affect activity in itself and any subsequent child goroutines.  The signal 
WILL propagate upward to the parent goroutine, but the parent goroutine will not terminate 
until there are no other active child goroutines, or all other child goroutines are blocked.

### Caveat #1
In the `TaskRunner` struct, you could listen for context cancellation to terminate the goroutine.  However, that would push the responsibility of termination downstream, effectively making it "opt-in" logic.  The main goroutine still cedes control in this situation.
```go
if runner.Delay > 0 {
    select {
    case <-ctx.Done():
    case <-time.After(runner.Delay):
    }
}
```
