package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// RespMsg : http响应数据的通用结构
type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResp struct {
	Code      int         `json:"code"`
	Page      int64       `json:"page"`      //当前页
	Pagesize  int64       `json:"pagesize"`  //每页条数
	Total     int64       `json:"total"`     //总条数
	PageCount int64       `json:"pagecount"` //总页数
	Data      interface{} `json:"data"`
}

// NewRespMsg : 生成response对象
func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSONBytes : 对象转json格式的二进制数组
func (resp *RespMsg) JSONBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}

// JSONString : 对象转json格式的string
func (resp *RespMsg) JSONString() string {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

// GenSimpleRespStream : 只包含code和message的响应体([]byte)
func GenSimpleRespStream(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
}

// GenSimpleRespString : 只包含code和message的响应体(string)
func GenSimpleRespString(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg)
}

func SuccessResult(msg string, data interface{}) (httpcode int, contentType string, respdata []byte) {
	resp := NewRespMsg(200, msg, data)
	return http.StatusOK, "application/json", resp.JSONBytes()
}

func ErrorResult(msg string) (httpcode int, contentType string, respdata []byte) {
	resp := NewRespMsg(-1, msg, nil)
	return http.StatusOK, "application/json", resp.JSONBytes()
}

func LoginFailureResult() (httpcode int, contentType string, respdata []byte) {
	resp := NewRespMsg(-2, "登录失效,请重新登录", nil)
	return http.StatusOK, "application/json", resp.JSONBytes()
}

func PageResult(page, pagesize, total, pagecount int64, data interface{}) (httpcode int, contentType string, respdata []byte) {
	resp := PageResp{
		Code:      200,
		Page:      page,
		Pagesize:  pagesize,
		Total:     total,
		PageCount: pagecount,
		Data:      data,
	}
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return http.StatusOK, "application/json", r
}

// 测试提交
