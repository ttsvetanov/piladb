package pila

import (
	"reflect"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := NewDatabase("test-1")

	if db.ID.String() != "2b87e5d8b7d3d853514c8d0801fbcf46" {
		t.Errorf("db.ID is %v, expected %v", db.ID, "2b87e5d8b7d3d853514c8d0801fbcf46")
	}
	if db.Name != "test-1" {
		t.Errorf("db.Name is %v, expected %v", db.Name, "test-1")
	}
	if db.Pila != nil {
		t.Error("db.Pila is not nil")
	}
}

func TestDatabaseCreateStack(t *testing.T) {
	db := NewDatabase("test-db")
	id := db.CreateStack("test-stack")

	if id == nil {
		t.Fatal("stack ID is nil")
	}

	stack, ok := db.Stacks[id]
	if !ok {
		t.Fatal("stack not found in database")
	}
	if stack.ID != id {
		t.Errorf("stack ID is %v, expected %v", stack.ID, id)
	}

	if !reflect.DeepEqual(stack.Database, db) {
		t.Errorf("stack Database is %v, expected %v", stack.Database, db)
	}
}