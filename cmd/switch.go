package cmd

import (
	"os"
	"path/filepath"

	"github.com/mirceanton/kube-switcher/internal/kubeconfig"
	"github.com/mirceanton/kube-switcher/pkg/prompt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CLI flags
var (
	contextName string
)

// switchCmd represents the switch command to change Kubernetes contexts.
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch Kubernetes context",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine the kubeconfig directory
		if configDir == "" {
			configDir = os.Getenv("KUBESWITCHER_CONFIG_DIR")
			if configDir == "" {
				log.Fatal("KUBESWITCHER_CONFIG_DIR environment variable is not set")
			}
		}

		// Parse all kubeconfigs in the specified directory
		contexts, err := kubeconfig.ParseKubeconfigs(configDir)
		if err != nil {
			log.Fatalf("Error parsing kubeconfig files: %v", err)
		}

		// If context name is not provided, prompt user to select one
		if len(args) == 0 {
			contextName, err = prompt.SelectContext(contexts)
			if err != nil {
				log.Fatalf("Error selecting context: %v", err)
			}

		} else if len(args) == 1 {
			contextName = args[0]
		} else {
			log.Fatalf("Too many arguments!")
		}
		log.Debugf("Selected context: %s", contextName)

		// Find and copy the selected kubeconfig file
		for _, context := range contexts {
			if context.Name == contextName {
				err = kubeconfig.CopyConfig(context.FilePath, filepath.Join(os.Getenv("HOME"), ".kube", "config"))
				if err != nil {
					log.Fatalf("Error copying kubeconfig file: %v", err)
				}
				log.Infof("Switched to context '%s' from file '%s'", context.Name, context.FilePath)
				return
			}
		}

		log.Fatalf("Context '%s' not found", contextName)
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
