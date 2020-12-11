package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateRepoIvalidJSONRequest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/create/github", strings.NewReader(``))
	c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	var errResp ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &errResp); err != nil {
		fmt.Println("can not unmarshal errorResponse")
	}

	assert.EqualValues(t, "Invalid JSON body", errResp.Message)
	assert.EqualValues(t, http.StatusBadRequest, errResp.StatusCode)

}

func TestCreateRepoInvalidProvider(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "/create/bitbucket", strings.NewReader(`{"name": "test", "description": "testtest", "private": true}`))
	c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	var errResp ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &errResp); err != nil {
		fmt.Println("can not unmarshal response")
	}

	assert.EqualValues(t, http.StatusBadRequest, errResp.StatusCode)
	assert.EqualValues(t, "There is no provider with given name", errResp.Message)

}

func TestCreateMultipleRepositoryInvalidJson(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(http.MethodPost, "createMultiple", strings.NewReader(`{misssingQuote":[
		{"name": "test4", "description": "testtest", "private": true},
		{"name": "test5", "description": "testtest", "private": true},
		{"name": "test6", "description": "testtest", "private": true}
		]}`))

	c.Request = request

	CreateMultipleRepository(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	var errResp ErrorResponse
	if err := json.Unmarshal(response.Body.Bytes(), &errResp); err != nil {
		fmt.Println("can not unmarshal errorResponse")
	}

	assert.EqualValues(t, "Invalid JSON body", errResp.Message)
	assert.EqualValues(t, http.StatusBadRequest, errResp.StatusCode)
}
