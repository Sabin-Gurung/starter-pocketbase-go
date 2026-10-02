package api

import (
	"github.com/pocketbase/pocketbase/core"
)

func registerExampleRoutes(se *core.ServeEvent) {
	se.Router.GET("/api/hello", func(e *core.RequestEvent) error {
		return e.JSON(200, map[string]string{
			"message": "Hello from PocketBase starter backend!",
		})
	})
}
