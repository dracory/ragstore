package ragstore

import "github.com/doug-martin/goqu/v9"

// DocumentQueryInterface defines the interface for querying chats
type DocumentQueryInterface interface {
	// Validation method
	Validate() error

	// Count related methods
	IsCountOnlySet() bool
	GetCountOnly() bool
	SetCountOnly(countOnly bool) DocumentQueryInterface

	// Soft delete related query methods
	IsWithSoftDeletedSet() bool
	GetWithSoftDeleted() bool
	SetWithSoftDeleted(withSoftDeleted bool) DocumentQueryInterface

	IsOnlySoftDeletedSet() bool
	GetOnlySoftDeleted() bool
	SetOnlySoftDeleted(onlySoftDeleted bool) DocumentQueryInterface

	// Dataset conversion methods
	ToSelectDataset(store *store) (selectDataset *goqu.SelectDataset, columns []any, err error)

	// Field query methods

	IsCreatedAtGteSet() bool
	GetCreatedAtGte() string
	SetCreatedAtGte(createdAt string) DocumentQueryInterface

	IsCreatedAtLteSet() bool
	GetCreatedAtLte() string
	SetCreatedAtLte(createdAt string) DocumentQueryInterface

	IsIDSet() bool
	GetID() string
	SetID(id string) DocumentQueryInterface

	IsIDInSet() bool
	GetIDIn() []string
	SetIDIn(ids []string) DocumentQueryInterface

	IsLimitSet() bool
	GetLimit() int
	SetLimit(limit int) DocumentQueryInterface

	IsOffsetSet() bool
	GetOffset() int
	SetOffset(offset int) DocumentQueryInterface

	IsOrderBySet() bool
	GetOrderBy() string
	SetOrderBy(orderBy string) DocumentQueryInterface

	IsOrderDirectionSet() bool
	GetOrderDirection() string
	SetOrderDirection(orderDirection string) DocumentQueryInterface

	IsStatusSet() bool
	GetStatus() string
	SetStatus(status string) DocumentQueryInterface
	SetStatusIn(statuses []string) DocumentQueryInterface

	IsUpdatedAtGteSet() bool
	GetUpdatedAtGte() string
	SetUpdatedAtGte(updatedAt string) DocumentQueryInterface

	IsUpdatedAtLteSet() bool
	GetUpdatedAtLte() string
	SetUpdatedAtLte(updatedAt string) DocumentQueryInterface
}
