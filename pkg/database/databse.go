package database

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	GetConn() *gorm.DB
	GetTables(dbName string) []TableInfo
	GetColumns(tbName string) []ColumnInfo
}

//Register
func Register(dbType DbType, cfg DbConfig) IDatabase {

	switch dbType {

	case Mysql:
		return newMysqlDatabase(cfg)
	case Sqlserver:
		return newSqlServerDatabase(cfg)
	default:
		panic("不存在该数据库类型")

	}

}
