package cmd

import (
	"github.com/spf13/cobra"
	"go_tour/internal/word"
	"log"
	"strings"
)

const (
	MODE_UPPER                    = iota + 1 // 转化为大写
	MODE_LOWER                               // 转化为小写
	MODE_UNDERSCORE_TO_UPPER_CAME            // 下划线转大写驼峰
	MODE_UNDERSCORE_TO_LOWER_CAEE            // 下划线转小写驼峰
	MODE_CAME_TO_UNDERSCORE                  // 驼峰转下划线
)

var desc = strings.Join([]string{
	"该子命令支持如下转换，模式如下：",
	"1: 全部单词转化成大写",
	"2: 全部单词转化成小写",
	"3: 全部单词下划线转大写驼峰",
	"4: 全部单词下划线转小写驼峰",
	"5: 全部单词驼峰转下划线",
}, "\n")

var mode int8
var str string
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAME:
			content = word.UnderscoreToUpperCame(str)
		case MODE_UNDERSCORE_TO_LOWER_CAEE:
			content = word.UnderscoreToLowerCame(str)
		case MODE_CAME_TO_UNDERSCORE:
			content = word.CameToUnderscore(str)
		default:
			log.Fatal(" not supported, execute 'help word' for help")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "input")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "转换模式")

}
