package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available Kubernetes context",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine the kubeconfig directory
		if configDir == "" {
			configDir = os.Getenv("KUBESWITCHER_CONFIG_DIR")
			if configDir == "" {
				log.Fatal("KUBESWITCHER_CONFIG_DIR environment variable is not set")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
