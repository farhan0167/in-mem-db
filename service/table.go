package service

import (
	"farhan0167/mem-db/database"
	"fmt"
)

type Table struct {
	Id            string
	Name          string
	NumberOfItems int
}

func GetTables(db *database.DB) []Table {
	tables := db.GetTables()
	table_res := make([]Table, 0)

	for _, table := range tables {
		var count int
		t := Table{
			Id:   table.Id,
			Name: table.Name,
		}
		for range table.Items {
			count++
		}
		t.NumberOfItems = count
		table_res = append(table_res, t)
	}
	return table_res
}

type GetTableParams struct {
	Id   string
	Name string
}

func GetTable(db *database.DB, options GetTableParams) (Table, error) {
	t := Table{}
	if options.Id != "" {
		table, err := db.GetTableById(options.Id)
		if err != nil {
			return t, fmt.Errorf("%v", err)
		}
		t.Id, t.Name = table.Id, table.Name

	} else {
		table, err := db.GetTableByName(options.Name)
		if err != nil {
			return t, fmt.Errorf("%v", err)
		}
		t.Id, t.Name = table.Id, table.Name
	}
	return t, nil
}

func AddTable(db *database.DB, table database.Table) error {
	return db.AddTable(table)
}
