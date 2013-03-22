package kuaipan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

func TestLoginKuaipan(t *testing.T) {
	jar := new(Jar)
	client := http.Client{nil, nil, jar}
	indexRep, errIndex := http.NewRequest("GET", "https://www.kuaipan.cn/account_login.htm", nil)
	indexRep.Header.Set("Content-Type", "text/html; charset=utf-8")
	indexRep.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.172 Safari/537.22")
	if errIndex != nil {
		t.Fatal(errIndex)
	}
	indexResp, indexRespErr := client.Do(indexRep)
	if indexRespErr != nil {
		t.Fatal(indexRespErr)
	}
	cookies := indexResp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}

	loginReq, err := http.NewRequest("POST", "https://www.kuaipan.cn/index.php?ac=account&op=login", nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, cookie := range cookies {
		loginReq.AddCookie(cookie)
	}

	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	loginReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.172 Safari/537.22")
	loginReq.Header.Set("Accept", "application/json, text/javascript, */*")
	loginReq.Header.Set("Referer", "https://www.kuaipan.cn/account_login.htm")

	loginReq.Form = make(url.Values)
	loginReq.Form.Set("username", "Noahs-Ark@163.com")
	loginReq.Form.Set("userpwd", "1234567")
	loginReq.Form.Set("isajax", "yes")
	loginReq.Form.Set("rememberme", "1")

	fmt.Println("login request:", loginReq)

	resp, err1 := client.Do(loginReq)
	if err1 != nil {
		t.Fatal(err1)
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		t.Fatal(err2)
	}
	fmt.Println("login result:", string(b))
	fmt.Println("login resp:", resp)
	resp.Body.Close()

	//resp2, _ := client.Get("http://www.kuaipan.cn/index.php?ac=common&op=usersign")

	//b2, _ := ioutil.ReadAll(resp2.Body)
	//resp.Body.Close()

	//fmt.Println("sign result:", string(b2))
}
