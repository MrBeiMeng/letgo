package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var HeaderMap = make(map[string]string)

var Cookies = "_gid=GA1.2.1633828563.1672150479; gr_user_id=fdfd41b7-6fc9-4c54-a5c0-b39e42bb9cd0; _bl_uid=CmlpqcaU6U2b2m79XzRFmy9qq9ma; a2873925c34ecbd2_gr_last_sent_cs1=beimengclub; csrftoken=Qg2oPO181zv8CQdSl68cyiD2pXyVFVEkgtEvMXvY4NKvqU31Xjj9mk7HlKO9YnPC; _ga_PDVPZYN3CW=GS1.1.1672223327.2.1.1672223365.0.0.0; _ga=GA1.2.1932331919.1672150479; NEW_QUESTION_DETAIL_PAGE_V2=1; aliyungf_tc=9104dc1f720e8887754329aa4765af05fb52108a4d5d57cec2691c437bef2f46; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1672150058,1672223220,1672223328,1672302020; NEW_PROBLEMLIST_PAGE=1; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNDA3MDI5OCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiZDVjNWJhYmUzYjZmODQxYzYwN2ViNGIxMGZhNjY0N2Y1MmEwMGZmZmRhNDhlZTlmNDdhZDhkOGMwNGJmNDY3NiIsImlkIjo0MDcwMjk4LCJlbWFpbCI6IjExOTIzODQ3MjJAcXEuY29tIiwidXNlcm5hbWUiOiJiZWltZW5nY2x1YiIsInVzZXJfc2x1ZyI6ImJlaW1lbmdjbHViIiwiYXZhdGFyIjoiaHR0cHM6Ly9hc3NldHMubGVldGNvZGUuY24vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9iZWltZW5nY2x1Yi9hdmF0YXJfMTY2NzUzNTc0OC5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTY3MjIyMzM2NC43NDk1MjAzLCJleHBpcmVkX3RpbWVfIjoxNjc0NzU5NjAwLCJ2ZXJzaW9uX2tleV8iOjAsImxhdGVzdF90aW1lc3RhbXBfIjoxNjcyMzE3MDk4fQ.G5OZMva64jdi73u1bbrYu1K5kRTfaq1QXtw6lwjItn4; Hm_lpvt_f0faad39bcf8471e3ab3ef70125152c3=1672317101; a2873925c34ecbd2_gr_session_id=28d7810b-e93b-4933-97ee-ad2d688b4e10; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=28d7810b-e93b-4933-97ee-ad2d688b4e10; a2873925c34ecbd2_gr_cs1=beimengclub; a2873925c34ecbd2_gr_session_id_28d7810b-e93b-4933-97ee-ad2d688b4e10=true; _gat=1"

func init() {
	HeaderMap["random-uuid"] = "ee0085a4-4dcd-1b24-d939-a573dd5c93ff"
	HeaderMap["content-type"] = "application/json"
	HeaderMap["origin"] = "https://leetcode.cn"
	HeaderMap["x-csrftoken"] = "Qg2oPO181zv8CQdSl68cyiD2pXyVFVEkgtEvMXvY4NKvqU31Xjj9mk7HlKO9YnPC"
	HeaderMap["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
}

func HttpPost(url string, cookies string, headerMap map[string]string, requestBody string) []byte {
	client := &http.Client{}
	// ...
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		panic(err)
	}

	for _, cookie := range strings.Split(cookies, ";") {
		cookie = strings.TrimSpace(cookie)
		cs := strings.Split(cookie, "=")
		name := cs[0]
		value := cs[1]
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return all
}
