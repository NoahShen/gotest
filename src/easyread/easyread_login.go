package easyread

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"io"
	"strings"
	"utils"
)

const (
	LOING_URL = "https://easyread.163.com/sns/login/login.atom"
)

type EasyreadSession struct {
	cookies []*http.Cookie
}

func (self *EasyreadSession) login(username, password string) error {
	loginInfo := make(map[string]interface{})
	loginInfo["accountType"] = 0
	loginInfo["auth"] = utils.MD5Encode(password)
	loginInfo["username"] = username
	loginInfoJson, _ := json.Marshal(loginInfo)
	b := strings.NewReader(string(loginInfoJson))
	req := self.createHttpRequest("POST", "application/json", LOING_URL, b)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	content, respErr := self.getResponseContent(resp)
	if respErr != nil {
		return respErr
	}
	fmt.Println("content:", string(content))
	return nil
}

func (self *EasyreadSession) createHttpRequest(method, contentType, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", contentType)
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
