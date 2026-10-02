package api

import (
	"github.com/pocketbase/pocketbase/core"
)

// RegisterRoutes sets up custom HTTP endpoints on the PocketBase Chi router.
func RegisterRoutes(se *core.ServeEvent) {
	registerExampleRoutes(se)
}
