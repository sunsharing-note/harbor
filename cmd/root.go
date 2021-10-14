package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "harbor",
	Short: "a docker image registry",
	Long: "a registry to store and pull or push image",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("uncognized command"))
	},
}


func Excute() {
	rootCmd.Execute()
}