package kubesh

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func GetPrompt(pods []string, containers map[string][]string) (pod string, container string) {

	prompt_pod := promptui.Select{
		Label: "Select pod",
		Items: pods,
	}

	_, pod_result, err := prompt_pod.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
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

	return pod_result, pod_container

}
