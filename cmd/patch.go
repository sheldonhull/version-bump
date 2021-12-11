package cmd

import (
	"github.com/sheldonhull/version-bump/bump"

	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Increment a patch version",
	Run: func(cmd *cobra.Command, args []string) {
		run(bump.Patch)
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
