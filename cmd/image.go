package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	harbor "harbor/harbor_registry"
	"os"
	"strconv"
)

var tagCmd = &cobra.Command{
	Use: "tag",
	Short: "to operator image",
	Run: func(cmd *cobra.Command, args []string) {

		output, err := ExecuteCommand("harbor","tag", args...)
		if err != nil {
			Error(cmd,args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

var tagLsCmd = &cobra.Command{
	Use: "ls",
	Short: "list  all tags of the repository you have specified which you should specified project at the same time",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		project, _ := cmd.Flags().GetString("project")
		repo, _ := cmd.Flags().GetString("repo")
		ss := harbor.GetTags(url, project, repo)

		fmt.Println(ss)
	},
}

var tagDelCmd = &cobra.Command{
	Use: "del",
	Short: "delete the tags of the repository you have specified which you should specified project at the same time",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		project, _ := cmd.Flags().GetString("project")
		repo, _ := cmd.Flags().GetString("repo")
		tag,_ := cmd.Flags().GetString("tag")
		count,_ := cmd.Flags().GetString("count")
		ret,_ := strconv.Atoi(count)
		if len(tag) != 0 && ret != 0 {
				fmt.Println("You can't choose both between count and tag")
				os.Exit(2)
		} else if len(tag) == 0 && ret != 0 {
			retu := harbor.GetTags(url, project, repo)
			rTagCount, _ := strconv.Atoi((retu[0])["count"])
			if ret == rTagCount {
				fmt.Printf("the repository %s of the project %s only have %d tags, so you can't delete tags and we will do nothing!!\n", repo, project,ret)
			} else if ret > rTagCount {
				fmt.Printf("the repository %s of the project %s only have %d tags, but you want to delete %d tags, so we suggest you to have a rest and we will do nothing!!\n", repo, project,rTagCount, ret)
			} else {
				fmt.Printf("we will save the latest %d tags  and delete other %d tags !!!\n", ret, (rTagCount - ret))
				tags := harbor.GetTags(url, project, repo)
				retu := harbor.DeleTagsByCount(tags, (rTagCount - ret))
				for i := 0 ; i < len(retu); i++ {
					code, msg := harbor.DelTags(url, project, repo, retu[i])
					fmt.Printf("the tag %s is deleted,status code is %d, msg is %s\n", retu[i], code, msg)
				}
			}
		} else {
			code, msg := harbor.DelTags(url, project, repo, tag)
			fmt.Println(code, msg["errors"])
		}
	},
}



func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.AddCommand(tagLsCmd)
	tagLsCmd.Flags().StringP("url", "u", "","defaults: [https://127.0.0.1]")
	tagLsCmd.Flags().StringP("project", "p", "","the project")
	tagLsCmd.Flags().StringP("repo", "r", "","the repository")

	tagCmd.AddCommand(tagDelCmd)
	tagDelCmd.Flags().StringP("url", "u", "","defaults: [https://127.0.0.1]")
	tagDelCmd.Flags().StringP("project", "p", "","the project which you should specified if you want to delete the tag of any repository ")
	tagDelCmd.Flags().StringP("repo", "r", "","the repository which you should specified if you want to delete the tag")
	tagDelCmd.Flags().StringP("tag", "t", "","the tag, You can't choose  it with tag together")
	tagDelCmd.Flags().StringP("count", "c", "","the total number you want to save.for example: you set --count=10, we will save the 10 latest tags by use push_time to sort,can't choose it with tag together")
}