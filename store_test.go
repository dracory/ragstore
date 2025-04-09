package ragstore

import (
	"database/sql"
	"errors"
	"os"

	"github.com/gouniverse/utils"
)

const testDocument_O1 = "00000000000000000000000000000010"
const testDocument_O2 = "00000000000000000000000000000020"

// const testChunk_O1 = "00000000000000000000000000000050"
// const testChunk_O2 = "00000000000000000000000000000060"

func initDB(filepath string) *sql.DB {
	if filepath != ":memory:" && utils.FileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			panic(err)
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func initStore(filepath string) (StoreInterface, error) {
	db := initDB(filepath)

	store, err := NewStore(NewStoreOptions{
		DB:                     db,
		TableDocumentName:      "document_table",
		TableDocumentChunkName: "document_chunk_table",
		AutomigrateEnabled:     true,
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}
