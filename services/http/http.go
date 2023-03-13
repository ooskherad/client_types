package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	setHeader(req)
	if err != nil {
		return nil, err
	}
	var httpClient = http.Client{}
	httpClient.Timeout = 1 * time.Second
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func GetAndString(url string) (string, error) {
	resp, err := Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func setHeader(req *http.Request) *http.Request {
	req.Header.Set(ConnectionKeyEnum, ConnectionValueEnum)
	req.Header.Set(CookieKeyEnum, CookieValueEnum)
	req.Header.Set(AcceptLanguageKeyEnum, AcceptLanguageValueEnum)
	req.Header.Set(HostKeyEnum, HostValueEnum)
	req.Header.Set(UpgradeInsecureRequestsKeyEnum, UpgradeInsecureRequestsValueEnum)
	req.Header.Set(UserAgentKeyEnum, UserAgentValueEnum)
	req.Header.Set(ProxyAuthorizationKeyEnum, ProxyAuthorizationValueEnum)
	return req
}
