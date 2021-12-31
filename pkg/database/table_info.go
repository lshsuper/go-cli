package database

type TableInfo struct {
	TableName string `gorm:"column:table_name"`
}
