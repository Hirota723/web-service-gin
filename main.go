package main

import "example/web-service-gin/router"

func main() {
	r := router.SetupRouter()
	r.Run("localhost:8080")
}
