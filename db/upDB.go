package db

import (
	"database/sql"
	"fmt"
)

//var createTables = []string{
//	CreateTableQuery,
//}

//func Up() error {
//	for i, table := range createTables {
//		_, err := GetDBConn().Exec(table)
//		if err != nil {
//			return errors.New(
//				fmt.Sprintf("error occurred while creating table №%d, error is: %s", i, err.Error()))
//		}
//	}
//	return nil
//}

func Up(db *sql.DB) error {
	for i, table := range createTables {
		_, err := db.Exec(table)
		if err != nil {
			return fmt.Errorf("error occurred while creating table №%d, error is: %w", i, err)
		}
	}
	return nil
}
