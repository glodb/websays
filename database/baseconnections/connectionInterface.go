package baseconnections

import "websays/database/basetypes"

type ConntectionInterface interface {
	CreateConnection() (ConntectionInterface, error)
	GetDB(basetypes.DbType) interface{}
}
