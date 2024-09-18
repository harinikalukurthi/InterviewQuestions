package main

import (
    "fmt"
    "encapsulation/project"
)

func main() {
    // Create a new instance of Person
    p := project.Person{Age: 30}

    // Use the setter method to set the name
    p.SetName("John")

    // Use the getter method to get the name
    fmt.Println("Name:", p.GetName()) // Name: John

    // Access the public field directly
    fmt.Println("Age:", p.Age) // Age: 30

    // Direct access to the private field 'name' is not possible
    //fmt.Println(p.name) // This would cause a compile-time error (uncommenting this line will result in an error)
}
