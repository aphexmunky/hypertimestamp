package main

import (
	"fmt"
	"os"

	"github.com/aphexmunky/hypertimestamp/cmd/hypertimestamp/version"
	"github.com/aphexmunky/hypertimestamp/controller"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        "hypertimestamp",
	Short:      "hypertimestamp agent",
	SuggestFor: []string{"hypertimestamp"},
	RunE:       runFunc,
}

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCmd.AddCommand(
		version.NewCommand(),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "hypertimestmap failed %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runFunc(cmd *cobra.Command, args []string) error {
	rpcchainvm.Serve(controller.New())
	return cmd.Help()
}
