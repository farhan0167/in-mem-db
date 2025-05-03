package database

import (
	"fmt"
	"slices"
)

type DB struct {
	Tables      []Table
	DBIndex     Collections
	DBNameIndex Collections
}

func (db *DB) GetTables() []Table {
	return db.Tables
}
func (db *DB) GetTableById(Id string) (Table, error) {
	//index, ok := db.DBIndex[Id]
	index, err := db.DBIndex.Search(Id)
	if err != nil {
		return Table{}, fmt.Errorf("table with id %s not found", Id)
	}
	return db.Tables[index], nil
}

func (db *DB) GetTableByName(Name string) (Table, error) {
	index, err := db.DBNameIndex.Search(Name)
	if err != nil {
		return Table{}, fmt.Errorf("table with name %s not found", Name)
	}
	return db.Tables[index], nil
}

func (db *DB) AddTable(table Table) error {
	if db.DBIndex.Index == nil {
		db.DBIndex.Init()
	}
	if db.DBNameIndex.Index == nil {
		db.DBNameIndex.Init()
	}

	if table.Id == "" {
		return fmt.Errorf("Table id is empty. Please set a table id.")
	}

	_, err := db.DBIndex.Search(table.Id)
	if err != nil {
		fmt.Errorf("Table with Id %v already exists", table.Id)
	}

	db.Tables = append(db.Tables, table)

	db.DBIndex.Add(table.Id, len(db.Tables)-1)
	db.DBNameIndex.Add(table.Name, len(db.Tables)-1)

	return nil
}

func (db *DB) DeleteTable(id string) error {
	table_index, err := db.DBIndex.Search(id)
	if err != nil {
		return fmt.Errorf("No Table with id %v found", id)
	}
	table := db.Tables[table_index]
	// delete(db.DBIndex, id)
	// delete(db.DBNameIndex, table.Name)

	db.DBIndex.Delete(id)
	db.DBNameIndex.Delete(table.Name)

	db.Tables = slices.Delete(db.Tables, table_index, table_index+1)
	db.DBIndex.Build(db.Tables, "Id")
	db.DBNameIndex.Build(db.Tables, "Name")

	return nil
}

type Table struct {
	Id    string
	Name  string
	Items []Item
	Index map[string]any
}

func (t *Table) GetItems() []Item {
	return t.Items
}

func (t *Table) GetItemByKey(k string) (Item, error) {
	return Item{}, nil
}

func (t *Table) AddItem(item Item) {
	// If index is empty, create a new one
	if t.Index == nil {
		t.Index = make(map[string]any)
	}
	t.Items = append(t.Items, item)
	t.Index[item.Key] = len(t.Items) - 1
}
