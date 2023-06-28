# Go Context
Go Context is a Go library for carrying deadlines, cancellations signals, and other request-scoped values across API boundaries and between processes.

Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc. Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers. Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. The go vet tool checks that CancelFuncs are used on all control-flow paths.

The WithCancelCause function returns a CancelCauseFunc, which takes an error and records it as the cancellation cause. Calling Cause on the canceled context or any of its children retrieves the cause. If no cause is specified, `Cause(ctx)` returns the same value as `ctx.Err()`.

Programs that use Contexts should follow these rules to keep interfaces consistent across packages and enable static analysis tools to check context propagation:

- Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:
```go
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```
- Do not pass a nil Context, even if a function permits it. Pass `context.TODO` if you are unsure about which Context to use.

- Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.

- The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
