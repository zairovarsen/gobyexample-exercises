//go:build solution

package recovery

import "fmt"

func DidPanic(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()
	return false
}

func PanicValue(f func()) (value any) {
	defer func() {
		value = recover()
	}()
	f()
	return nil
}

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			result, err = 0, fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

func Try(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = fmt.Errorf("panic: %w", e)
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	return f()
}

func RunAll(tasks []func()) []error {
	errs := make([]error, len(tasks))
	for i, task := range tasks {
		errs[i] = runTask(i, task)
	}
	return errs
}

// runTask gives every task its own stack frame to recover in.
func runTask(i int, task func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("task %d panicked: %v", i, r)
		}
	}()
	task()
	return nil
}
