package database

import (
	"fmt"
	"slices"
)

type DB struct {
	Tables      []Table
	DBIndex     map[string]int
	DBNameIndex map[string]int
}

func (db *DB) GetTables() []Table {
	return db.Tables
}
func (db *DB) GetTableById(Id string) (Table, error) {
	index, ok := db.DBIndex[Id]
	if ok {
		return db.Tables[index], nil
	}
	return Table{}, fmt.Errorf("table with id %s not found", Id)
}

func (db *DB) GetTableByName(Name string) (Table, error) {
	index, ok := db.DBNameIndex[Name]
	if ok {
		return db.Tables[index], nil
	}
	return Table{}, fmt.Errorf("table with name %s not found", Name)
}

func (db *DB) AddTable(table Table) error {
	if db.DBIndex == nil {
		db.DBIndex = make(map[string]int)
	}
	if db.DBNameIndex == nil {
		db.DBNameIndex = make(map[string]int)
	}
	if table.Id == "" {
		return fmt.Errorf("Table id is empty. Please set a table id.")
	}
	_, ok := db.DBIndex[table.Id]
	if ok {
		return fmt.Errorf("Table with Id %v already exists", table.Id)
	}
	db.Tables = append(db.Tables, table)
	db.DBIndex[table.Id] = len(db.Tables) - 1
	db.DBNameIndex[table.Name] = len(db.Tables) - 1

	return nil
}

func (db *DB) DeleteTable(id string) error {
	table_index, ok := db.DBIndex[id]
	if ok != true {
		return fmt.Errorf("No Table with id %v found", id)
	}
	table := db.Tables[table_index]
	delete(db.DBIndex, id)
	delete(db.DBNameIndex, table.Name)

	db.Tables = slices.Delete(db.Tables, table_index, table_index+1)
	for index, table := range db.Tables {
		db.DBIndex[table.Id] = index
		db.DBNameIndex[table.Name] = index
	}

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
