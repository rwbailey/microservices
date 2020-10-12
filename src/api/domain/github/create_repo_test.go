package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "test-repo",
		Description: "This is a Test Repo Created Via The API.",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	// assert.EqualValues(t, string(bytes), `{"name":"test-repo","description":"This is a Test Repo Created Via The API.","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`)

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
}
