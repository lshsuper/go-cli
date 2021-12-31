package database

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

type SqlServerDatabase struct {
	db  *gorm.DB
	cfg DbConfig
}

func newSqlServerDatabase(opt DbConfig) *MysqlDatabase {
	conn, err := gorm.Open(sqlserver.Open(fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s",
		opt.Host,
		opt.Port,
		opt.DefaultDb,
		opt.UserName,
		opt.Password)), &gorm.Config{})
	if err != nil {
		panic("数据库初始异常")
	}

	db, _ := conn.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)

	return &MysqlDatabase{db: conn, cfg: opt}
}

//GetConn
func (m *SqlServerDatabase) GetConn() *gorm.DB {
	return m.db
}

func (m *SqlServerDatabase) GetTables(dbName string) (tbs []TableInfo) {
	m.db.Raw(`SELECT name as 'table_name' FROM SysObjects Where XType='U'`, dbName).Scan(&tbs)
	return
}
func (m *SqlServerDatabase) GetColumns(tbName string) (columns []ColumnInfo) {
	m.db.Raw("SELECT data_type as 'type',column_name as 'field' FROM INFORMATION_SCHEMA.columns WHERE TABLE_NAME=?", tbName).Scan(&columns)
	return
}
