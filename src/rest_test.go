package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

const (
	REST_URL = "https://api.cosm.com/v2/"
	KEY      = "oMyWvwF_rWI8e5ULO1tW9pHAUOqSAKxSUm9nZDFIOGhwWT0g"
)

func NoTestNewRequest(t *testing.T) {
	req, e := http.NewRequest("GET", "http://example.com/", nil)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Log(string(body))
	}
}

func TestGetFeeds(t *testing.T) {
	var feedId = 98339
	var format = "json"
	var url = REST_URL + "feeds/" + strconv.Itoa(feedId) + "." + format + "?key=" + KEY
	t.Log(url)
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(string(bytes))
	var jsonObj = make(map[string]interface{})
	json.Unmarshal(bytes, &jsonObj)
	location := jsonObj["location"].(map[string]interface{})
	t.Log(location["domain"])
}
