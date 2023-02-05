package utils

import (
	"bytes"
	"dcr-gin/app/global"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

// CurlPost post请求
// reqMap 可以是结构体 或 map    返回 map[string]any
func CurlPost(url string, reqMap interface{}) map[string]interface{} {
	// body体：序列化为 json str
	body, err := json.Marshal(reqMap)
	if err != nil {
		loggerStr := fmt.Sprintf("url:%+v,reqMap:%s", url, reqMap)
		global.Logger.Info("post请求参数解析错误", zap.String("http", loggerStr))
		fmt.Println("curlPost json.Marshal err:", err)
		return nil
	}
	// 请求方法 url 以及 body体
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		loggerStr := fmt.Sprintf("url:%+v,reqMap:%s,req:%+v", url, reqMap, req)
		global.Logger.Error("post请求错误", zap.String("http", loggerStr))
		fmt.Println("curlPost http.NewRequest err:", err)
		return nil
	}
	// 设置header头
	req.Header.Set("Content-Type", "application/json")

	//设置超时时间
	client := &http.Client{Timeout: 5 * time.Second} // 设置请求超时时长5s
	// 发起请求拿到 返回
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Error("Post请求错误", zap.String("http", err.Error()))
		fmt.Println("Req http.DefaultClient.Do() err: ", err)
		return nil
	}
	// defer 关闭连接
	defer resp.Body.Close()

	// 从返回拿到返回body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error("curlPost ioutil.ReadAll() err", zap.String("http1", err.Error()))
		fmt.Println("curlPost ioutil.ReadAll() err: ", err)
		return nil
	}
	fmt.Println("respBody: ", string(respBody))

	rspMap := make(map[string]interface{})
	// 将body 字符串json 反序列化为 map
	err = json.Unmarshal(respBody, &rspMap)
	if err != nil {
		global.Logger.Error("json解析错误", zap.Error(err))
		return nil
	}
	//global.Logger.Info(fmt.Sprintf("success%+v", rspMap))
	return rspMap
}

type Rsp struct {
	Code int    `json:"code"`
	Rsp  string `json:"rsp"`
}

// CurlGet get请求
// reqMap 可以是结构体 或 map    返回 map[string]any
func CurlGet(url string, reqMap map[string]interface{}) Rsp {
	// 请求参数：序列化为 json str
	body, err := json.Marshal(reqMap)
	if err != nil {
		loggerStr := fmt.Sprintf("url:%+v,reqMap:%s", url, reqMap)
		global.Logger.Info("get请求参数解析错误", zap.String("http", loggerStr))
		fmt.Println("CurlGet json.Marshal err:", err)
		body = nil
	}
	// 请求方法 url 以及 body体
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("TestGetReq http.NewRequest err:", err)
		return Rsp{}
	}
	//设置超时时间
	client := &http.Client{Timeout: 5 * time.Second} // 设置请求超时时长5s
	// 发起请求拿到 返回
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("TestGetReq http.DefaultClient.Do() err: ", err)
		return Rsp{}
	}
	// defer 关闭连接
	defer resp.Body.Close()
	// 拿返回body体
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("TestGetReq ioutil.ReadAll() err: ", err)
		return Rsp{}
	}
	// fmt.Println("respBody: ", string(respBody))
	var rsp Rsp
	// 将body 字符串json 反序列化为 map
	err = json.Unmarshal(respBody, &rsp)
	if err != nil {
		fmt.Println("TestGetReq json.Unmarshal() err: ", err)
		return Rsp{}
	}
	return rsp
}
