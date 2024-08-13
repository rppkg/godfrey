package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rs/xid"
)

func main() {
	// 创建一个默认的命令行标志集
	flagSet := flag.NewFlagSet("xid", flag.ExitOnError)
	help := flagSet.Bool("h", false, "显示帮助信息")
	flagSet.BoolVar(help, "help", false, "显示帮助信息")

	// 解析命令行参数
	_ = flagSet.Parse(os.Args[1:])

	if *help || len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "new":
		handleNewCommand(os.Args[2:])
	default:
		fmt.Printf("未知的子命令: %s\n", subcommand)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("使用方法:")
	fmt.Println("  xid <子命令> [参数]")
	fmt.Println("可用的子命令:")
	fmt.Println("  new    生成新的ID")
	fmt.Println("		-count=[生成xid数量]")
	fmt.Println("全局标志:")
	fmt.Println("  -h, --help    显示帮助信息")
}

func handleNewCommand(args []string) {
	// 子命令参数的flag.FlagSet
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	// 定义标志
	count := newCmd.Int("count", 1, "生成数量")
	// 解析子命令的标志
	if err := newCmd.Parse(args); err != nil {
		fmt.Println("解析子命令参数失败")
		os.Exit(1)
	}

	if count != nil {
		for i := 0; i < *count; i++ {
			fmt.Println(xid.New())
		}
	}
}
