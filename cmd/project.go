package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	harbor "harbor/harbor_registry"
	"os"
)


var projectCmd = &cobra.Command{
	Use: "project",
	Short: "to operator project",
	Run: func(cmd *cobra.Command, args []string) {

		output, err := ExecuteCommand("harbor","project", args...)
		if err != nil {
			Error(cmd,args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

var projectLsCmd = &cobra.Command{
	Use: "ls",
	Short: "list  all project",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		if len(url) == 0 {
			fmt.Println("url is null,please specified the harbor url first !!!!")
			os.Exit(2)
		}
		output := harbor.GetProject(url)
		fmt.Println("项目名 访问级别 仓库数量")
		for i := 0; i < len(output); i++ {

			fmt.Println(output[i]["name"], output[i]["public"], output[i]["repo_count"])
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectLsCmd)
	projectLsCmd.Flags().StringP("url", "u", "","defaults: [https://127.0.0.1]")
}
