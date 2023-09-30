package baseconnections

import (
	"database/sql"
	"websays/config"
	"websays/database/basetypes"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlConnection represents a MySQL database connection.
type MysqlConnection struct {
	dbName string
	db     *sql.DB
}

// CreateConnection creates a MySQL database connection using the configuration values from the singleton config instance.
// It returns the created connection and any error encountered during connection setup.
func (u *MysqlConnection) CreateConnection() (ConnectionInterface, error) {
	dsn := config.GetInstance().Database.Username + ":" + config.GetInstance().Database.Password + "@tcp(" + config.GetInstance().Database.Host + ":" + config.GetInstance().Database.Port + ")/" + config.GetInstance().Database.DBName
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	u.db = db
	return u, nil
}

// GetDB returns the MySQL database instance associated with this connection.
func (u *MysqlConnection) GetDB(dbType basetypes.DbType) interface{} {
	return u.db
}
