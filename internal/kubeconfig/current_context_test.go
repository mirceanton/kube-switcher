package kubeconfig

import (
	"path/filepath"
	"testing"
)

func TestCurrentContext(t *testing.T) {
	filePath := filepath.Join("testdata", "config-2.yaml")

	// Get the current context from the kubeconfig file
	currentContext, err := getCurrentContext(filePath)
	if err != nil {
		t.Fatalf("Failed to get current context from kubeconfig: %v", err)
	}

	// validate the current context
	expectedCurrentContext := "context-2"
	if currentContext != expectedCurrentContext {
		t.Errorf("unexpected current context, got %s, want %s", currentContext, expectedCurrentContext)
	}

	// set a valid current context
	newContext := "context-3"
	err = setCurrentContext(filePath, newContext)
	if err != nil {
		t.Fatalf("Failed to set current context: %v", err)
	}

	// validate the current context
	currentContext, err = getCurrentContext(filePath)
	if currentContext != newContext {
		t.Errorf("unexpected current context after setting, got %s, want %s", currentContext, newContext)
	}

	// set an invalid current context
	invalidContext := "invalid-context"
	err = setCurrentContext(filePath, invalidContext)
	if err == nil {
		t.Fatalf("expected error when setting invalid context, got nil")
	}
}