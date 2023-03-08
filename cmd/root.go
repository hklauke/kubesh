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

var namespace string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubesh pod-query",
	Short: "kubeshell into a pod",
	Long:  "a long description",
	RunE:  run,
}

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
	errCheck(err)

	if myClientConfig == nil {
		panic("Empty Client config")
	}

	matchSlice, containerMap := kubesh.GetResources(ctx, namespace, myClientSet)
	pod, container := kubesh.GetPrompt(matchSlice, containerMap)

	kubesh.StartConn(ctx, namespace, myClientSet, myClientConfig, pod, container)
	return nil

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

	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "desired kubernetes namespace, defaults to all")
}
