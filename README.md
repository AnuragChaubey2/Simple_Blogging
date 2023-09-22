# Simple_Blogging

```Simple Blogging API```

The Simple Blogging API is designed to provide a reliable and efficient way to manage blog posts. It incorporates best practices and follows proper separation of concerns, making it a robust solution for building blog-based applications.

`Key Features:`

`Proper Separation of Concerns:`
The project follows a clean architecture with distinct layers: Store, Service, and HTTP.
Each layer is responsible for specific tasks, ensuring modularity and maintainability.

`Dependency Injection:`
Dependency injection is used to facilitate communication between layers.
This design pattern promotes code reusability and testability.

`Error Handling:`
Robust error handling is implemented throughout the application.
Different HTTP status codes are used to indicate the result of each operation.

`HTTP Handlers:`
HTTP handlers are implemented using the Gorilla Mux router.
They handle various use cases such as creating, reading, updating, and deleting posts.

`CRUD Operations:`
The API supports all CRUD (Create, Read, Update, Delete) operations for managing blog posts.
Proper validation and error handling are in place for each operation.

`Custom Response Messages:`
Custom success and error messages are included in JSON responses to provide clear feedback to clients.
Success messages confirm the successful execution of operations.

```Usage```

`Endpoints:`

/posts (GET): Retrieve a list of all blog posts.
/posts/{id} (GET): Retrieve a specific blog post by ID.
/posts (POST): Create a new blog post.
/posts/{id} (PUT): Update an existing blog post by ID.
/posts/{id} (DELETE): Delete a blog post by ID.