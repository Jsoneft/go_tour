package cmd

import (
	"github.com/spf13/cobra"
	"go_tour/internal/sql2struct"
	"log"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:                    "sql",
	Aliases:                nil,
	SuggestFor:             nil,
	Short:                  "sql process",
	Long:                   "SQL process",
	Example:                "",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
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
var sql2structCmd = &cobra.Command{
	Use:                    "struct",
	Aliases:                nil,
	SuggestFor:             nil,
	Short:                  "sql2struct",
	Long:                   "sql2struct",
	Example:                "",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Hidden:                 false,
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			PassWord: password,
			CharSet:  charset,
		}
		dbModel := sql2struct.NewModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect() err = %v", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns() err = %v", err)
		}
		Tpl := sql2struct.NewStructTemplate()
		structCol := Tpl.AssemblyColumns(columns)
		err = Tpl.Generate(tableName, structCol)
		if err != nil {
			log.Fatalf("Tpl.Generate() err = %v", err)
		}

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
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "数据库HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "数据库 字符串编码")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "数据库 名称")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "数据库 实例类型")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "数据库 表名称")
}
