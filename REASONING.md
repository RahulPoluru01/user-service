### Architecture
<img width="1297" height="606" alt="image" src="https://github.com/user-attachments/assets/d17fc857-b112-4efa-b81a-2fe45c1a9bbe" />

- Postman acts as the client that sends HTTP requests to my GoFiber backend.
- The GoFiber server receives each request and routes it to the appropriate handler.
- From there, the request flows through internal layers where each layer has a single responsibility:
- parsing and validating input, applying business logic, interacting with the database, and finally constructing the response.
- The backend talks to PostgreSQL to store and fetch data, and returns the result back to the client as JSON.

---
### End-to-End Flow: GET /users/{id}

When the server starts, the `main` function initializes the logger, registers middleware, establishes the PostgreSQL database connection, creates the SQLC queries instance, and wires all layers together using dependency injection. This sets up the full request handling pipeline before accepting traffic.

1. A client sends an HTTP GET request to `/users/{id}`.

2. The request first passes through Fiber middleware.  
   The request ID middleware attaches a UUID to the request, and the logging middleware measures and logs the request lifecycle and duration.

3. The router matches the path and forwards the request to the `GetUserByID` handler.

4. The handler extracts the `id` path parameter and performs basic validation.  
   Since this is a GET request, there is no request body to parse.

5. The handler calls the service layer with the extracted user ID.

6. The service layer invokes the repository to fetch user data.  
   The service does not know how data is stored; it only depends on the repository interface.

7. The repository implementation (`PostgresUserRepository`) calls the SQLC-generated `GetUserByID` method, which executes the SQL query against PostgreSQL and maps the result to a Go struct.

8. The fetched user data is returned back to the service layer.  
   The service calculates the age dynamically from the date of birth using the current date.

9. The service constructs a response model including the calculated age and returns it to the handler.

10. The handler sends the response as JSON back to the client with an HTTP 200 status.

This flow ensures that:
- HTTP concerns stay in the handler layer
- Business logic stays in the service layer
- Database access stays in the repository and SQLC layers
- Each layer has a single responsibility

---
### Why I built it this way

I implemented the project using a layered architecture to clearly separate responsibilities across the system.

Each layer has a focused role:
- Handlers deal with HTTP concerns like parsing requests and returning responses
- Services contain business logic such as calculating age
- Repositories handle data persistence and abstract database access

This structure helps:
- Keep the code organized and easy to navigate
- Ensure each layer has a single responsibility
- Make debugging easier by isolating issues to a specific layer
- Maintain a clear and predictable request flow from API to database

By following this pattern, the code becomes easier to understand, extend, and maintain as the project grows.

---
### discussing SQLC 
SQLC generates type-safe Go code directly from SQL queries and the database schema.  
This allows me to write raw SQL while still getting compile-time validation.

- SQL errors are caught at build time instead of runtime
- Queries are explicit and easy to reason about
- No hidden or auto-generated queries as in ORMs
- Strong typing between SQL and Go models

This approach gives better control over performance and avoids surprises in database behavior.

---
### Learnings and Reflections

#### How interfaces helped reduce tight coupling

Using a repository interface allowed the service layer to depend only on behavior rather than a concrete database implementation.  
The service does not need to know how data is fetched or stored, it only calls the methods defined by the interface.

This design:
- Decouples business logic from persistence
- Improves flexibility to switch databases or implementations
- Makes the code easier to reason about and test

It helped me understand how abstraction can keep layers independent while still maintaining control over behavior.


#### How I adapted to Go for this task

This was my first backend project in Go. I focused on learning only the concepts required to build this service effectively.

What helped me adapt:
- Go’s simplicity and minimal syntax made it beginner friendly
- Structs and explicit methods felt intuitive coming from Java classes
- I could relate layered architecture and dependency injection concepts from my Java backend experience

Although I did not aim to master the entire language in a short time, I learned the necessary concepts step by step while building the project.  
This approach helped me stay productive and gain confidence in working with Go in a real backend scenario.


