package db

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

var ErrUnrecognisedTable error = fmt.Errorf("unrecognised table name")
var ErrNotFound error = fmt.Errorf("record not found")
var ErrRecordAlreadyExists error = fmt.Errorf("record already exists")

/*
	Fake database implements the database client.
	For the purpose of this test, there is no real database connection; instead we're saving objects to a map
*/
type FakeDatabase struct {
	// represents a map of tables
	database map[string]map[string]interface{}

	// required if the service was to include additional go routines for data processing.
	sync.Mutex
}

func NewFakeDatabase() *FakeDatabase {
	// we only have the one port table
	portsTable := make(map[string]interface{}, 0)
	database := map[string]map[string]interface{}{"port": portsTable}
	return &FakeDatabase{database: database}
}

func (fdb *FakeDatabase) Count(tableName string) int {
	defer fdb.Unlock()
	fdb.Lock()
	return len(fdb.database[tableName])
}

func (fdb *FakeDatabase) Get(id string, tableName string) interface{} {
	defer fdb.Unlock()
	fdb.Lock()

	if table, err := fdb.getTable(tableName); err != nil {
		return err
	} else {
		return table[id]
	}
}

func (fdb *FakeDatabase) Create(id string, object interface{}) error {
	defer fdb.Unlock()
	fdb.Lock()

	return fdb.queryTable(id, object, func(id string, object interface{}, tableName string) error {
		if record := fdb.database[tableName][id]; record == nil {
			fdb.database[tableName][id] = object
			return nil
		}
		return ErrRecordAlreadyExists
	})
}

func (fdb *FakeDatabase) Update(id string, object interface{}, changes map[string]interface{}) error {
	defer fdb.Unlock()
	fdb.Lock()

	return fdb.queryTable(id, object, func(id string, object interface{}, tableName string) error {
		if record := fdb.database[tableName][id]; record == nil {
			return ErrNotFound
		} else {
			return fdb.updateObject(object, changes)
		}
	})
}

func (fdb *FakeDatabase) Delete(id string, tableName string) error {
	defer fdb.Unlock()
	fdb.Lock()

	if table, err := fdb.getTable(tableName); err != nil {
		return err
	} else {
		if row := table[id]; row == nil {
			return ErrNotFound
		}
		delete(fdb.database[tableName], id)
		return nil
	}
}

func (fdb *FakeDatabase) updateObject(dbObject interface{}, changes map[string]interface{}) error {
	v := reflect.ValueOf(dbObject).Elem()
	t := reflect.TypeOf(dbObject).Elem()

	// iterate over the fields on the object and update those that have been included in the changes map
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		fieldChanges := changes[strings.ToLower(f.Name)]

		// if the field is not to be changed, move onto the next
		if fieldChanges != nil {
			fieldVal := v.FieldByName(f.Name)

			switch fieldVal.Kind() {
			case reflect.Ptr:
				rv := reflect.New(reflect.TypeOf(fieldChanges))
				rv.Elem().Set(reflect.ValueOf(fieldChanges))

				// TODO: add handling for pointer time.Location field and structs
				fieldVal.Set(rv)
			default:
				switch v.Field(i).Kind() {
				case reflect.Struct:
					if _, ok := v.Field(i).Interface().(time.Location); ok {
						loc, err := parseLocation(fieldChanges)
						if err != nil {
							return err
						}

						v.Field(i).Set(reflect.ValueOf(*loc))
					} else {
						v.Field(i).Set(reflect.ValueOf(fieldChanges))
					}

				default:
					v.Field(i).Set(reflect.ValueOf(fieldChanges))
				}
			}

		}

	}
	return nil
}

func parseLocation(fieldChanges interface{}) (*time.Location, error) {
	locStrChange := fieldChanges.(string)
	return time.LoadLocation(locStrChange)
}

func (fdb *FakeDatabase) getTable(tableName string) (map[string]interface{}, error) {
	table := fdb.database[tableName]
	if table == nil {
		return nil, ErrUnrecognisedTable
	}
	return table, nil
}

func (fdb *FakeDatabase) queryTable(id string, object interface{}, fn func(i string, o interface{}, t string) error) error {
	table := strings.ToLower(reflect.TypeOf(object).Elem().Name())
	if fdb.database[table] == nil {
		return ErrUnrecognisedTable
	}

	return fn(id, object, table)
}
