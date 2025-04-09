package ragstore

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"
)

// ============================================================================
// == TYPE
// ============================================================================

type store struct {
	tableDocument      string
	tableDocumentChunk string
	db                 *sql.DB
	dbDriverName       string
	automigrateEnabled bool
	debugEnabled       bool
	logger             *slog.Logger
}

// ============================================================================
// == INTERFACE
// ============================================================================

var _ StoreInterface = (*store)(nil) // verify it extends the interface

// ============================================================================
// == PUBLIC METHODS
// ============================================================================

// AutoMigrate auto migrate
func (store *store) AutoMigrate() error {
	if store.db == nil {
		return errors.New("chatstore: database is nil")
	}

	sqls := []string{
		store.sqlDocumentTableCreate(),
		store.sqlDocumentChunkTableCreate(),
	}

	for _, sql := range sqls {
		if sql == "" {
			return errors.New("table create sql is empty")
		}

		_, err := store.db.Exec(sql)

		if err != nil {
			return err
		}
	}

	return nil
}

// EnableDebug - enables the debug option
func (st *store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

// GetLogger - returns the logger
func (st *store) GetLogger() *slog.Logger {
	if st.logger == nil {
		st.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return st.logger
}

// SetLogger - sets the logger
func (st *store) SetLogger(logger *slog.Logger) {
	st.logger = logger
}
