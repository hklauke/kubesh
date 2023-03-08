/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hklauke/kubesh/client"
	"github.com/hklauke/kubesh/kubesh"

	"github.com/spf13/cobra"
)

var namespaceVar string

type shQuery struct {
	pod       string
	container string
	namespace string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubesh pod-query",
	Short: "kubeshell into a pod",
	//Long:  "a long description",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if len(os.Args) < 2 {
	// 		panic("Must include at least one argument")
	// 	}
	// },

	//TODO: If no args print help
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	config := shQuery{
		namespace: namespaceVar,
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	myClientSet, myClientConfig, err := client.GetLocal()
	errCheck(err)

	if myClientConfig == nil {
		panic("Empty Client config")
	}

	matchSlice, containerMap := kubesh.GetResources(ctx, config, myClientSet)
	pod, container := kubesh.GetPrompt(matchSlice, containerMap)
	fmt.Printf("%T", myClientSet)
	fmt.Println(pod, container)
	//kubesh.StartConn(ctx, namespaceVar, myClientSet, myClientConfig, pod, container)

}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubesh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().StringVarP(&namespaceVar, "namespace", "n", "", "desired kubernetes namespace")
}
