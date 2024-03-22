/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-04-30 12:07:17
 * LastEditTime: 2023-05-01 23:54:01
 * Description:
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_server_util

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// RegisterHelloWorldMethod_1 第一种注册一个hello接口的方法
//
//	@param serveMux 注册到指定的 multiplexer
//	@param pattern 接口url
func RegisterHelloWorldMethod_1(serveMux *http.ServeMux, pattern string) {
	serveMux.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SimplePrintRequest(r)
		fmt.Fprintln(w, "hello world!")
	}))
}

// RegisterHelloWorldMethod_2 第二种注册一个hello接口的方法
//
//	@param pattern 接口url
func RegisterHelloWorldMethod_2(serveMux *http.ServeMux, pattern string) {
	serveMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		SimplePrintRequest(r)
		fmt.Fprintln(w, "hello world!")
	})
}

// RegisterStatic 注册一个静态文件夹服务
//
//	@param serverMux
//	@param pattern
//	@param staticDir
func RegisterStatic(serverMux *http.ServeMux, pattern string, staticDir string) {
	serverMux.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(staticDir))))
}

// RegisterStaticWithLog 注册一个静态文件夹服务,会打印访问日志
// 因为做了去前缀的逻辑，所以pattern最好由斜杠结尾
//
//	@param serverMux
//	@param pattern "/gatewayx/"
//	@param staticDir "gatewayx/"
func RegisterStaticWithLog(serverMux *http.ServeMux, pattern string, staticDir string) {
	fs := http.FileServer(http.Dir(staticDir))
	serverMux.Handle(pattern, http.StripPrefix(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SimplePrintRequest(r)
		fs.ServeHTTP(w, r)
	})))
}

// SimplePrintRequest 简单日志输出
//
//	@param r 请求对象
func SimplePrintRequest(r *http.Request) {
	log.Printf("RemoteAddr: %s, Method: %s, URL: %s\n", r.RemoteAddr, r.Method, r.URL.Path)
}

// SimpleWriteCookie 简单的写入cookie
//
//	@param w http输出对象
//	@param key cookie key
//	@param value cookie value
func SimpleWriteCookie(w http.ResponseWriter, key string, value string) {
	ck := http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(value),
		HttpOnly: false,
	}
	w.Header().Add("Set-Cookie", ck.String())
}

// RegisterSingleFile 注册一个页面，这个页面返回固定的文件
//
//	@param serverMux
//	@param pattern
//	@param singleFile
func RegisterSingleFile(serverMux *http.ServeMux, pattern string, singleFile string) {
	serverMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(singleFile)
		if err != nil {
			log.Printf("singleFile: %s is not found\n", singleFile)
			ResponseInternalServerError(w)
			return
		}
		if strings.HasSuffix(singleFile, ".html") {
			w.Header().Add("Content-Type", "text/html")
		}
		w.Write(data)
	})
}

// ResponseInternalServerError 返回服务器内部错误
//
//	@param w
func ResponseInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

// RegisterReverseProxy 注册一个反向代理接口
//
//	@param serverMux
//	@param pattern
//	@param reverseServer
func RegisterReverseProxy(serverMux *http.ServeMux, pattern string, reverseServer string) {
	targetURL, _ := url.Parse(reverseServer)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	serverMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		SimplePrintRequest(r)
		proxy.ServeHTTP(w, r)
	})
}

// RegisterReverseProxySkipTLSVerify 注册的一个反向代理接口，忽略https证书校验，返回给前端的请求头增加 "X-Frame-Options", "SAMEORIGIN"
//
//	@param serverMux
//	@param pattern
//	@param reverseServer
func RegisterReverseProxySkipTLSVerify(serverMux *http.ServeMux, pattern string, reverseServer string) {
	targetURL, _ := url.Parse(reverseServer)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	proxy.ModifyResponse = func(r *http.Response) error {
		r.Header.Del("X-Frame-Options")
		r.Header.Add("X-Frame-Options", "SAMEORIGIN")
		return nil
	}
	serverMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		SimplePrintRequest(r)
		proxy.ServeHTTP(w, r)
	})
}

// ParseCookie 提取cookie为map
//
//	@param r
//	@return map
func ParseCookie(r *http.Request) map[string]*http.Cookie {
	cookies := r.Cookies()
	var _map = make(map[string]*http.Cookie)
	for _, cookie := range cookies {
		_map[cookie.Name] = cookie
	}
	return _map
}

// ExtractCookie 提取cookie
func ExtractCookie(r *http.Request, name string) (*http.Cookie, bool) {
	cookie, err := r.Cookie(name)
	if err == http.ErrNoCookie {
		return nil, false
	}
	return cookie, true
}

// RequestRemoveCookie 本函数用于在转发的请求中，移除掉指定的cookie
//
//	@param r
//	@param removeCks
//	@return *http.Request
func RequestRemoveCookie(r *http.Request, removeCks []string) *http.Request {
	newReq := r.Clone(context.Background())
	allCookies := ParseCookie(newReq)
	newReq.Header.Del("Cookie")
	for _, rmCk := range removeCks {
		delete(allCookies, rmCk)
	}
	for _, ck := range allCookies {
		newReq.AddCookie(ck)
	}
	return newReq
}
