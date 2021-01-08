package main

import "github.com/kataras/iris"

func main() {
	// Run the function to create the new Iris App
	app := newApp()

	// Start the web server on port 8080
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func newApp() *iris.Application {
	// Initialize a new Iris App
	app := iris.New()

	// Register the request handler for the endpoint "/"
	app.Get("/", func(ctx iris.Context) {
		// Return something by adding it to the context
		ctx.Text("Hello World")
	})

	app.Get("/ola", func(ctx iris.Context) {
		// Return something by adding it to the context
		ctx.Text("rota ola")
	})

	return app
}
