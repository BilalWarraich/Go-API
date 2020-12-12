package db

import (
	"context"

	"github.com/jinzhu/gorm"
)

// transaction start a transaction as a block,
// return error will rollback, otherwise to commit.
func transaction(ctx context.Context, db *gorm.DB, fc func(tx *gorm.DB) error) (err error) {
	panicked := true
	tx := db.BeginTx(ctx, nil)
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	err = fc(tx)
	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
