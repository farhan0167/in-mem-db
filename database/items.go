package database

type Item struct {
	Key       string      `json:"key"`
	Attribute []Attribute `json:"attribute"`
	Ttl       int         `json:"ttl"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}
