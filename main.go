package main

import (
	"farhan0167/mem-db/database"
	"fmt"
)

func main() {
	db := database.DB{}

	err := db.AddTable(database.Table{
		Id:   "123",
		Name: "users",
	})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AddTable(database.Table{
		Id:   "124",
		Name: "models",
	})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AddTable(database.Table{
		Id:   "125",
		Name: "clients",
	})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AddTable(database.Table{
		Id:   "126",
		Name: "tokens",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Tables: %v \n", db.GetTables())
	fmt.Println(db.GetTableByName("clients"))
	err = db.DeleteTable("124")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db.GetTableByName("clients"))
	// table, err := database.GetTable("users")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(table)
	// }
}
