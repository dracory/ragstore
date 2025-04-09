package ragstore

import "github.com/doug-martin/goqu/v9"

// ChunkQueryInterface defines the interface for querying chunks
type ChunkQueryInterface interface {
	// Validation method
	Validate() error

	// Basic query methods
	IsCreatedAtGteSet() bool
	GetCreatedAtGte() string
	SetCreatedAtGte(createdAt string) ChunkQueryInterface

	IsCreatedAtLteSet() bool
	GetCreatedAtLte() string
	SetCreatedAtLte(createdAt string) ChunkQueryInterface

	IsIDSet() bool
	GetID() string
	SetID(id string) ChunkQueryInterface

	IsIDInSet() bool
	GetIDIn() []string
	SetIDIn(ids []string) ChunkQueryInterface

	IsIDNotInSet() bool
	GetIDNotIn() []string
	SetIDNotIn(ids []string) ChunkQueryInterface

	IsLimitSet() bool
	GetLimit() int
	SetLimit(limit int) ChunkQueryInterface

	IsDocumentIDSet() bool
	GetDocumentID() string
	SetDocumentID(documentID string) ChunkQueryInterface

	IsDocumentIDInSet() bool
	GetDocumentIDIn() []string
	SetDocumentIDIn(chatIDs []string) ChunkQueryInterface

	IsOffsetSet() bool
	GetOffset() int
	SetOffset(offset int) ChunkQueryInterface

	IsOrderBySet() bool
	GetOrderBy() string
	SetOrderBy(orderBy string) ChunkQueryInterface

	IsOrderDirectionSet() bool
	GetOrderDirection() string
	SetOrderDirection(orderDirection string) ChunkQueryInterface

	// Count related methods
	IsCountOnlySet() bool
	GetCountOnly() bool
	SetCountOnly(countOnly bool) ChunkQueryInterface

	// Soft delete related query methods
	IsWithSoftDeletedSet() bool
	GetWithSoftDeleted() bool
	SetWithSoftDeleted(withSoftDeleted bool) ChunkQueryInterface

	IsOnlySoftDeletedSet() bool
	GetOnlySoftDeleted() bool
	SetOnlySoftDeleted(onlySoftDeleted bool) ChunkQueryInterface

	ToSelectDataset(store *store) (selectDataset *goqu.SelectDataset, columns []any, err error)
}
