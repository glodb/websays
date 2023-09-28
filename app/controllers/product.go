package controllers

import (
	"websays/config"
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basecontrollers/baseinterfaces"
	"websays/httpHandler/basevalidators"
)

type Product struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
}

func (pro *Product) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

func (pro *Product) GetCollectionName() basetypes.CollectionName {
	return "products"
}

func (pro *Product) DoIndexing() error {
	return nil
}

func (pro *Product) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	pro.BaseFucntionsInterface = inter
}

func (pro *Product) RegisterApis() {
}
