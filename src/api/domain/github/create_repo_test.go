package github

import(
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T){
	request := CreateRepoRequest{
		Name: "golang introduction",
		Description: "a golang introduction repository",
		Homepage: "https://github.com",
		Private: true,
		Hasissues: true,
		Hasprojects: true,
		Haswiki: true,
	}

	JSONbytes, err := json.MarshalIndent(request, "", "  ")
	assert.Nil(t, err)
	assert.NotNil(t, JSONbytes)


	var target CreateRepoRequest
	err = json.Unmarshal(JSONbytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Hasissues, request.Hasissues)
}