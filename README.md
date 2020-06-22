# Xmen-meli
Aplicación que detecta a partir de un ADN si un humano es mutante.

## Requerimientos

Para el uso de la aplicación es necesario tener instalado [GoLang](https://golang.org/dl/) y [MongoDB](https://docs.mongodb.com/manual/installation/).

## Puesta en marcha

Dentro del proyecto ejecutar los siguientes comandos:
- Nos situamos en el path donde contiene nuestro archivo main para ejecutar
```bash
cd cmd/app/
```
- De manera opcional se pueden especificar variables de entorno para configurar el MongoDB
```bash
MONGO_HOST=localhost

MONGO_PORT=27017

MONGO_DB=xmendb
```
- Corremos nuestra aplicación
```bash
go run main.go
```
-------------------------
## Uso

Con la aplicación corriendo se habilitan los siguientes contextos, que pueden ser consultadas a través de [postman](https://www.postman.com/downloads/).

####Chequeo de ADN 

En el contexto `/mutant/` recibe una llamada `POST` donde en el body se envía un `JSON` con la información de ADN para que este analice si es ADN mutante o humano.

- Por ejemplo hacemos una petición `POST → /mutant/` con el body:

```json
{
    "dna": [
      "ATGCGA",
      "CAGTGC",
      "TTATGT",
      "AGAAGG",
      "CCCCTA",
      "TCACTG"
    ]
}
```

En caso de ser un mutante, debería devolver un `HTTP 200-OK`, de lo contrario un
`403-Forbidden`. Para este ejemplo nos va a devolver un `200-OK`.

####Consulta de estadísticas 

En el contexto `/stats` recibe una llamada `GET` y devuelve un `JSON` con las estadísticas de las verificaciones de ADN realizadas

- Respuesta de ejemplo:
 
```json
{
    "count_mutant_dna": 40,
    "count_human_dna": 100,
    "ratio": 0.4
}
```