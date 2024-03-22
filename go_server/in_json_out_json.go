/**
 * Author: anarckk anarckk@gmail.com
 * Date: 2023-04-30 17:46:47
 * LastEditTime: 2023-04-30 18:12:00
 * Description: 再go_server_demo1中被引用过了，可以再那边看
 * 这个就是一个简单的包装方法，简便的在go server中使用 application/json 和 application/xml
 *
 * Copyright (c) 2023 by anarckk, All Rights Reserved.
 */
package go_server_util

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseMsg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

/*
	 使用示例
	  func main() {
		 http.Handle("/", http.FileServer(http.Dir("static")))
		 http.HandleFunc("/helloworld", OutJson(HandleHelloWorld))
		 http.HandleFunc("/now", OutJson(HandleNow))
		 log.Println("server run in: 35080")
		 err := http.ListenAndServe(":35080", nil)
		 if err != nil {
			 log.Println(err)
		 }
	 }

	 func HandleHelloWorld(ctx context.Context) (int, interface{}, string) {
		 return 0, nil, "helloworld"
	 }

	 func HandleNow(ctx context.Context) (int, interface{}, string) {
		 return 0, map[string]string{"now": strconv.FormatInt(time.Now().UnixMilli(), 10)}, "ok"
	 }
*/
func InXmlOutJson[T interface{}, R interface{}](handler func(ctx context.Context, in T) (int, R, string)) func(http.ResponseWriter, *http.Request) {
	getBody := func(req *http.Request) []byte {
		defer func() {
			err := req.Body.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		log.Printf("request[%s]", string(b))
		return b
	}
	returnM := func(rw http.ResponseWriter, msg *ResponseMsg) {
		rw.Header().Set("Content-Type", "application/json")
		vByte, err := json.Marshal(msg)
		if err != nil {
			write, err := rw.Write([]byte(err.Error()))
			if err != nil {
				log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
			}
			return
		}
		write, err := rw.Write(vByte)
		if err != nil {
			log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
		}
		log.Printf("response[%s]", string(vByte))
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		data := getBody(req)
		var reqXml T
		err := xml.Unmarshal(data, &reqXml)
		if err != nil {
			returnM(rw, &ResponseMsg{-1, nil, err.Error()})
		}

		code, resData, msg := handler(context.Background(), reqXml)
		returnM(rw, &ResponseMsg{code, resData, msg})
	}
}

func OutJson[R interface{}](handler func(ctx context.Context) (int, R, string)) func(http.ResponseWriter, *http.Request) {
	returnM := func(rw http.ResponseWriter, msg *ResponseMsg) {
		rw.Header().Set("Content-Type", "application/json")
		vByte, err := json.Marshal(msg)
		if err != nil {
			write, err := rw.Write([]byte(err.Error()))
			if err != nil {
				log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
			}
			return
		}
		write, err := rw.Write(vByte)
		if err != nil {
			log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
		}
		log.Printf("response[%s]", string(vByte))
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		code, resData, msg := handler(context.Background())
		returnM(rw, &ResponseMsg{code, resData, msg})
	}
}

func InJsonOutJson[T interface{}, R interface{}](handler func(ctx context.Context, in T) (int, R, string)) func(http.ResponseWriter, *http.Request) {
	getBody := func(req *http.Request) []byte {
		defer func() {
			err := req.Body.Close()
			if err != nil {
				log.Println(err)
			}
		}()
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		log.Printf("request[%s]", string(b))
		return b
	}
	returnM := func(rw http.ResponseWriter, msg *ResponseMsg) {
		rw.Header().Set("Content-Type", "application/json")
		vByte, err := json.Marshal(msg)
		if err != nil {
			write, err := rw.Write([]byte(err.Error()))
			if err != nil {
				log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
			}
			return
		}
		write, err := rw.Write(vByte)
		if err != nil {
			log.Printf("send msg error write[%d] bytes, err[%s]", write, err)
		}
		log.Printf("response[%s]", string(vByte))
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		data := getBody(req)
		var reqJson T
		err := json.Unmarshal(data, &reqJson)
		if err != nil {
			returnM(rw, &ResponseMsg{-1, nil, err.Error()})
		}

		code, resData, msg := handler(context.Background(), reqJson)
		returnM(rw, &ResponseMsg{code, resData, msg})
	}
}
