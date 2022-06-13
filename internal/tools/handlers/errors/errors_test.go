package error_handler

import (
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("code has panicked as expected: %v", r)
		}
	}()

	t.Run("do not panic", func(t *testing.T) {
		Handle(nil)
	})

	t.Run("panics", func(t *testing.T) {
		t.Skip()
		defer func() {
			if r := recover(); r != nil {
				t.Logf("code has panicked as expected: %v", r)
			}
		}()
		Handle(fmt.Errorf("any-error"))
	})
}
