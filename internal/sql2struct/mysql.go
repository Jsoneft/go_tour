package sql2struct

import (
	"database/sql"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo *DBInfo
}

type DBInfo struct {
	DBType string
	Host string
	UserName string
	PassWord string
	CharSet string
}

type TableColumn struct {
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string
	
}

func NewModel(info *DBInfo) *DBModel{
	return &DBModel{DBInfo: info}
}

func (m *DBModel)Connect()error	 {
	var err error
	dsn := fmt.Sprint(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.PassWord,
		m.DBInfo.Host,
		m.DBInfo.CharSet,
		)

	m.DBEngine ,err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil


}