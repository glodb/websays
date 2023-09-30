package baseconnections

import "websays/database/basetypes"

//File connection does not require an active connection like dbs
type MemoryConnection struct {
	folderName string
}

func (u *MemoryConnection) CreateConnection() (ConnectionInterface, error) {
	return nil, nil
}

func (u *MemoryConnection) GetDB(dbType basetypes.DbType) interface{} {
	return nil
}
