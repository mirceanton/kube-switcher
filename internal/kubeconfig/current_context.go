package kubeconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

// getCurrentContext returns the current context in the kubeconfig file
func getCurrentContext(filePath string) (string, error) {
	// Read the kubeconfig file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	// Unmarshal the kubeconfig file
	var kubeconfig clientcmdapi.Config
	if err := yaml.Unmarshal(data, &kubeconfig); err != nil {
		return "", fmt.Errorf("failed to unmarshal kubeconfig from file %s: %v", filePath, err)
	}

	return kubeconfig.CurrentContext, nil
}

// setCurrentContext sets the current context in the kubeconfig file
func setCurrentContext(filePath string, contextName string) error {
	// Read the kubeconfig file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	// Unmarshal the kubeconfig file
	var kubeconfig clientcmdapi.Config
	if err := yaml.Unmarshal(data, &kubeconfig); err != nil {
		return fmt.Errorf("failed to unmarshal kubeconfig from file %s: %v", filePath, err)
	}

	// Set the current context if it exists
	for _, context := range kubeconfig.Contexts {
		if context.Name == contextName {
			kubeconfig.CurrentContext = contextName
			return nil
		}
	}
	return fmt.Errorf("context '%s' does not exist in file %s", contextName, filePath)
}
