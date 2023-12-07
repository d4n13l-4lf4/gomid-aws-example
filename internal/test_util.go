package internal

import (
	"errors"
	"fmt"
	"log"
)

type (
	// Asserter is an assertion implementation for tests.
	Asserter struct {
		err error
	}

	// TestKey custom test key for contexts.
	TestKey string
)

// NewAsserter creates a new *Asserter.
func NewAsserter() *Asserter {
	return &Asserter{}
}

// Errorf formats an error according to fmt.Errorf rules.
func (a *Asserter) Errorf(format string, args ...any) {
	a.err = fmt.Errorf(format, args...)
}

// Logf logs an error according to log.Printf rules.
func (a *Asserter) Logf(format string, args ...any) {
	log.Printf(format, args...)
}

// FailNow asserter as failed.
func (a *Asserter) FailNow() {
	a.err = errors.New("fail now")
}

// Error returns the inner error.
func (a *Asserter) Error() error {
	return a.err
}
