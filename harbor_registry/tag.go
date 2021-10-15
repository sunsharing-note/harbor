package harbor

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Tag struct {
	ArtifactId int 	`json:"artifact_id"`
	Id int `json:"id"`
	Name string `json:"name"`
	RepositoryId int `json:"repository_id"`
	PushTimte string `json:"push_time"`
}

type Tag2 struct {
	ArtifactId string 	`json:"artifact_id"`
	Id string `json:"id"`
	Name string `json:"name"`
	RepositoryId string `json:"repository_id"`
	PushTimte string `json:"push_time"`
}

type Tag2s []Tag2
// delete tag by specified count
func DeleTagsByCount(tags []map[string]string ,count int) []string {
	var re []string

	tt := tags[0]["tags"]
	ss := Tag2s{}
	json.Unmarshal([]byte(tt), &ss)

	// have a sort
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i].PushTimte > ss[j].PushTimte {
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}
	// get all tags
	for i := 0; i < len(ss); i++ {
		re = append(re, ss[i].Name)
	}

	return re[0:count]
}

// delete tag by specified tag
func DelTags(url string, project string, repo string, tag string) (int, map[string]interface{})  {
	url = url + "/api/v2.0/projects/" + project + "/repositories/" + repo + "/artifacts/" + tag + "/tags/" + tag
	request, _ := http.NewRequest(http.MethodDelete, url,nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	request.Header.Set("accept", "application/json")
	request.SetBasicAuth("admin", "Harbor12345")
	response,_ := client.Do(request)
	defer response.Body.Close()

	var result map[string]interface{}
	bd, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(bd, &result)
	}
	return response.StatusCode,result

}





type ArtiData struct {
	Id int `json:"id"`
	ProjectId int `json:"project_id"`
	RepositoryId int `json:"repository_id"`
	//Digest string `json:"digest"`
	Tags []Tag `json:"tags"`
}

type AData []ArtiData

func GetTags(url string, project string, repo string) []map[string]string {
	url = url + "/api/v2.0/projects/" + project + "/repositories/" + repo + "/artifacts"
	request, _ := http.NewRequest(http.MethodGet, url,nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	request.Header.Set("accept", "application/json")
	request.Header.Set("X-Accept-Vulnerabilities", "application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0")
	request.SetBasicAuth("admin", "Harbor12345")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("excute failed")
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	ret := AData{}
	json.Unmarshal([]byte(string(body)),&ret)
	var ps = []map[string]string{}
	sum := 0
	RData := make(map[string]string)
	RData["name"] = repo
	for i := 0; i < len(ret); i++ {

		RData["id"] = (strconv.Itoa(ret[i].Id))
		RData["project_id"] = (strconv.Itoa(ret[i].ProjectId))
		RData["repository_id"] =(strconv.Itoa(ret[i].RepositoryId))
		//RData["digest"] = ret[i].Digest
		var tdata = []map[string]string{}
		sum = len((ret[i].Tags))
		for j := 0; j < len((ret[i].Tags)); j++ {
			TagData := make(map[string]string)
			TagData["artifact_id"] = strconv.Itoa((ret[i].Tags)[j].ArtifactId)
			TagData["id"] = strconv.Itoa((ret[i].Tags)[j].Id)
			TagData["name"] = (ret[i].Tags)[j].Name
			TagData["repository_id"] = strconv.Itoa((ret[i].Tags)[j].RepositoryId)
			TagData["push_time"] = (ret[i].Tags)[j].PushTimte
			tdata = append(tdata, TagData)
		}
		RData["count"] = strconv.Itoa(sum)
		ss, err := json.Marshal(tdata)
		if err != nil {
			fmt.Println("failed")
			os.Exit(2)
		}
		RData["tags"] = string(ss)
		ps = append(ps, RData)

	}
	return ps
}
