package harbor

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ReposiData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ProjectId int `json:"project_id"`
	PullCount int `json:"pull_count"`
}

type RepoData []ReposiData

func GetRepoData(url string, proj string)  []map[string]string {
	// /api/v2.0/projects/goharbor/repositories
	url = url + "/api/v2.0/projects/" + proj +  "/repositories"
	request, _ := http.NewRequest(http.MethodGet, url,nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	request.Header.Set("accept", "application/json")
	//request.SetBasicAuth("admin", "Pwd123456")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("excute failed")
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	ret := RepoData{}
	json.Unmarshal([]byte(string(body)), &ret)
	var ps = []map[string]string{}
	for i := 0; i < len(ret); i++ {
		RData := make(map[string]string)
		RData["name"] = (ret[i].Name)
		pId := strconv.Itoa(ret[i].ProjectId)
		RData["project_id"] = pId
		RData["id"] =(strconv.Itoa(ret[i].Id))
		RData["pullCount"] = (strconv.Itoa(ret[i].PullCount))
		ps = append(ps, RData)
	}
	return ps
}
