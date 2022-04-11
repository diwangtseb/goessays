package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func SubCommand() {
	rootCmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	firstLevel := cobra.Command{
		Use: "first-level",
	}
	SecondLevel := cobra.Command{
		Use: "second-level",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("second-level")
		},
	}
	firstLevel.AddCommand(&SecondLevel)
	rootCmd.AddCommand(&firstLevel)
	rootCmd.Execute()
}

func DemoCommand() {
	rootCmd := cobra.Command{
		// 我们定义这个工具默认输出帮助信息
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// 定义一个子命令，直接输出命令行参数
	echoCmd := cobra.Command{
		// 子命令描述，第一个单词作为子命令名称
		Use: "echo inputs",

		// 子命令别名，可以使用任何一个别名替代上面的 echo 来运行子命令
		Aliases: []string{
			"copy",
			"repeat",
		},

		// 命令提示，默认情况如果输入的子命令不存在会提示 unknown cmd，
		// 但是如果定义了 SuggestFor 的情况下，如果输入的命令不存在，会去 SuggestFor
		// 里面查找是否有匹配的字符串，如果有，则提示是否期望输入的是 echo 命令
		SuggestFor: []string{
			"ech", "cp", "rp",
		},

		DisableSuggestions:         true,
		SuggestionsMinimumDistance: 2,
		DisableFlagsInUseLine:      true,

		// 简单扼要概括下命令的用途
		Short: "echo is a command to echo command line inputs",

		// 想说什么都在这里说吧，越详细越好，可以使用 `` 来跨行输入
		Long: `echo is a command to echo command line inputs.
It is a very simple command used to display how to implement command line tools
using cobra, which is a very famous library to build command line interface tools.
`,

		// 当你在新版本废弃这个命令的时候，可以先隐藏，让用户优先使用替代品或者看不到，
		// 但是处于向下兼容目的，这个命令仍然是可用的，只是在帮助列表里面看不到
		//Hidden: true,

		// 当你需要废弃这个命令的时候设置。废弃的意思意味着未来版本可能删除这个命令。
		// 标注为废弃的命令在执行的时候，都会打印命令已废弃的提示信息以及这个设置的提示信息。
		//Deprecated: "will be deleted in version 2.0",

		// 注解，用于代码层面的命令分组，不会显示在命令行输出中
		Annotations: map[string]string{
			"group":        "user",
			"require-auth": "none",
		},

		SilenceErrors: true,
		//SilenceUsage:  true,
		// Version 定义版本
		Version: "1.0.0",

		// 是否禁用选项解析
		//DisableFlagParsing: true,

		//DisableAutoGenTag: true,

		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//     fmt.Println("hahha! let me check echk")
		// },

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			//    return errors.New("invalid parameter")
			return nil
		},

		// 子命令执行过程
		// Run: func(cmd *cobra.Command, args []string) {
		//     fmt.Println(strings.Join(args, " "))
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(strings.Join(args, " "))
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("i am post run")
		},
	}

	// 添加子命令根命令
	rootCmd.AddCommand(&echoCmd)

	// 执行根命令
	rootCmd.Execute()
}

func GlobalCommand() {
	var namespace string
	rootCmd := cobra.Command{
		Use: "kubectl",
	}
	// 添加全局的选项，所有的子命令都可以继承
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "If present, the namespace scope for this CLI request")

	// 一级子命令 get
	var outputFormat string
	var labelSelector string
	getCmd := cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print flags...")
			fmt.Printf("Flags: namespace=[%s], selector=[%s], output=[%s]\n", namespace, labelSelector, outputFormat)
			fmt.Println("Print args...")
			for _, arg := range args {
				fmt.Println("Arg:", arg)
			}
		},
	}
	// 添加命令选项，这些选项仅 get 命令可用
	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format")
	getCmd.Flags().StringVarP(&labelSelector, "selector", "l", "", "Selector (label query) to filter on")

	// 组装命令
	rootCmd.AddCommand(&getCmd)

	// 执行命令
	rootCmd.Execute()
}

func main() {
	// DemoCommand()
	// SubCommand()
	GlobalCommand()
}
