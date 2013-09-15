package easyread

import (
	"encoding/json"
	"encoding/xml"
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
	XMLName xml.Name          `xml:"usrsubsummary"`
	Entries []SubSummaryEntry `xml:"entry"`
}

type SubSummaryEntry struct {
	XMLName xml.Name    `xml:"entry"`
	Id      string      `xml:"id"`
	Name    string      `xml:"title"`
	Status  EntryStatus `xml:"entry_status"`
}

type EntryStatus struct {
	XMLName xml.Name `xml:"entry_status"`
	Type    string   `xml:"type,attr"`
	Style   string   `xml:"style,attr"`
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

func (self *EasyreadSession) getSubSummary() (SubSummary, error) {
	var subSummar = SubSummary{}
	url := fmt.Sprintf(GET_SUBSUMMARY_URL, time.Now().UTC().Unix())
	req := self.createHttpRequest("GET", url, "", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return subSummar, err
	}
	content, respErr := self.getResponseContent(resp)
	if respErr != nil {
		return subSummar, respErr
	}
	unmarshalErr := xml.Unmarshal(content, &subSummar)
	if unmarshalErr != nil {
		return subSummar, unmarshalErr
	}
	return subSummar, nil
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
