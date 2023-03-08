package kubesh

import (
	"context"
	"fmt"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetResouces returns a list pods and their containers
func GetResources(ctx context.Context, namespace string, client kubernetes.Interface) (podsList []string, containers map[string][]string) {
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	var match []string
	cMap := make(map[string][]string)
	if err != nil {
		panic(err)
	}
	for _, v := range pods.Items {

		res := strings.Contains(v.Name, os.Args[1])
		//TODO make sure status is ready
		if res {
			match = append(match, v.Name)
			var containerSlice []string
			for _, k := range v.Spec.Containers {
				containerSlice = append(containerSlice, k.Name)

			}
			cMap[v.Name] = containerSlice

		}
	}

	if len(match) == 0 {
		fmt.Println("There are no pods that match in any namespace ")
		os.Exit(1)
	}
	return match, cMap
}
