package main

import "http_server/server"

func main() {
	router := server.CreateRouter()
	router.Use(server.CORSMiddleware())
	router.Run("localhost:9999")
}

// curl localhost:9999/v1/foo
// curl localhost:9999/v1/foo/0
