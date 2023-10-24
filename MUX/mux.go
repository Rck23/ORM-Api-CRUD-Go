package main

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Name  string `validate:"required,min=3,max=16"`
	Email string `validate:"required,email"`
}

func main() {
	dsn := "root:admin@tcp(localhost:3306)/GoPerson?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a BD.")
	}

	// AutoMigrate crea automáticamente las tablas en la base de datos basadas en las estructuras definidas
	db.AutoMigrate(&Person{})

	router := mux.NewRouter()

	// Obtener todos los registros
	router.HandleFunc("/GetAllPersons", func(w http.ResponseWriter, r *http.Request) {
		var Persons []Person
		db.Find(&Persons)
		json.NewEncoder(w).Encode(Persons)
	}).Methods("GET")

	// Obtener un registro por ID
	router.HandleFunc("/GetPersonById/{id}", func(w http.ResponseWriter, r *http.Request) {
		var Person Person
		db.First(&Person, mux.Vars(r)["id"])
		json.NewEncoder(w).Encode(Person)
	}).Methods("GET")

	// Crear un nuevo registro
	router.HandleFunc("/CreatePerson", func(w http.ResponseWriter, r *http.Request) {
		var Person Person
		json.NewDecoder(r.Body).Decode(&Person)
		db.Create(&Person)
		json.NewEncoder(w).Encode(Person)
	}).Methods("POST")

	// Actualizar un registro existente
	router.HandleFunc("/UpdatePerson/{id}", func(w http.ResponseWriter, r *http.Request) {
		var Person Person
		db.First(&Person, mux.Vars(r)["id"])
		json.NewDecoder(r.Body).Decode(&Person)
		db.Save(&Person)
		json.NewEncoder(w).Encode(Person)
	}).Methods("PUT")

	// Eliminar un registro
	router.HandleFunc("/DeletePerson/{id}", func(w http.ResponseWriter, r *http.Request) {
		var Person Person
		db.Delete(&Person, mux.Vars(r)["id"])
		json.NewEncoder(w).Encode(map[string]string{"message": "Registro eliminado"})
	}).Methods("DELETE")

	// Ejecutar el servidor
	log.Fatal(http.ListenAndServe(":8080", router))

}
