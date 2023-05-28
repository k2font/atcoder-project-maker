package mkatcproj

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1" // TODO: バージョン情報を他から取得するようにしたい

var rootCmd = &cobra.Command{
	Use:     "mkatcproj",
	Version: version,
	Short:   "mkatcproj make quickly an atcoder project.",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error is occured!")
		os.Exit(1)
	}
}
