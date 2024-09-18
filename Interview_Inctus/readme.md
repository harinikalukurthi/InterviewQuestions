what are design patterns? types
Singleton design pattern. how to create a singleton struct in golang? benefits. 
package main

import (
    "fmt"
    "sync"
)

type Singleton struct {
    // Add fields here
}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}

func main() {
    s1 := GetInstance()
    s2 := GetInstance()

    if s1 == s2 {
        fmt.Println("Both instances are the same")
    }
}

In this example, the Singleton pattern ensures that only one instance of the Singleton struct is created and used throughout the application. The sync.Once mechanism guarantees that the initialization code runs only once, even if accessed from multiple goroutines.


need to write test cases for a function.
what is integration testing.
what and all will happen before the code changes reaches prod.
what ci/cd?
how do you scalable your microservices.
what is graphql?
what is 
difference between http and https.
what hotfix,feature branch. how do you manage branches in git.
difference of param and query. both of them we use for GET only right, why do we have these two.
how do you stop the goroutine on demand. sol: using context or channels.
how to check the type of the variable in runtime in go? lets say i dont want to print the type, but i want to check do something if that is equal to int and do something else if it is string.

what is schema ? 