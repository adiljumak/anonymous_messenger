package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
var keeper = getDummyKeeper()

func handleTestRequest(w *httptest.ResponseRecorder, r *http.Request) {
	keyBuilder := getKeyBuilder()
	router := getRouter(keyBuilder, keeper)
	router.ServeHTTP(w, r)
}

func TestIndexPage(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("index page is not 200")
	}
}

func TestSaveMessage(t *testing.T) {
	testMessage := "foo"
	postData := strings.NewReader(fmt.Sprintf("message=%s", testMessage))
	request, _ := http.NewRequest("POST", "/", postData)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("save is not 200")
	}
	keyBuilder := getKeyBuilder()
	key, _ := keyBuilder.Get()
	savedMessage, _ := keeper.Get(key)

	if savedMessage != testMessage {
		t.Error("message was not saved")
	}

	result := w.Result()
	defer result.Body.Close()
	data, _ := ioutil.ReadAll(result.Body)
	if !strings.Contains(string(data), key) {
		t.Error("result page without key")
	}
}

func TestReadMessage(t *testing.T) {
	testMessage := "helloMessage"
	keyBuilder := getKeyBuilder()
	key, _ := keyBuilder.Get()
	keeper.Set(key, testMessage)

	request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 200 {
		t.Error("read message is not 200")
	}

	result := w.Result()
	defer result.Body.Close()
	data, _ := ioutil.ReadAll(result.Body)
	if !strings.Contains(string(data), testMessage) {
		t.Error("result page without key")
	}

	_, err := keeper.Get(key)
	if err == nil {
		t.Error("delete now working, keeper value must be empty")
	}
}

func TestReadMessageNotFound(t *testing.T) {
	keyBuilder := getKeyBuilder()
	key, _ := keyBuilder.Get()
	request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	w := httptest.NewRecorder()
	handleTestRequest(w, request)
	if w.Code != 404 {
		t.Error("empty message must be 404")
	}
}
