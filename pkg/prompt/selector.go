package prompt

import (
	"github.com/mirceanton/kube-switcher/internal/kubeconfig"

	"github.com/manifoldco/promptui"
)

// SelectContext prompts the user to select a Kubernetes context.
func SelectContext(configFiles []string) (string, error) {
	index := 0
	items := []string{}

	for _, file := range configFiles {
		contexts, err := kubeconfig.GetContexts(file)
		if err != nil {
			return "", err
		}

		for _, context := range contexts {
			items[index] = context.Name
			index++
		}
	}

	prompt := promptui.Select{
		Label: "Select Kubernetes Context",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
