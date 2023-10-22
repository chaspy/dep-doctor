package bundler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://guides.rubygems.org/rubygems-org-api/
const RUBYGEMS_ORG_API = "https://rubygems.org/api/v1/gems/%s.json"

type GemResponse struct {
	Name          string `json:"name"`
	SourceCodeUri string `json:"source_code_uri"`
}

func FetchFromRubyGems(name string) string {
	url := fmt.Sprintf(RUBYGEMS_ORG_API, name)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)

	var Gem GemResponse
	json.Unmarshal(body, &Gem)

	return Gem.SourceCodeUri
}
