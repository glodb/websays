package basetypes

type CollectionName string
type DBName string
type DbType int

const (
	MYSQL  DbType = 1
	FILE   DbType = 2
	MEMORY DbType = 3
)
