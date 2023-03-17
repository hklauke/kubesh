package kubesh

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func GetPrompt(pods []string, containers map[string][]string) (pod string, container string) {
	var pod_result string
	if len(pods) > 1 {
		pod_result = callPrompt(pods)
	} else {
		pod_result = pods[0]
	}

	prompt_container := promptui.Select{
		Label: "Select container",
		Items: containers[pod_result],
	}

	_, pod_container, err := prompt_container.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println("You choose ", pod_result, pod_container)
	fmt.Println("")
	return pod_result, pod_container

}

func callPrompt(pods []string) string {
	prompt_pod := promptui.Select{
		Label: "Select pod",
		Items: pods,
	}

	_, pod_result, err := prompt_pod.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return pod_result

}
