package database

import (
	"fmt"
	"reflect"
)

type Index interface {
	Add(k string, v any) error
	Delete(k string)
	Search(k string) (int, error)
	Init()
	Build(structs any, fieldName string) error
}

type CollectionsIndex struct {
	Index map[string]int
}

func (c *CollectionsIndex) Init() {
	c.Index = make(map[string]int)
}

func (c *CollectionsIndex) Add(k string, v any) error {
	index, ok := v.(int)
	if ok {
		c.Index[k] = index
		return nil
	}
	return fmt.Errorf("v is of type %T. int expected", v)
}

func (c *CollectionsIndex) Search(k string) (int, error) {
	index, ok := c.Index[k]
	if ok {
		return index, nil
	}
	return -1, fmt.Errorf("key %v does not exist", k)
}

func (c *CollectionsIndex) Delete(k string) {
	delete(c.Index, k)
}

func (c *CollectionsIndex) Build(structs any, fieldName string) error {
	iterable := reflect.ValueOf(structs)

	if iterable.Kind() != reflect.Slice {
		return fmt.Errorf("structs is of type %T. slice expected", structs)
	}

	for i := 0; i < iterable.Len(); i++ {
		structValue := iterable.Index(i)
		fieldValue := structValue.FieldByName(fieldName)
		c.Add(fieldValue.String(), i)
	}

	return nil
}
