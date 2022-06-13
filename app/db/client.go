package db

/*
	A database client that describes basic CRUD actions
*/
type Client interface {
	Count(tableName string) int
	Get(id string, tableName string) interface{}
	Create(id string, object interface{}) error
	Update(id string, object interface{}, changes map[string]interface{}) error
	Delete(id string, tableName string) error
}
