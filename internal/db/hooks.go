package db

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
)

// RegisterHooks binds global event hooks to the PocketBase instance.
// Call this function from main.go or OnServe if you need custom record event hooks.
func RegisterHooks(app core.App) {
	// Example: Before creating a record in any collection
	app.OnRecordCreateExecute().BindFunc(func(e *core.RecordEvent) error {
		log.Printf("[Hook] Record creating in collection: %s", e.Record.Collection().Name)
		return e.Next()
	})
}
