package db

import (
	"fmt"
)

type DB struct {
	Tables      []Table
	DBIndex     map[string]int
	DBNameIndex map[string]int
}

type Table struct {
	Id    string
	Name  string
	Items []Item
}

type Item struct {
	Key       string
	Attribute []Attribute
	Ttl       int
}

type Attribute struct {
	Name  string
	Value any
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
	index, ok := db.DBIndex[Name]
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
