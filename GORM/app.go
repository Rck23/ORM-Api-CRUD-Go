package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/Rck23/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(localhost:3306)/GoPerson?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a BD.")
	}

	// AutoMigrate crea automáticamente las tablas en la base de datos basadas en las estructuras definidas
	db.AutoMigrate(&models.Person{})

	router := gin.Default()

	// Obtener todos los registros
	router.GET("/GetAllPersons", func(c *gin.Context) {
		var Persons []models.Person
		db.Find(&Persons)
		c.JSON(http.StatusOK, Persons)
	})

	// Obtener un registro por ID
	router.GET("/GetPersonById/:id", func(c *gin.Context) {
		var Person models.Person
		db.First(&Person, c.Param("id"))
		c.JSON(http.StatusOK, Person)
	})

	// Crear un nuevo registro
	router.POST("/CreatePerson", func(c *gin.Context) {
		var Person models.Person
		c.BindJSON(&Person)
		db.Create(&Person)
		c.JSON(http.StatusOK, Person)
	})

	// Actualizar un registro existente
	router.PUT("/UpdatePerson/:id", func(c *gin.Context) {
		var Person models.Person
		db.First(&Person, c.Param("id"))
		c.BindJSON(&Person)
		db.Save(&Person)
		c.JSON(http.StatusOK, Person)
	})

	// Eliminar un registro
	router.DELETE("/DeletePerson/:id", func(c *gin.Context) {
		var Person models.Person
		db.Delete(&Person, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "Registro eliminado"})
	})

	// Ejecutar el servidor
	router.Run(":8080")

}
