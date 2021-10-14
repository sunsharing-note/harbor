package cmd

import (
	"fmt"
	"harbor/harbor_registry"
	"github.com/spf13/cobra"
	"os"
)

var repoCmd = &cobra.Command{
	Use: "repo",
	Short: "to operator repository",
	Run: func(cmd *cobra.Command, args []string) {

		output, err := ExecuteCommand("harbor","repo", args...)
		if err != nil {
			Error(cmd,args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

var repoLsCmd = &cobra.Command{
	Use: "ls",
	Short: "list  project's repository",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		project, _ := cmd.Flags().GetString("project")
		if len(project) == 0 {
			fmt.Println("sorry, you must specified the project which you want to show repository !!!")
			os.Exit(2)
		}
		output := harbor.GetRepoData(url, project)
		fmt.Println("仓库名----------拉取次数")
		for i := 0; i < len(output); i++ {
			fmt.Println(output[i]["name"],output[i]["pullCount"])
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoLsCmd)
	repoLsCmd.Flags().StringP("url", "u", "","defaults: [https://127.0.0.1]")
	repoLsCmd.Flags().StringP("project", "p","", "the project")
}
