package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version() {
	//fmt.Println("Version:", config.C.Info.Version)
	//fmt.Println("Commit:", config.C.Info.Commit)
	//fmt.Println("Build:", config.C.Info.Build)
}
