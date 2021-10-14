package harbor

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func HarborParams() map[string]string {
	url := flag.String("url", "https://127.0.0.1", "must: docker registry url")
	project := flag.String("project", "", "optional: harbor registry project,sometime maybe must")
	repository := flag.String("repository", "", "optional: harbor registry repository,sometime maybe must")
	tagCount := flag.Int("tagCount",0, "optional: tag count,sometime maybe must")
	flag.Parse()
	if url == nil || *url == "https://127.0.0.1" {
		fmt.Println("warn: please give the docker registry url")
		flag.PrintDefaults()
		os.Exit(2)
	}
	var param = make(map[string]string)
	param["url"] = *url
	param["project"] = *project
	param["repository"] = *repository
	param["tagCount"] = strconv.Itoa(*tagCount)

	return param
}
