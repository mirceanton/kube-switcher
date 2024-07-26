// FILEPATH: /workspaces/kube-switcher/internal/kubeconfig/context_test.go

package kubeconfig

import (
	"path/filepath"
	"testing"
)

func TestContexts(t *testing.T) {
	filePath := filepath.Join("testdata", "config-2.yaml")

	// Get contexts from the kubeconfig file
	contexts, err := getContexts(filePath)
	if err != nil {
		t.Fatalf("Failed to get contexts from kubeconfig: %v", err)
	}

	// Verify the number of contexts
	expectedContexts := 2
	if len(contexts) != expectedContexts {
		t.Errorf("Expected %d contexts, but got %d", expectedContexts, len(contexts))
	}

	// Verify the context names
	expectedContextNames := []string{"context-2", "context-3"}
	for i, context := range contexts {
		if context.Name != expectedContextNames[i] {
			t.Errorf("Expected context name %s, but got %s", expectedContextNames[i], context.Name)
		}
	}

	// Remove one of the contexts
	contextNameToRemove := "context-2"
	err = removeContext(filePath, contextNameToRemove)
	if err != nil {
		t.Fatalf("Failed to remove context from kubeconfig: %v", err)
	}

	// Get contexts from the kubeconfig file again
	contexts, err = getContexts(filePath)
	if err != nil {
		t.Fatalf("Failed to get contexts from kubeconfig: %v", err)
	}

	// Verify the number of contexts after removal
	expectedContextsAfterRemoval := 1
	if len(contexts) != expectedContextsAfterRemoval {
		t.Errorf("Expected %d contexts after removal, but got %d", expectedContextsAfterRemoval, len(contexts))
	}

	// Verify the context name after removal
	expectedContextNameAfterRemoval := "context-3"
	if contexts[0].Name != expectedContextNameAfterRemoval {
		t.Errorf("Expected context name %s after removal, but got %s", expectedContextNameAfterRemoval, contexts[0].Name)
	}
}
