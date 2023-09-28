package controllers

import (
	"websays/config"
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basecontrollers/baseinterfaces"
	"websays/httpHandler/basevalidators"
)

type Category struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
}

func (cat *Category) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

func (cat *Category) GetCollectionName() basetypes.CollectionName {
	return "categories"
}

func (cat *Category) DoIndexing() error {
	return nil
}

func (cat *Category) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	cat.BaseFucntionsInterface = inter
}

func (cat *Category) RegisterApis() {
}
