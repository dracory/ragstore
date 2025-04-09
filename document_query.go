package ragstore

import (
	"errors"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/sb"
)

// documentQuery implements the DocumentQueryInterface
type documentQuery struct {
	params map[string]interface{}
}

var _ DocumentQueryInterface = (*documentQuery)(nil)

// DocumentQuery creates a new document query
func DocumentQuery() DocumentQueryInterface {
	return &documentQuery{
		params: map[string]interface{}{},
	}
}

// Validate validates the query parameters
func (q *documentQuery) Validate() error {
	if q.IsOwnerIDSet() && q.GetOwnerID() == "" {
		return errors.New("document query: owner_id cannot be empty")
	}

	if q.IsCreatedAtGteSet() && q.GetCreatedAtGte() == "" {
		return errors.New("document query: created_at_gte cannot be empty")
	}

	if q.IsCreatedAtLteSet() && q.GetCreatedAtLte() == "" {
		return errors.New("document query: created_at_lte cannot be empty")
	}

	if q.IsIDSet() && q.GetID() == "" {
		return errors.New("document query: id cannot be empty")
	}

	if q.IsIDInSet() && len(q.GetIDIn()) < 1 {
		return errors.New("document query: id_in cannot be empty array")
	}

	if q.IsLimitSet() && q.GetLimit() < 0 {
		return errors.New("document query: limit cannot be negative")
	}

	if q.IsOffsetSet() && q.GetOffset() < 0 {
		return errors.New("document query: offset cannot be negative")
	}

	if q.IsStatusSet() && q.GetStatus() == "" {
		return errors.New("document query: status cannot be empty")
	}

	if q.IsStatusInSet() && len(q.GetStatusIn()) < 1 {
		return errors.New("document query: status_in cannot be empty array")
	}

	return nil
}

func (q *documentQuery) ToSelectDataset(st *store) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if st == nil {
		return nil, []any{}, errors.New("store cannot be nil")
	}

	if err := q.Validate(); err != nil {
		return nil, []any{}, err
	}

	sql := goqu.Dialect(st.dbDriverName).From(st.tableDocument)

	// Created At filter
	if q.IsCreatedAtGteSet() {
		sql = sql.Where(goqu.C(COLUMN_CREATED_AT).Gte(q.GetCreatedAtGte()))
	}

	if q.IsCreatedAtLteSet() {
		sql = sql.Where(goqu.C(COLUMN_CREATED_AT).Lte(q.GetCreatedAtLte()))
	}

	// ID filter
	if q.IsIDSet() {
		sql = sql.Where(goqu.C(COLUMN_ID).Eq(q.GetID()))
	}

	// ID IN filter
	if q.IsIDInSet() {
		sql = sql.Where(goqu.C(COLUMN_ID).In(q.GetIDIn()))
	}

	// Status filter
	if q.IsStatusSet() {
		sql = sql.Where(goqu.C(COLUMN_STATUS).Eq(q.GetStatus()))
	}

	// Status IN filter
	if q.IsStatusInSet() {
		sql = sql.Where(goqu.C(COLUMN_STATUS).In(q.GetStatusIn()))
	}

	// Updated At filter
	if q.IsUpdatedAtGteSet() {
		sql = sql.Where(goqu.C(COLUMN_UPDATED_AT).Gte(q.GetUpdatedAtGte()))
	}

	if q.IsUpdatedAtLteSet() {
		sql = sql.Where(goqu.C(COLUMN_UPDATED_AT).Lte(q.GetUpdatedAtLte()))
	}

	if !q.IsCountOnlySet() {
		if q.IsLimitSet() {
			sql = sql.Limit(uint(q.GetLimit()))
		}

		if q.IsOffsetSet() {
			sql = sql.Offset(uint(q.GetOffset()))
		}
	}

	sortOrder := sb.DESC
	if q.IsOrderDirectionSet() {
		sortOrder = q.GetOrderDirection()
	}

	if q.IsOrderBySet() {
		if strings.EqualFold(sortOrder, sb.ASC) {
			sql = sql.Order(goqu.I(q.GetOrderBy()).Asc())
		} else {
			sql = sql.Order(goqu.I(q.GetOrderBy()).Desc())
		}
	}

	// Limit (if count only is not set)
	if !q.IsCountOnlySet() || !q.GetCountOnly() {
		if q.IsLimitSet() {
			sql = sql.Limit(uint(q.GetLimit()))
		}

		if q.IsOffsetSet() {
			sql = sql.Offset(uint(q.GetOffset()))
		}
	}

	// Sort order
	if q.IsOrderBySet() {
		sortOrder := q.GetOrderDirection()

		if strings.EqualFold(sortOrder, sb.ASC) {
			sql = sql.Order(goqu.I(q.GetOrderBy()).Asc())
		} else {
			sql = sql.Order(goqu.I(q.GetOrderBy()).Desc())
		}
	}

	// Soft delete filters

	// Only soft deleted
	if q.IsOnlySoftDeletedSet() && q.GetOnlySoftDeleted() {
		sql = sql.Where(goqu.C(COLUMN_SOFT_DELETED_AT).Lte(carbon.Now(carbon.UTC).ToDateTimeString()))
		return sql, []any{}, nil
	}

	// Include soft deleted
	if q.IsWithSoftDeletedSet() && q.GetWithSoftDeleted() {
		return sql, []any{}, nil
	}

	// Exclude soft deleted, not in the past (default)
	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	sql = sql.Where(softDeleted)

	return sql, []any{}, nil
}

// ============================================================================
// == Getters and Setters
// ============================================================================

func (q *documentQuery) IsOwnerIDSet() bool {
	return q.hasProperty("owner_id")
}

func (q *documentQuery) GetOwnerID() string {
	if q.IsOwnerIDSet() {
		return q.params["owner_id"].(string)
	}

	return ""
}

func (q *documentQuery) SetOwnerID(ownerID string) DocumentQueryInterface {
	q.params["owner_id"] = ownerID
	return q
}

func (q *documentQuery) IsCountOnlySet() bool {
	return q.hasProperty("count_only")
}

func (q *documentQuery) GetCountOnly() bool {
	if q.IsCountOnlySet() {
		return q.params["count_only"].(bool)
	}

	return false
}

func (q *documentQuery) SetCountOnly(countOnly bool) DocumentQueryInterface {
	q.params["count_only"] = countOnly
	return q
}

func (q *documentQuery) IsCreatedAtGteSet() bool {
	return q.hasProperty("created_at_gte")
}

func (q *documentQuery) GetCreatedAtGte() string {
	if q.IsCreatedAtGteSet() {
		return q.params["created_at_gte"].(string)
	}

	return ""
}

func (q *documentQuery) SetCreatedAtGte(createdAtGte string) DocumentQueryInterface {
	q.params["created_at_gte"] = createdAtGte
	return q
}

func (q *documentQuery) IsCreatedAtLteSet() bool {
	return q.hasProperty("created_at_lte")
}

func (q *documentQuery) GetCreatedAtLte() string {
	if q.IsCreatedAtLteSet() {
		return q.params["created_at_lte"].(string)
	}

	return ""
}

func (q *documentQuery) SetCreatedAtLte(createdAtLte string) DocumentQueryInterface {
	q.params["created_at_lte"] = createdAtLte
	return q
}

func (q *documentQuery) IsIDSet() bool {
	return q.hasProperty("id")
}
func (q *documentQuery) GetID() string {
	if q.IsIDSet() {
		return q.params["id"].(string)
	}

	return ""
}

func (q *documentQuery) SetID(id string) DocumentQueryInterface {
	q.params["id"] = id
	return q
}

func (q *documentQuery) IsIDInSet() bool {
	return q.hasProperty("id_in")
}

func (q *documentQuery) GetIDIn() []string {
	if q.IsIDInSet() {
		return q.params["id_in"].([]string)
	}

	return []string{}
}

func (q *documentQuery) SetIDIn(idIn []string) DocumentQueryInterface {
	q.params["id_in"] = idIn
	return q
}

func (q *documentQuery) IsLimitSet() bool {
	return q.hasProperty("limit")
}

func (q *documentQuery) GetLimit() int {
	if q.IsLimitSet() {
		return q.params["limit"].(int)
	}

	return 0
}

func (q *documentQuery) SetLimit(limit int) DocumentQueryInterface {
	q.params["limit"] = limit
	return q
}

func (q *documentQuery) IsOffsetSet() bool {
	return q.hasProperty("offset")
}

func (q *documentQuery) GetOffset() int {
	if q.IsOffsetSet() {
		return q.params["offset"].(int)
	}

	return 0
}

func (q *documentQuery) SetOffset(offset int) DocumentQueryInterface {
	q.params["offset"] = offset
	return q
}

func (q *documentQuery) IsOnlySoftDeletedSet() bool {
	return q.hasProperty("only_soft_deleted")
}

func (q *documentQuery) GetOnlySoftDeleted() bool {
	if q.IsOnlySoftDeletedSet() {
		return q.params["only_soft_deleted"].(bool)
	}

	return false
}

func (q *documentQuery) SetOnlySoftDeleted(onlySoftDeleted bool) DocumentQueryInterface {
	q.params["only_soft_deleted"] = onlySoftDeleted
	return q
}

func (q *documentQuery) IsOrderDirectionSet() bool {
	return q.hasProperty("order_direction")
}

func (q *documentQuery) GetOrderDirection() string {
	if q.IsOrderDirectionSet() {
		return q.params["order_direction"].(string)
	}

	return ""
}

func (q *documentQuery) SetOrderDirection(orderDirection string) DocumentQueryInterface {
	q.params["order_direction"] = orderDirection
	return q
}

func (q *documentQuery) IsOrderBySet() bool {
	return q.hasProperty("order_by")
}

func (q *documentQuery) GetOrderBy() string {
	if q.IsOrderBySet() {
		return q.params["order_by"].(string)
	}

	return ""
}

func (q *documentQuery) SetOrderBy(orderBy string) DocumentQueryInterface {
	q.params["order_by"] = orderBy
	return q
}

func (q *documentQuery) IsStatusSet() bool {
	return q.hasProperty("status")
}

func (q *documentQuery) GetStatus() string {
	if q.IsStatusSet() {
		return q.params["status"].(string)
	}

	return ""
}

func (q *documentQuery) SetStatus(status string) DocumentQueryInterface {
	q.params["status"] = status
	return q
}

func (q *documentQuery) IsStatusInSet() bool {
	return q.hasProperty("status_in")
}

func (q *documentQuery) GetStatusIn() []string {
	if q.IsStatusInSet() {
		return q.params["status_in"].([]string)
	}

	return []string{}
}

func (q *documentQuery) SetStatusIn(statusIn []string) DocumentQueryInterface {
	q.params["status_in"] = statusIn
	return q
}

func (q *documentQuery) IsUpdatedAtGteSet() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *documentQuery) GetUpdatedAtGte() string {
	if q.IsUpdatedAtGteSet() {
		return q.params["updated_at_gte"].(string)
	}

	return ""
}

func (q *documentQuery) SetUpdatedAtGte(updatedAt string) DocumentQueryInterface {
	q.params["updated_at_gte"] = updatedAt
	return q
}

func (q *documentQuery) IsUpdatedAtLteSet() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *documentQuery) GetUpdatedAtLte() string {
	if q.IsUpdatedAtLteSet() {
		return q.params["updated_at_lte"].(string)
	}

	return ""
}

func (q *documentQuery) SetUpdatedAtLte(updatedAt string) DocumentQueryInterface {
	q.params["updated_at_lte"] = updatedAt
	return q
}

func (q *documentQuery) IsWithSoftDeletedSet() bool {
	return q.hasProperty("with_soft_deleted")
}

func (q *documentQuery) GetWithSoftDeleted() bool {
	if q.IsWithSoftDeletedSet() {
		return q.params["with_soft_deleted"].(bool)
	}

	return false
}

func (q *documentQuery) SetWithSoftDeleted(withSoftDeleted bool) DocumentQueryInterface {
	q.params["with_soft_deleted"] = withSoftDeleted
	return q
}

func (q *documentQuery) hasProperty(key string) bool {
	return q.params[key] != nil
}
