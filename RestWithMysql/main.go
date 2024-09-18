package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// Employee represents the employee model
type Employee struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    City string `json:"city"`
}

var db *gorm.DB
var err error

func main() {
    // Database connection
    dsn := "root:new_password@tcp(127.0.0.1:3306)/employee_db"
    db, err = gorm.Open(mysql.Open(dsn))
    if err != nil {
        panic("failed to connect database")
    }


    // Initialize Gin router
    r := gin.Default()

    // Routes
    r.GET("/employees", getEmployees)
    r.GET("/employees/:id", getEmployeeByID)
    r.POST("/employees", createEmployee)
    r.PUT("/employees/:id", updateEmployee)
    r.DELETE("/employees/:id", deleteEmployee)

    // Run the server
    r.Run(":8080")
}

// Handlers
func getEmployees(c *gin.Context) {
    var employees []Employee
    db.Find(&employees)
    c.JSON(http.StatusOK, employees)
}

func getEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    var employee Employee
    if result := db.First(&employee, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }
    c.JSON(http.StatusOK, employee)
}

func createEmployee(c *gin.Context) {
    var employee Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&employee)
    c.JSON(http.StatusOK, employee)
}

func updateEmployee(c *gin.Context) {
    id := c.Param("id")
    var employee Employee
    if result := db.First(&employee, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }

    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&employee)
    c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c *gin.Context) {
    id := c.Param("id")
    if result := db.Delete(&Employee{}, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
