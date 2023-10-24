# Go API ORM

## Introducción

Esta API está desarrollada en Go 1.21.1 y utiliza GORM y Gorilla Mux para acceder a la base de datos MySql y definir las rutas HTTP, respectivamente.

## Requisitos

Para ejecutar esta API, necesitas tener instalado lo siguiente:

* Go
* MySQL
* GORM
* Gorilla Mux

## Ejecutar la API

Para ejecutar la API, clona el repositorio y ejecuta el siguiente comando:

```
go run app.go
```
O 

```
go run mux.go
```

La API se ejecutará en el puerto 8080. Puedes acceder a ella en tu navegador web en `http://localhost:8080`.

## Rutas HTTP

La API proporciona las siguientes rutas HTTP:

| Ruta | Método | Descripción |
|---|---|---|
| `/GetAllPersons` | GET | Obtener todos los registros de personas. |
| `/GetPersonById/:id` | GET | Obtener un registro de persona por su ID. |
| `/CreatePerson` | POST | Crear un nuevo registro de persona. |
| `/UpdatePerson/:id` | PUT | Actualizar un registro de persona existente. |
| `/DeletePerson/:id` | DELETE | Eliminar un registro de persona. |

## Ejemplos de uso

### Obtener todos los registros de personas

Para obtener todos los registros de personas, envía una solicitud GET a la siguiente ruta:

```
http://localhost:8080/GetAllPersons
```

Respuesta:

```json
[{
  "id": 1,
  "name": "Juan Pérez",
  "email": "juan.perez@example.com"
}, {
  "id": 2,
  "name": "María González",
  "email": "maria.gonzalez@example.com"
}]
```

### Obtener un registro de persona por su ID

Para obtener un registro de persona por su ID, envía una solicitud GET a la siguiente ruta:

```
http://localhost:8080/GetPersonById/1
```

Respuesta:

```json
{
  "id": 1,
  "name": "Juan Pérez",
  "email": "juan.perez@example.com"
}
```

### Crear un nuevo registro de persona

Para crear un nuevo registro de persona, envía una solicitud POST a la siguiente ruta:

```
http://localhost:8080/CreatePerson
```

En el cuerpo de la solicitud, envía un objeto JSON con los datos de la nueva persona:

```json
{
  "name": "Pedro López",
  "email": "pedro.lopez@example.com"
}
```

Respuesta:

```json
{
  "id": 3,
  "name": "Pedro López",
  "email": "pedro.lopez@example.com"
}
```

### Actualizar un registro de persona existente

Para actualizar un registro de persona existente, envía una solicitud PUT a la siguiente ruta:

```
http://localhost:8080/UpdatePerson/2
```

En el cuerpo de la solicitud, envía un objeto JSON con los datos actualizados de la persona:

```json
{
  "name": "María González (Actualizada)",
  "email": "maria.gonzalez@example.com"
}
```

Respuesta:

```json
{
  "id": 2,
  "name": "María González (Actualizada)",
  "email": "maria.gonzalez@example.com"
}
```

### Eliminar un registro de persona

Para eliminar un registro de persona, envía una solicitud DELETE a la siguiente ruta:

```
http://localhost:8080/DeletePerson/3
```

Respuesta:

```json
{
  "message": "Registro eliminado"
}
```
