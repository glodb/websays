package baseconnections

import "websays/database/basetypes"

//Keeping it open for multiple db or own db connections in microservices
type MemoryConnection struct {
	folderName string
}

func (u *MemoryConnection) CreateConnection() (ConntectionInterface, error) {
	return nil, nil
}

func (u *MemoryConnection) GetDB(dbType basetypes.DbType) interface{} {
	return nil
}
