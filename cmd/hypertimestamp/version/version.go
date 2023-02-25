package version

import (
	"fmt"

	"github.com/aphexmunky/hypertimestamp/consts"
	"github.com/aphexmunky/hypertimestamp/version"
	"github.com/spf13/cobra"
)

func init() {
	cobra.EnablePrefixMatching = true
}

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of hypertimestamp",
		RunE:  versionFunc,
	}
}

func versionFunc(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s@%s (%s)\n", consts.Name, version.Version, consts.ID)
	return nil
}
