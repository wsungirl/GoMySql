package db

import (
	"fmt"
	"strings"

	"log"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetDatabaseTables(dbMod *model.Database) (tables []model.DBTable, err error) {
	var tableNames []string

	rows, err := db.Raw("SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?", dbMod.GetStoredName()).Rows()
	defer rows.Close()

	for rows.Next() {
		var col string
		rows.Scan(&col)
		tableNames = append(tableNames, col)
	}

	log.Println(tableNames)

	var cols []model.TableColumn
	var table model.DBTable

	for _, t := range tableNames {

		err = db.Raw(`
			SELECT
				COLUMN_NAME as field,
				COLUMN_TYPE as type
			FROM information_schema.COLUMNS
			WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`,
			dbMod.GetStoredName(), t,
		).Scan(&cols).Error
		if err != nil {
			err = fmt.Errorf("Can't parse columns: %v", err)
			return
		}

		table.DB = dbMod
		table.Name = t
		table.Columns = cols

		tables = append(tables, table)
	}

	return
}

func (db *DB) CreateTable(table *model.DBTable) (err error) {

	var colStrings []string

	for _, c := range table.Columns {
		colStrings = append(colStrings, c.Field+" "+c.Type)
	}

	query := fmt.Sprintf(
		`CREATE TABLE %s.%s (%s)`,
		table.DB.GetStoredName(), table.Name,
		strings.Join(colStrings, ","),
	)

	err = db.Exec(query).Error
	if err != nil {
		return fmt.Errorf("Can't create table: %v", err)
	}

	return
}

func (db *DB) AddRowToTable (newRow * model.TableRow) (err error) {
	columnValuesPattern := strings.Repeat(",?", len(newRow.Table.Columns)-1)
	var columnNames string
	for i :=0; i < len(newRow.Table.Columns); i++ {
		columnNames += newRow.Table.Columns[i].Field+","
	}
	columnNames = columnNames[:len(columnNames)-1]
	fullTableName := newRow.Table.DB.Name+"."+newRow.Table.Name
	query := "INSERT INTO "+fullTableName+"("+columnNames+") VALUES (?"+columnValuesPattern+");"
	err = db.Exec(query, newRow.Values...).Error
	if err != nil {
		return fmt.Errorf("can`t insert values into table %s", fullTableName)
	}
	return nil
}
