package repositories

import (
	"../../clients/restclient"
	"../../domains/repositories"
	"../../utils/errors"
	"../../utils/test_utils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockUps()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/repostiories", strings.NewReader(``))
	c := test_utils.GetMockedContext(req, res)
	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, res.Code)

	apiErr, err := errors.NewApiErrorFromBytes(res.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoFromGitHub(t *testing.T) {
	restclient.FlushMockUps()
	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": ""}`)),
		},
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/repostiories", strings.NewReader(`{"name": "lib_test"}`))
	c := test_utils.GetMockedContext(req, res)
	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, res.Code)

	apiErr, err := errors.NewApiErrorFromBytes(res.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoSuccess(t *testing.T) {
	restclient.FlushMockUps()
	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/repostiories", strings.NewReader(`{"name": "lib_test"}`))
	c := test_utils.GetMockedContext(req, res)
	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, res.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
}
