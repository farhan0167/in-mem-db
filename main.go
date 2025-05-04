package main

import (
	"farhan0167/mem-db/database"
	"fmt"
)

func main() {
	db := database.DB{}

	table := database.Table{
		Id:   "123",
		Name: "users",
	}
	//fmt.Println("Table pre db add", table)
	err := db.AddTable(table)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("table post db add", table)
	table_db, _ := db.GetTableByName("users")
	//fmt.Println("table_db returned by GetTableByName", table_db)
	item := database.Item{
		Key: "123",
		Attribute: []database.Attribute{
			{
				Name:  "name",
				Value: "Farhan",
			},
		},
		Ttl: 3600,
	}
	item2 := database.Item{
		Key: "1234",
		Attribute: []database.Attribute{
			{
				Name:  "name",
				Value: "Ishraq",
			},
		},
		Ttl: 3600,
	}
	err = table_db.AddItem(item)
	if err != nil {
		fmt.Println(err)
	}
	err = table_db.AddItem(item2)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(table.GetItems())
	fmt.Println(table_db.GetItemByKey("1234"))

}
