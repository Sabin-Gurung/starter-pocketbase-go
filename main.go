package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/untitledev-np/go-pocketbase-starter/internal/api"
	"github.com/untitledev-np/go-pocketbase-starter/internal/bootstrap"
	"github.com/untitledev-np/go-pocketbase-starter/internal/config"
)

func main() {
	cfg := config.Load()
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Bootstrap initial admin account in dev mode
		if cfg.DevMode && cfg.AdminEmail != "" && cfg.AdminPassword != "" {
			if err := bootstrap.EnsureAdmin(app, cfg.AdminEmail, cfg.AdminPassword); err != nil {
				log.Printf("Failed to bootstrap admin account: %v", err)
			}
		}

		// Register custom API handlers
		api.RegisterRoutes(se)

		// Serve static frontend assets from ./pb_public fallback
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	log.Printf("Starting PocketBase backend server on %s", cfg.Port)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
