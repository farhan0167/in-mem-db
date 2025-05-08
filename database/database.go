package database

import (
	"fmt"
	"slices"
)

const (
	ITEM_INDEX = "ItemsIndex"
)

type DB struct {
	Tables      []Table
	DBIndex     CollectionsIndex
	DBNameIndex CollectionsIndex
}

func (db *DB) GetTables() []Table {
	return db.Tables
}
func (db *DB) GetTableById(Id string) (*Table, error) {
	index, err := db.DBIndex.Search(Id)
	if err != nil {
		return &Table{}, fmt.Errorf("table with id %s not found", Id)
	}
	return &db.Tables[index], nil
}

func (db *DB) GetTableByName(Name string) (*Table, error) {
	index, err := db.DBNameIndex.Search(Name)
	if err != nil {
		return &Table{}, fmt.Errorf("table with name %s not found", Name)
	}
	return &db.Tables[index], nil
}

func (db *DB) AddTable(table Table) error {
	if db.DBIndex.Index == nil {
		db.DBIndex.Init()
	}
	if db.DBNameIndex.Index == nil {
		db.DBNameIndex.Init()
	}

	if table.Id == "" {
		return fmt.Errorf("table id is empty. Please set a table id")
	}

	_, err := db.DBIndex.Search(table.Id)
	if err == nil {
		return fmt.Errorf("table with Id %v already exists", table.Id)
	}

	// Initialize Table's Items Index
	table.index = make(map[string]Index)

	// Initialize a basic Items Index
	table.index[ITEM_INDEX] = &CollectionsIndex{
		Index: make(map[string]int),
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

	db.DBIndex.Delete(id)
	db.DBNameIndex.Delete(table.Name)

	db.Tables = slices.Delete(db.Tables, table_index, table_index+1)
	db.DBIndex.Build(db.Tables, "Id")
	db.DBNameIndex.Build(db.Tables, "Name")

	return nil
}

type Table struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
	index map[string]Index
}

func (t *Table) GetItems() []Item {
	return t.Items
}

func (t *Table) GetItemByKey(k string) (Item, error) {
	index, err := t.index[ITEM_INDEX].Search(k)
	if err != nil {
		return Item{}, fmt.Errorf("Item with key %v not found", k)
	}
	return t.Items[index], nil
}

func (t *Table) AddItem(item Item) error {
	_, err := t.index[ITEM_INDEX].Search(item.Key)
	if err == nil {
		return fmt.Errorf("Item with key %v already exists", item.Key)
	}
	t.Items = append(t.Items, item)
	t.index[ITEM_INDEX].Add(item.Key, len(t.Items)-1)
	return nil
}
