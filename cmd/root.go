/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/hklauke/kubesh/client"
	"github.com/hklauke/kubesh/kubesh"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&customCom, "command", "c", "/bin/sh", "use a custom command, defaults to /bin/sh")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "desired kubernetes namespace, defaults to default")
}

// rootCmd represents the base command when called without any subcommands
var (
	namespace string
	customCom string

	rootCmd = &cobra.Command{
		Use:   "kubesh pod-query",
		Short: "kubeshell into a pod",
		Long:  "a long description",
		RunE:  run,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {

	if len(args) == 0 {
		return errors.New("missing pod query")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	myClientSet, myClientConfig, err := client.GetLocal()
	if err != nil {
		os.Exit(1)
	}

	if myClientConfig == nil {
		errors.New("Empty Client config")
	}

	matchSlice, containerMap := kubesh.GetResources(ctx, namespace, myClientSet)
	pod, container := kubesh.GetPrompt(matchSlice, containerMap)

	kubesh.StartConn(ctx, namespace, myClientSet, myClientConfig, pod, container, customCom)
	return nil

}
