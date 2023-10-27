package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://docs.npmjs.com/cli/v8/using-npm/registry
const NODEJS_REGISTRY_API = "https://registry.npmjs.org/%s"

type NodejsRegistryResponse struct {
	Repository struct {
		Url string `json:"url"`
	}
}

type Nodejs struct {
}

func (n *Nodejs) fetchURLFromRegistry(name string) (string, error) {
	url := fmt.Sprintf(NODEJS_REGISTRY_API, name)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)

	var NodejsRegistryResponse NodejsRegistryResponse
	err := json.Unmarshal(body, &NodejsRegistryResponse)
	if err != nil {
		return "", nil
	}

	return NodejsRegistryResponse.Repository.Url, nil
}
