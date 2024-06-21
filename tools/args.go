package tools

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Name   string
	URL    string
	Proxy  string
	Dnslog string
	Lhost  string
	Lport  string
}

// GetCommandOptions 解析命令行参数并填充 Options 结构体
func GetCommandOptions() (*Options, error) {

	opts := &Options{}

	if len(os.Args) <= 2 {
		if len(os.Args) == 2 {
			opts.Name = os.Args[1]
			switch opts.Name {
			case "week":
				printTypeHelp(opts.Name)
			case "unauth":
				printTypeHelp(opts.Name)
			case "token":
				printTypeHelp(opts.Name)
			case "help":
				printHelp(opts.Name)
			default:
				err := fmt.Errorf("invalid operation option: %s", opts.Name)
				printHelp(opts.Name)
				return opts, err
			}
		} else {
			err := errors.New("not enough arguments provided")
			printHelp("")
			return opts, err
		}

	} else {
		// 获取命令行参数列表
		args := os.Args[2:]

		// 检查是否有足够的参数
		if len(args) < 1 {
			err := errors.New("not enough arguments provided")
			printHelp(opts.Name)
			return opts, err
		}

		// 将第一个字符串作为类型标识
		opts.Name = os.Args[1]

		validTypes := []string{"help", "week", "unauth", "token"}
		if !contains(validTypes, opts.Name) {
			return nil, fmt.Errorf("invalid operation option: %s", opts.Name)
		}

		flag.StringVar(&opts.URL, "u", "", "target url")
		flag.StringVar(&opts.Proxy, "p", "", "proxy string eg: socks://127.0.0.1:8080")
		flag.StringVar(&opts.Dnslog, "d", "", "dnslog domain")
		flag.StringVar(&opts.Lhost, "lh", "", "reverse ip")
		flag.StringVar(&opts.Lport, "lp", "", "reverse port")

		helpFlag := flag.Bool("h", false, "help")

		// 解析标志
		flag.CommandLine.Parse(os.Args[2:])

		// 检查是否请求了帮助
		if *helpFlag {
			printHelp(opts.Name)
		}

	}

	return opts, nil
}

func printHelp(typeName string) {
	fmt.Println("Usage: go run xxl-job.go [type] [options]")
	fmt.Println()
	fmt.Println("Available types:")
	for _, t := range []string{"help", "week", "unauth", "token"} {
		fmt.Printf(" - %s\n", t)
	}
	fmt.Println()
	fmt.Println("Options:")

	flag.PrintDefaults()
	if typeName != "" {
		validTypes := []string{"help", "week", "unauth", "token"}
		if !contains(validTypes, typeName) {
			fmt.Printf("  invalid operation option: %s\n", typeName)
			fmt.Println()
		} else {
			if typeName == "help" {

			} else {
				fmt.Printf("\nUsage for type %s:\n", typeName)
				fmt.Println("  go run xxl-job.go", typeName, "[options]")
				fmt.Println()
			}

		}

	} else {
		fmt.Printf("  invalid or missing command line arguments\n")
		fmt.Println()
	}
}

func printTypeHelp(t string) {
	switch t {
	case "week":
		fmt.Printf("Usage for type %s:\n", t)
		fmt.Println("  go run main.go", t, "[-u url] [-p proxy]")
		fmt.Println("Options:")
		fmt.Println("  -u url     input target url.")
		fmt.Println("  -p proxy   input proxy string[Optional parameters].")
		os.Exit(0)
	case "unauth":
		fmt.Printf("Usage for type %s:\n", t)
		fmt.Println("  go run main.go", t, "[-u url] [-d dnslog] [-lh lhost] [-lp lport] [-p proxy]")
		fmt.Println("Options:")
		fmt.Println("  -u url     input target url.")
		fmt.Println("  -p proxy   input proxy string[Optional parameters].")
		fmt.Println("  -d dnslog  input dnslog string.")
		fmt.Println("  -lh lhost  input reverse ip.")
		fmt.Println("  -lp lport  input reverse port.")
		os.Exit(0)
	case "token":
		fmt.Printf("Usage for type %s:\n", t)
		fmt.Println("  go run main.go", t, "[-u url] [-d dnslog] [-lh lhost] [-lp lport] [-p proxy]")
		fmt.Println("Options:")
		fmt.Println("  -u url     input target url.")
		fmt.Println("  -p proxy   input proxy string[Optional parameters].")
		fmt.Println("  -d dnslog  input dnslog string.")
		fmt.Println("  -lh lhost  input reverse ip.")
		fmt.Println("  -lp lport  input reverse port.")
		os.Exit(0)
	}

}

// contains 函数用于检查切片中是否包含某个元素
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
