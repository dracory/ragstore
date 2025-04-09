package ragstore

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"

	"github.com/gouniverse/sb"
)

// NewStoreOptions define the options for creating a new block store
type NewStoreOptions struct {
	TableDocumentName      string
	TableDocumentChunkName string
	DB                     *sql.DB
	DbDriverName           string
	AutomigrateEnabled     bool
	DebugEnabled           bool
	Logger                 *slog.Logger
}

// NewStore creates a new block store
func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	if opts.TableDocumentName == "" {
		return nil, errors.New("chat store: TableDocumentName is required")
	}

	if opts.TableDocumentChunkName == "" {
		return nil, errors.New("chat store: TableDocumentChunkName is required")
	}

	if opts.DB == nil {
		return nil, errors.New("shop store: DB is required")
	}

	if opts.DbDriverName == "" {
		opts.DbDriverName = sb.DatabaseDriverName(opts.DB)
	}

	if opts.Logger == nil {
		opts.Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	store := &store{
		tableDocument:      opts.TableDocumentName,
		tableDocumentChunk: opts.TableDocumentChunkName,
		automigrateEnabled: opts.AutomigrateEnabled,
		db:                 opts.DB,
		dbDriverName:       opts.DbDriverName,
		debugEnabled:       opts.DebugEnabled,
		logger:             opts.Logger,
	}

	if store.automigrateEnabled {
		err := store.AutoMigrate()

		if err != nil {
			return nil, err
		}
	}

	return store, nil
}
