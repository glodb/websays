package baseconnections

import (
	"database/sql"
	"websays/config"
	"websays/database/basetypes"

	_ "github.com/go-sql-driver/mysql"
)

//Keeping it open for multiple db or own db connections in microservices
type MysqlConnection struct {
	dbName string
	db     *sql.DB
}

func (u *MysqlConnection) CreateConnection() (ConntectionInterface, error) {
	dsn := config.GetInstance().Database.Username + ":" + config.GetInstance().Database.Password + "@tcp(" + config.GetInstance().Database.Host + ":" + config.GetInstance().Database.Port + ")/" + config.GetInstance().Database.DBName
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	u.db = db
	return u, nil
}

func (u *MysqlConnection) GetDB(dbType basetypes.DbType) interface{} {
	return u.db
}
