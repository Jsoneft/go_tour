package cmd

import (
	"github.com/spf13/cobra"
	"go_tour/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:                        "time",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "时间格式处理",
	Long:                       "",
	Example:                    "",
	ValidArgs:                  nil,
	ValidArgsFunction:          nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Hidden:                     false,
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run: func(cmd *cobra.Command, args []string) {

	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

var nowTimeCmd = &cobra.Command{
	Use:                        "now",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "获取当前时间",
	Long:                       "获取当前时间",
	Example:                    "",
	ValidArgs:                  nil,
	ValidArgsFunction:          nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Hidden:                     false,
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"),nowTime.Unix())
	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

var calculateTimeCmd = &cobra.Command{
	Use:                        "calc",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "计算所需要的时间",
	Long:                       "计算所需要的时间",
	Example:                    "",
	ValidArgs:                  nil,
	ValidArgsFunction:          nil,
	Args:                       nil,
	ArgAliases:                 nil,
	BashCompletionFunction:     "",
	Deprecated:                 "",
	Hidden:                     false,
	Annotations:                nil,
	Version:                    "",
	PersistentPreRun:           nil,
	PersistentPreRunE:          nil,
	PreRun:                     nil,
	PreRunE:                    nil,
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime ==""{
			currentTimer = timer.GetNowTime()
		}else {
			var err  error
			if !strings.Contains(calculateTime, " "){
				layout = "2006-01-02"
			}
			currentTimer,err = time.Parse(layout, calculateTime)
			if err != nil {
				t,_:=strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		 calculateTime,err := timer.CalculateTime(currentTimer, duration)

		if err != nil {
			log.Fatalf("timer.CalculateTime err : %v",err)
		}
		log.Printf("输出结果 : %s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
	TraverseChildren:           false,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位或为时间戳或者是格式化的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间， 有效单位为"ns","us","ms", "s", "m", "h"`)
}