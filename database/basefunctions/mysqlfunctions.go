package basefunctions

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"strings"
	"websays/database/baseconnections"

	"websays/database/basetypes"
)

// MySqlFunctions is a concrete implementation of the BaseFucntionsInterface for MySQL database.
type MySqlFunctions struct {
}

// GetFunctions returns the MySqlFunctions instance as a BaseFucntionsInterface.
func (u *MySqlFunctions) GetFunctions() BaseFucntionsInterface {
	return u
}

// EnsureIndex ensures an index for the specified database and collection in MySQL.
// It takes the database name, collection name, and a sample data interface for table schema.
// This function creates a table with the schema based on the data interface.
func (u *MySqlFunctions) EnsureIndex(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) error {
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	query := `CREATE TABLE IF NOT EXISTS ` + string(collectionName) + ` (`
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()

	if dataType.Kind() != reflect.Struct {
		return errors.New("Required a struct for data")
	}

	columns := ""

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		tags := strings.Split(field.Tag.Get("db"), ",")

		if columns != "" {
			columns += ","
		}

		columns += strings.Join(tags, " ")
	}

	query += columns + ");"
	_, err := conn.Exec(query)
	return err
}

// GetNextID returns the next available ID for MySQL storage.
func (u *MySqlFunctions) GetNextID() int {
	return 0
}

// Add inserts data into the MySQL database.
// It takes the database name, collection name, and data interface to be inserted.
// This function dynamically generates an SQL INSERT statement based on the data interface and inserts the data.
func (u *MySqlFunctions) Add(dbName basetypes.DBName, collectionName basetypes.CollectionName, data interface{}) (int, error) {
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	query := "INSERT INTO " + string(collectionName)

	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()

	if dataType.Kind() != reflect.Struct {
		return 0, errors.New("Required a struct for data")
	}

	var columns []string
	var placeholders []string
	values := make([]interface{}, 0)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		tag := strings.Split(field.Tag.Get("db"), ",")[0]

		if tag == "" {
			continue
		}

		value := dataValue.Field(i).Interface()
		values = append(values, value)

		columns = append(columns, tag)
		placeholders = append(placeholders, "?")
	}

	query += "(" + strings.Join(columns, ", ") + ")"
	query += " VALUES(" + strings.Join(placeholders, ", ") + ")"

	res, err := conn.Exec(query, values...)
	lastId, _ := res.LastInsertId()
	return int(lastId), err
}

// FindOne retrieves data from the MySQL database based on a condition.
// It takes the database name, collection name, and a condition map to filter data.
// This function dynamically generates an SQL SELECT statement based on the condition and retrieves one record.
func (u *MySqlFunctions) FindOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, cond interface{}) (interface{}, error) {

	condition := cond.(map[string]interface{})
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	query := "SELECT * FROM " + string(collectionName)

	whereClause := ""
	values := make([]interface{}, 0)

	for key, val := range condition {
		if whereClause != "" {
			whereClause += " AND "
		} else {
			whereClause += " WHERE "
		}
		whereClause += key + "= ? "
		values = append(values, val)
	}

	query += whereClause
	log.Println(query, values)
	rows, err := conn.Query(query, values...)

	return rows, err
}

// UpdateOne updates data in the MySQL database based on a query condition.
// It takes the database name, collection name, a query map for filtering, data to update, and an upsert flag.
// This function dynamically generates an SQL UPDATE statement based on the query condition and updates one record.
func (u *MySqlFunctions) UpdateOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, query interface{}, data interface{}, upsert bool) error {
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	dbQuery := "UPDATE " + string(collectionName) + " SET "

	values := make([]interface{}, 0)
	dataMap := data.(map[string]interface{})
	setClause := ""

	for key, val := range dataMap {
		if setClause != "" {
			setClause += "," + key + "= ? "
		} else {
			setClause += key + "= ?"
		}
		values = append(values, val)
	}

	condition := query.(map[string]interface{})
	whereClause := ""
	for key, val := range condition {
		if whereClause != "" {
			whereClause += " AND "
		} else {
			whereClause += " WHERE "
		}
		whereClause += key + "= ? "
		values = append(values, val)
	}

	dbQuery += setClause + whereClause + " LIMIT 1"
	_, err := conn.Exec(dbQuery, values...)
	return err
}

// DeleteOne deletes data from the MySQL database based on a query condition.
// It takes the database name, collection name, and a query map for filtering data to delete.
// This function dynamically generates an SQL DELETE statement based on the query condition and deletes one record.
func (u *MySqlFunctions) DeleteOne(dbName basetypes.DBName, collectionName basetypes.CollectionName, cond interface{}) error {
	conn := baseconnections.GetInstance().GetConnection(basetypes.MYSQL).GetDB(basetypes.MYSQL).(*sql.DB)
	condition := cond.(map[string]interface{})
	query := "DELETE FROM " + string(collectionName)
	whereClause := ""
	values := make([]interface{}, 0)

	for key, val := range condition {
		if whereClause != "" {
			whereClause += " AND "
		} else {
			whereClause += " WHERE "
		}
		whereClause += key + "= ? "
		values = append(values, val)
	}

	query += whereClause + " LIMIT 1"

	_, err := conn.Query(query, values...)

	return err
}
