package baseinterfaces

// BaseControllerFactory is an interface that defines the contract for a factory responsible for creating and providing instances of controllers within the application.
// Controllers encapsulate specific business logic, data manipulation, and API endpoints, and they serve as intermediaries between the HTTP layer and the underlying database.
// Implementing this interface enables the dynamic creation and retrieval of controllers based on their names or identifiers.
type BaseControllerFactory interface {
	// GetController retrieves a controller instance based on the provided controller name or identifier.
	// This factory method allows for the creation and retrieval of different controllers by name, facilitating dynamic controller management.
	// It returns a Controller interface and an error if the requested controller cannot be found or instantiated.
	// Example Usage:
	//   controller, err := factory.GetController("UserController")
	//   if err != nil {
	//       // Handle the error, possibly indicating that the requested controller does not exist.
	//   } else {
	//       // Use the retrieved controller to perform operations and handle API requests.
	//       controller.RegisterApis()
	//   }
	GetController(controllerName string) (Controller, error)
}
