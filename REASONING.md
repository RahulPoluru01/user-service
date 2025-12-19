# Architecture
<img width="1297" height="606" alt="image" src="https://github.com/user-attachments/assets/d17fc857-b112-4efa-b81a-2fe45c1a9bbe" />
- Postman acts as the client that sends HTTP requests to my GoFiber backend.
- The GoFiber server receives each request and routes it to the appropriate handler.
- From there, the request flows through internal layers where each layer has a single responsibility:
- parsing and validating input, applying business logic, interacting with the database, and finally constructing the response.
- The backend talks to PostgreSQL to store and fetch data, and returns the result back to the client as JSON.
