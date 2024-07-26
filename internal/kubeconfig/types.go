package kubeconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

type KubeConfigFile struct {
	Path     string
	Config clientcmdapi.Config
}


// parseFile parses a kubeconfig file and returns the parsed config
func parseFile(filePath string) (*KubeConfigFile, error) {
	// Read the kubeconfig file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	// Unmarshal the kubeconfig file
	var kubeconfig clientcmdapi.Config
	if err := yaml.Unmarshal(data, &kubeconfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kubeconfig from file %s: %v", filePath, err)
	}

	kubeconfigFile := KubeConfigFile{
		Path:   filePath,
		Config: kubeconfig,
	}
	return &kubeconfigFile, nil
}
