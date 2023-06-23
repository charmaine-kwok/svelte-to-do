# Svelte To-Do App Backend

This is the backend server for the Svelte To-Do App. The backend server is built using [Go](https://go.dev) and the [Gin Web Framework](https://gin-gonic.com). It is responsible for handling API requests and interacting with the database.

## Endpoints

### Todos

```sh
# Create a new to-do item
http POST http://localhost:8080/todo subject=$subject
```

```sh
# Get all to-do items
http GET http://localhost:8080/todo/todos
```

```sh
# Delete a to-do item
http DELETE http://localhost:8080/todo/:id
```

```sh
# Update an existing to-do item
http PATCH http://localhost:8080/todo/:id
```

### Dones

```sh
# Create a new done item
http POST http://localhost:8080/done subject=$subject
```

```sh
# Get all done items
http GET http://localhost:8080/done/dones
```

```sh
# Delete a done item
http DELETE http://localhost:8080/done/:id
```
