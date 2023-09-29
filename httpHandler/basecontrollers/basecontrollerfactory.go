package basecontrollers

import (
	"sync"
	"websays/app/controllers"
	"websays/app/validators"
	"websays/config"
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basecontrollers/baseinterfaces"
)

var instance *controllersObject
var once sync.Once

// Controllers struct
type controllersObject struct {
	controllers map[string]baseinterfaces.Controller
}

// Singleton. Returns a single object of Factory
func GetInstance() *controllersObject {
	// var instance
	once.Do(func() {
		instance = &controllersObject{}
		instance.controllers = make(map[string]baseinterfaces.Controller)
	})
	return instance
}

// createController is a factory to return the appropriate controller
func (c *controllersObject) GetController(controllerType string) (baseinterfaces.Controller, error) {
	if _, ok := c.controllers[controllerType]; ok {
		return c.controllers[controllerType], nil
	} else {
		c.registerControllers(controllerType, false)
		return c.controllers[controllerType], nil
	}
}

/**
*To all developers are future me,
*Although this is lazy flyweight factory it doesn't work as lazy factory for web server
*It will register all the controllers defined in the config for web but it will still be flyweight
*Don't call the RegisterControllers if its not web
 */
func (c *controllersObject) RegisterControllers() {
	localControllers := config.GetInstance().Controllers
	for i := range localControllers {
		c.registerControllers(localControllers[i], true)
	}
}

func (c *controllersObject) registerControllers(key string, registerApis bool) {
	var funcs *basefunctions.BaseFucntionsInterface
	switch key {
	case Article:
		c.controllers[key] = &controllers.Article{BaseControllerFactory: c, ValidatorInterface: &validators.ArticleValidator{}}
		funcs, _ = basefunctions.GetInstance().GetFunctions(basetypes.MEMORY, c.controllers[key].GetDBName())
	case Category:
		c.controllers[key] = &controllers.Category{BaseControllerFactory: c, ValidatorInterface: &validators.CategoryValidator{}}
		funcs, _ = basefunctions.GetInstance().GetFunctions(basetypes.FILE, c.controllers[key].GetDBName())
	case Product:
		c.controllers[key] = &controllers.Product{BaseControllerFactory: c, ValidatorInterface: &validators.ProductValidator{}}
		funcs, _ = basefunctions.GetInstance().GetFunctions(basetypes.MYSQL, c.controllers[key].GetDBName())

	}
	c.controllers[key].SetBaseFunctions(*funcs)
	c.controllers[key].DoIndexing()
	if registerApis {
		c.controllers[key].RegisterApis()
	}
}
