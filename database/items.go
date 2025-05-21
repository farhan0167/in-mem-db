package database

type Item struct {
	Key       string
	Attribute []Attribute
	Ttl       int
}

type Attribute struct {
	Name  string
	Value any
}
