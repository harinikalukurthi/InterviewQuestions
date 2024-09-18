A web framework is a collection of software components and libraries that provide the necessary tools for developing web applications.

libraries:- The standard library is a collection of Go packages that provide commonplace functionality, such as string manipulation, I/O, math, and an HTTP client/server.

router := gin.Default():- create a new instance of the Gin engine with some default middleware already attached to it. 

// Create Gin router
    router := gin.Default()

// Register Routes
    router.GET("/", homePage)

// Start the server
    router.Run()

Gin uses the : symbol followed by the name of the parameter you want to capture in URLs.

example:-router.POST("/recipes", recipesHandler.CreateRecipe)