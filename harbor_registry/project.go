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

type MetaData struct {
	Public string `json:"public"`
}
type ProjectData struct {
	MetaData MetaData `json:"metadata"`
	ProjectId int `json:"project_id"`
	Name string `json:"name"`
	RepoCount int `json:"repo_count"`
}

type PData []ProjectData

func GetProject(url string) []map[string]string {
	// curl -X GET "https://zhouhua.zaizai.com/api/v2.0/projects" -H "accept: application/json"
	url = url + "/api/v2.0/projects"
	//url = url + "/api/projects"
	request, _ := http.NewRequest(http.MethodGet, url,nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	request.Header.Set("accept", "application/json")
	//request.SetBasicAuth(username, password)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("excute failed")
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	ret := PData{}
	json.Unmarshal([]byte(string(body)), &ret)
	var ps = []map[string]string{}
	for i := 0; i < len(ret); i++ {
		RData := make(map[string]string)
		RData["name"] = (ret[i].Name)
		RData["project_id"] = strconv.Itoa(ret[i].ProjectId)
		RData["repo_count"] =strconv.Itoa(ret[i].RepoCount)
		RData["public"] = ret[i].MetaData.Public

		ps = append(ps, RData)
	}
	return ps
}
