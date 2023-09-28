package baseinterfaces

import (
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basevalidators"
)

type Controller interface {
	basefunctions.BaseFucntionsInterface
	BaseControllerFactory
	basevalidators.ValidatorInterface
	SetBaseFunctions(basefunctions.BaseFucntionsInterface)
	GetCollectionName() basetypes.CollectionName
	DoIndexing() error
	RegisterApis()
	GetDBName() basetypes.DBName
}
