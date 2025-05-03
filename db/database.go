package db

type Table struct {
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
