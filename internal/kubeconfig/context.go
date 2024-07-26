package kubeconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

// getContexts returns all contexts in the kubeconfig file
func getContexts(filePath string) ([]clientcmdapi.NamedContext, error) {
	// Parse the kubeconfig file
	kubeconfig, err := parseFile(filePath)
	if err != nil {
		return nil, err
	}

	return kubeconfig.Config.Contexts, nil
}

// removeContext removes a context from the kubeconfig file
func removeContext(filePath string, contextName string) error {
	// Read the kubeconfig file
	kubeconfig, err := parseFile(filePath)
	if err != nil {
		return err
	}

	// Remove the context if it exists
	for i, context := range kubeconfig.Config.Contexts {
		if context.Name == contextName {
			kubeconfig.Config.Contexts = append(
				kubeconfig.Config.Contexts[:i],
				kubeconfig.Config.Contexts[i+1:]...
			)

			// write back the file
			data, err := yaml.Marshal(kubeconfig.Config)
			if err != nil {
				return fmt.Errorf("failed to marshal kubeconfig to file %s: %v", filePath, err)
			}
			if err := os.WriteFile(filePath, data, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %v", filePath, err)
			}
			return nil
		}
	}

	return fmt.Errorf("context '%s' does not exist in file %s", contextName, filePath)
}
