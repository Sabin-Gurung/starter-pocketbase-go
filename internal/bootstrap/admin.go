package bootstrap

import (
	"database/sql"
	"errors"
	"log"

	"github.com/pocketbase/pocketbase/core"
)

// EnsureAdmin checks if a superuser with the given email exists, and creates one if missing.
func EnsureAdmin(app core.App, email, password string) error {
	_, err := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, email)
	if err == nil {
		return nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	superusersCol, err := app.FindCachedCollectionByNameOrId(core.CollectionNameSuperusers)
	if err != nil {
		return err
	}

	record := core.NewRecord(superusersCol)
	record.SetEmail(email)
	record.SetPassword(password)

	if err := app.Save(record); err != nil {
		return err
	}

	log.Printf("Bootstrapped superuser account: %s", email)
	return nil
}
