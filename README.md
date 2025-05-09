# mem-db


Mem-DB is an in memory database, inspired from AWS Dynamo DB and Redis, written in Go. The objective of this project is simple, we create a database that is able to:

- Store multiple tables, with each tables containing items. An item is simply a record in the table.
- Each item is made up of attributes, which is simply a key-val pair.
- The database operations should prioritize faster reads, and should be able to index on the attributes.

**Note** This project is my attempt at learning Go by creating something that's always intrigued me. How do database indexes work? More generally, can we build a very simple database similar to DynamoDB?

## Getting Started

To get started, simply clone the repository in your local machine, and do:
1. Make sure you have Go installed. Here's a [guide](https://go.dev/dl/) on how to install it locally.
2. Start the database server:
    ```bash
    go run ./cmd/mem-db-api
    ```
3. Get started by inserting a table:
    ```bash
    curl --location 'localhost:8080/tables' \
    --header 'Content-Type: application/json' \
    --data '{
        "id": "12345",
        "name": "first_table"
    }'
    ```
4. Insert items into the table:
    ```bash
    curl --location 'localhost:8080/item' \
    --header 'Content-Type: application/json' \
    --data '{
        "name" : "first_table",
        "attributes": {
            "name": "Farhan",
            "message": "Farhan says hello world!"
        }
    }'
    ```
5. Query Table items:
    ```bash
    curl --location 'localhost:8080/items?table_name=first_table'
    ```

## Roadmap
- [x] Define the database data models and methods.
- [ ] Finish up all the basic db CRUD operations.
- [ ] Implement proper data integrity. 
- [ ] Write the data structure and algos for Attribute level indexes. Implement an AVL Tree.
- [ ] Integrate the index into `GetItems()` method.
- [ ] Use go routines to spin up the server, and implement proper logging.