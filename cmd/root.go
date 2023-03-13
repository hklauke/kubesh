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
	"github.com/spf13/pflag"
)

var namespace string
var customCom string

type options struct {
	namespace string
	command   string
	pod       string
	container string
}

func (myOptions *options) initFlags(fs *pflag.FlagSet) {
	fs.StringVar(&myOptions.command, "c", "/bin/sh", "use a custom command, defaults to /bin/sh")
	fs.StringVarP(&myOptions.namespace, "namespace", "n", "", "desired kubernetes namespace, defaults to all")
}

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

	myOptions := options{}
	myOptions.initFlags(cmd.Flags())
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

	matchSlice, containerMap := kubesh.GetResources(ctx, myOptions.namespace, myClientSet)
	pod, container := kubesh.GetPrompt(matchSlice, containerMap)

	kubesh.StartConn(ctx, namespace, myClientSet, myClientConfig, pod, container)
	return nil

}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}
