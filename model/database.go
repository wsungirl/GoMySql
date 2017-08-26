package model

type IDatabase interface {
	GetDBInfo(idDb int64) (*Database, error)
	GetDBList(idUser int64) ([]Database, error)
	GetTableInfo(idUser int64, idDb int64, tableName string) (*DbTable, error)
	GetTableRows(idUser int64, idDb int64, tableName string, offset int64, limit int64, fields string) ([]string, error)
	GetTableRow(idUser int64, idDb int64, tableName string, rowId int64) (string, error)
	CreateDB(sDb *Database) error
	CreateTable(sDbT *DbTable) error
	//AddRow ()
	//DeleteDb ()
	//DeleteTable ()
	//DeleteRow ()
	//UpdateRow ()
	//UpdateRows ()
}

type TableColumn struct {
	ColumnType   string
	ColumnName   string
	//ColumnNotNull bool
	//ColumnDefault string
	//ColumnUnique bool
}

type DbTable struct {
	IdUser        int64
	IdDb          int64
	TableName     string
	Columns []TableColumn
}

type Database struct {
	ID     int64  `json:"id,omitempty"`
	Uid    int64  `json:"user_id,omitempty"`
	DBname string `json:"db_name,omitempty"`
}
