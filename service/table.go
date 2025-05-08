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

func DeleteTable(db *database.DB, id string) error {
	return db.DeleteTable(id)
}

type Attributes map[string]any

func GetItems(db *database.DB, table_name string) ([]Attributes, error) {
	res := make([]Attributes, 0)
	table, err := db.GetTableByName(table_name)
	if err != nil {
		return res, fmt.Errorf("%v", err)
	}

	items := table.GetItems()
	for _, item := range items {
		attr := make(Attributes)
		attr["key"] = item.Key
		for _, a := range item.Attribute {
			attr[a.Name] = a.Value
		}
		res = append(res, attr)
	}
	return res, nil
}
