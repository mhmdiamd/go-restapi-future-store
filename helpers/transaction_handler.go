package helpers

import (
	"database/sql"

	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errorRollback := tx.Rollback()
		exceptions.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		exceptions.PanicIfError(errorCommit)
	}
}
