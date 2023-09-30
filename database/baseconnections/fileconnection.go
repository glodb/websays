package baseconnections

import "websays/database/basetypes"

//File connection does not require an active connection like dbs
type FileConnection struct {
	folderName string
}

func (u *FileConnection) CreateConnection() (ConnectionInterface, error) {
	return nil, nil
}

func (u *FileConnection) GetDB(dbType basetypes.DbType) interface{} {
	return nil
}
