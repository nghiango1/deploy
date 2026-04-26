package push_test

import (
	"os"
	"testing"

	"github.com/nghiango1/deploy/webhook-handler/pkg/logger"
	"github.com/nghiango1/deploy/webhook-handler/pkg/webhook/push"
)


func TestDevRefHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
	}{
		{"default"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := push.DevRefHandler(push.Event{})
			if err != nil {
				t.Errorf("Failed to pull latest code")
			}
		})
	}
}

func TestMainRefHandler(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
	}{
		{"default"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := push.MainRefHandler(push.Event{})
			if err != nil {
				t.Errorf("Failed to pull latest code")
			}
		})
	}
}

func TestMain(m *testing.M) {
	logger.SetupLogger()
	os.Exit(m.Run())
}
