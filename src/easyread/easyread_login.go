package easyread

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"utils"
)

const (
	LOGIN_URL          = "https://easyread.163.com/sns/login/login.atom"
	GET_SUBSUMMARY_URL = "http://easyread.163.com/user/subsummary.atom?rand=%d"
)

type SubSummary struct {
	Id   string
	Name string
	Type string
}

type EasyreadSession struct {
	UserId   string
	Username string
	cookies  []*http.Cookie
}

func (self *EasyreadSession) login(username, password string) error {
	loginInfo := make(map[string]interface{})
	loginInfo["accountType"] = 0
	loginInfo["auth"] = utils.MD5Encode(password)
	loginInfo["username"] = username
	loginInfoJson, _ := json.Marshal(loginInfo)
	b := strings.NewReader(string(loginInfoJson))
	req := self.createHttpRequest("POST", LOGIN_URL, "application/json", b)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	content, respErr := self.getResponseContent(resp)
	if respErr != nil {
		return respErr
	}
	loginResult := make(map[string]interface{})
	unmarshalErr := json.Unmarshal(content, &loginResult)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	resCode := loginResult["resCode"].(float64)
	if resCode != 0 {
		return errors.New("resCode is not zero!")
	}
	userInfo := loginResult["userInfo"].(map[string]interface{})
	self.Username = userInfo["urs"].(string)
	self.UserId = userInfo["userId"].(string)
	return nil
}

func (self *EasyreadSession) getSubSummary() ([]SubSummary, error) {
	subsummaries := make([]SubSummary, 0)
	url := fmt.Sprintf(GET_SUBSUMMARY_URL, time.Now().UTC().Unix())
	req := self.createHttpRequest("GET", url, "", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return subsummaries, err
	}
	content, respErr := self.getResponseContent(resp)
	if respErr != nil {
		return subsummaries, respErr
	}
	fmt.Println("content:", string(content))
	return subsummaries, nil
}

func (self *EasyreadSession) createHttpRequest(method, url, contentType string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("User-Agent", "Pris/3.0.0")
	req.Header.Set("X-User-Agent", "PRIS/3.0.0 (768/1184; android; 4.3; zh-CN; android) 1.2.8")
	req.Header.Set("X-Pid", "(000000000000000;d41d8cd98f00b204e9800998ecf8427e;)")
	for _, cookie := range self.cookies {
		req.AddCookie(cookie)
	}
	return req
}

func (self *EasyreadSession) getResponseContent(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	cookies := resp.Cookies()
	self.cookies = cookies
	return ioutil.ReadAll(resp.Body)
}
