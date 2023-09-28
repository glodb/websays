package baseconnections

import "websays/database/basetypes"

//Keeping it open for multiple db or own db connections in microservices
type FileConnection struct {
	folderName string
}

func (u *FileConnection) CreateConnection() (ConntectionInterface, error) {
	return nil, nil
}

func (u *FileConnection) GetDB(dbType basetypes.DbType) interface{} {
	return nil
}
