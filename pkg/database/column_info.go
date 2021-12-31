package database

type ColumnInfo struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Null    string `json:"null"`
	Key     string `json:"key"`
	Default string `json:"default"`
	Extra   string `json:"extra"`
}
