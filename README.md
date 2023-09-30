# Websays Project

Websays is a project created for the Websays company's test. This project serves as an API server using the Gorilla/Mux library and utilizes MySQL as its database. Below, you will find information on how to set up and run the project, as well as details about its folder structure.

## Getting Started

To run the project, follow these steps:

1. Build and start the Docker containers:

```bash
sudo docker-compose up --build
```

2. Open another command prompt and access the MySQL container:

```bash
sudo docker exec -it mysql_db mysql -p
```

3. When prompted for a password, enter `gotest`.

4. In the MySQL prompt, run the following command to create the `websays` database:

```bash
CREATE DATABASE IF NOT EXISTS websays;
```


5. Go back to the original command prompt where you started the Docker containers and restart the containers. The project should now be up and running.

## Folder Structure

The project has a well-organized folder structure that separates different components. Here's an overview of the main folders:

### app
- **controllers**: Contains the API controllers that define how API endpoints handle requests and responses.
- **middlewares**: Houses middleware components used to perform tasks on all API requests, such as CORS handling.
- **models**: Contains data models used by the application, representing entities like articles, categories, and products.
- **validators**: Holds validator components responsible for validating data for various API endpoints.

### config
- **configModels**: Defines configuration models used throughout the project.

### database
- **baseconnections**: Contains database connection interfaces.
- **basefunctions**: Provides interfaces for basic database operations like CRUD.
- **basemodels**: Defines interfaces for database models.
- **basetypes**: Contains basic types used in the project's database operations.

### httpHandler
- **basecontrollers**: Defines base controller interfaces used by the application controllers.
- **basemodels**: Contains base models used by the application's models.
- **baserouter**: Provides routing functionality for handling HTTP requests.
- **basevalidators**: Defines base validator interfaces.
- **responses**: Handles HTTP responses and defines response codes.

### tests
- This folder is used for storing test cases and test-related files, ensuring the reliability of the application.

## Project Overview

The project's main functionality is exposed through API endpoints defined in the `app/controllers` directory. These controllers handle incoming requests, interact with the database, and return responses. The use of Gorilla/Mux for routing simplifies this process.

For the database, the project uses MySQL with the `sql` package for database operations. Instead of employing an ORM, it directly uses SQL queries for CRUD operations, keeping the application lightweight.

Feel free to explore each folder and the corresponding components to understand the project's structure and functionality better.

For any questions or issues, please don't hesitate to reach out.

