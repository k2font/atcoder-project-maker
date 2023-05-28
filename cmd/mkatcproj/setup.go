package mkatcproj

import (
	"fmt"

	"example.com/firstcli/package/mkatcproj"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:     "setup",
	Aliases: []string{"s"},
	Short:   "Setup a project or a file",
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		objectType, _ := cmd.Flags().GetString("type")
		fileName, _ := cmd.Flags().GetString("name")
		if lang == "cpp" {
			mkatcproj.CppObjMaker(fileName, objectType)
		} else if lang == "go" {
			// TODO: Go用オブジェクトを作成するための関数
			fmt.Println("Coming soon...")
		}
		fmt.Println("Project creation is complete!")
	},
}

func init() {
	// フラグ設定を追加
	setupCmd.Flags().StringP("lang", "l", "", "Languages you want to use")
	setupCmd.Flags().StringP("type", "t", "", "Objects you want to make")
	setupCmd.Flags().StringP("name", "n", "", "Name of project or file")
	// 必須のフラグを登録
	setupCmd.MarkFlagRequired("name")
	// setupコマンドを登録
	rootCmd.AddCommand(setupCmd)
}
