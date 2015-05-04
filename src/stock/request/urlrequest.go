/***************************************************************************
 *
 * Copyright (c) 2015 xxx.cn, Inc. All Rights Reserved
 * $Id$
 *
 **************************************************************************/

/**
 * @file urlrequest.go
 * @author divesli(divesli@gmail.com)
 * @date 2015/04/10 16:23:26
 * @version $Revision$
 * @filecoding UTF-8
 * @brief
 *
 */

package request

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type UrlRequest struct {
	method  string
	url     string
	refer   string
	timeout time.Duration
	params  map[string]string
	cookies map[string]string
	req     *http.Request
	resp    *http.Response
}

func Get(url string) *UrlRequest {
	return newUrlRequest(url, "GET")
}

func Post(url string) *UrlRequest {
	return newUrlRequest(url, "POST")
}

func newUrlRequest(url, method string) *UrlRequest {
	var resp http.Response
	req := http.Request{
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	return &UrlRequest{method: method, url: url, refer: "", timeout: 60 * time.Second, params: map[string]string{}, cookies: map[string]string{}, req: &req, resp: &resp}
}

func (r *UrlRequest) SetRefer(refer string) *UrlRequest {
	r.refer = refer
	return r
}

func (r *UrlRequest) SetParams(params map[string]string) *UrlRequest {
	r.params = params
	return r
}

func (r *UrlRequest) SetCookie(cookies map[string]string) *UrlRequest {
	r.cookies = cookies
	return r
}

func (r *UrlRequest) SetTimeout(n int64) *UrlRequest {
	r.timeout = time.Duration(n) * time.Second
	return r
}

func (r *UrlRequest) GetStatus() (int, error) {
	resp, err := r.getResponse()
	if err != nil {
		return -1, err
	}
	return int(resp.StatusCode), nil
}

func (r *UrlRequest) GetLocation() (string, error) {
	resp, err := r.getResponse()
	if err != nil {
		return "", err
	}
	return resp.Header.Get("Location"), nil
}

func (r *UrlRequest) GetCookies() (cookies []map[string]string, err error) {
	resp, err := r.getResponse()
	if err != nil {
		return nil, err
	}
	var httpcookies []*http.Cookie
	httpcookies = resp.Cookies()
	cookies = []map[string]string{}
	for _, httpcookie := range httpcookies {
		cookie := make(map[string]string, 5)
		cookie["name"] = httpcookie.Name
		cookie["value"] = httpcookie.Value
		cookie["path"] = httpcookie.Path
		cookie["domain"] = httpcookie.Domain
		cookie["expires"] = httpcookie.Expires.Format("2006-01-02 15:04:05")
		cookies = append(cookies, cookie)
	}
	return cookies, nil
}

func (r *UrlRequest) GetBody() (string, error) {
	resp, err := r.getResponse()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//defer resp.Body.Close()
	return string(body), nil
}

func (r *UrlRequest) Close() {
	if r.resp.StatusCode != 0 && r.resp.Body != nil {
		r.resp.Body.Close()
	}
}

func (r *UrlRequest) buildUrl(paramstr string) {
	if r.method == "GET" && len(paramstr) > 0 {
		if strings.Index(r.url, "?") != -1 {
			r.url += "&" + paramstr
		}
		r.url += "?" + paramstr
	}

	if r.method == "POST" && r.req.Body == nil {
		if len(paramstr) > 0 {
			r.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			bf := bytes.NewBufferString(paramstr)
			r.req.Body = ioutil.NopCloser(bf)
			r.req.ContentLength = int64(len(paramstr))
		}
	}
}
func (r *UrlRequest) getResponse() (*http.Response, error) {
	if r.resp.StatusCode != 0 {
		return r.resp, nil
	}

	if len(r.refer) > 0 {
		r.req.Header.Add("Refer", r.refer)
	}

	if len(r.cookies) > 0 {
		for name, value := range r.cookies {
			cookie := &http.Cookie{Name: name, Value: value}
			r.req.AddCookie(cookie)
		}
	}

	var paramstr string
	if len(r.params) > 0 {
		var buf bytes.Buffer
		for k, v := range r.params {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
		paramstr = buf.String()
		paramstr = paramstr[0 : len(paramstr)-1]
	}
	r.buildUrl(paramstr)
	trans := &http.Transport{
		/*Dial: func(network, addr string) (c net.Conn, err error) {
			return net.DialTimeout(network, addr, r.timeout)
		},*/
		Dial: TimeoutDialer(r.timeout),
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		},
		Transport: trans,
	}

	url, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	r.req.URL = url
	resp, err := client.Do(r.req)
	if err != nil && strings.Index(err.Error(), "Redirect") == -1 {
		return nil, err
	}

	r.resp = resp
	return r.resp, nil
}

func TimeoutDialer(timeout time.Duration) func(network, addr string) (c net.Conn, err error) {
	return func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, timeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(timeout))
		return conn, nil
	}
}
