package state

import "github.com/abdullin/fdb-go/fdb"

// abdullin: generated from DSLv7

// Deletes UserEmailIndex
func DeleteUserEmailIndex(tr fdb.Transaction) {
	tr.ClearRange(db.Sub("index"))
}
