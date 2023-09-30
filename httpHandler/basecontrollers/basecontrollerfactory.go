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

// controllersObject is a singleton factory responsible for creating and managing controller instances.
type controllersObject struct {
	controllers map[string]baseinterfaces.Controller
}

// GetInstance returns a single instance of the controllersObject.
// This function ensures that only one instance of the controllersObject is created and shared across the application.
func GetInstance() *controllersObject {
	// var instance
	once.Do(func() {
		instance = &controllersObject{}
		instance.controllers = make(map[string]baseinterfaces.Controller)
	})
	return instance
}

// GetController retrieves or creates a controller instance based on the provided controllerType.
// If the controller with the specified type exists, it is retrieved from the map. Otherwise, it is created and registered.
// It returns the Controller interface and an error if the controller cannot be found or instantiated.
func (c *controllersObject) GetController(controllerType string) (baseinterfaces.Controller, error) {
	if _, ok := c.controllers[controllerType]; ok {
		return c.controllers[controllerType], nil
	} else {
		c.registerControllers(controllerType, false)
		return c.controllers[controllerType], nil
	}
}

/**
 * Although this is a lazy flyweight factory, it doesn't work as a lazy factory for web servers.
 * It will register all the controllers defined in the config for web, but it will still be flyweight.
 * Don't call the RegisterControllers method if it's not intended for web use.
 */
func (c *controllersObject) RegisterControllers() {
	localControllers := config.GetInstance().Controllers
	for i := range localControllers {
		c.registerControllers(localControllers[i], true)
	}
}

// registerControllers creates and registers a specific controller based on the provided key.
// It sets the controller's base functions, performs indexing, and registers APIs if needed.
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
