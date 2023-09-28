package baseconnections

import "websays/database/basetypes"

//Keeping it open for multiple db or own db connections in microservices
type TextConnection struct {
	folderName string
}

func (u *TextConnection) CreateConnection() (ConntectionInterface, error) {
	return nil, nil
}

func (u *TextConnection) GetDB(dbType basetypes.DbType) interface{} {
	return nil
}
