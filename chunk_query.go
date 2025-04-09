package ragstore

import (
	"errors"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/sb"
)

// chunkQuery implements the ChunkQueryInterface
type chunkQuery struct {
	params map[string]interface{}
}

var _ ChunkQueryInterface = (*chunkQuery)(nil)

// ChunkQuery creates a new chunk query
func ChunkQuery() ChunkQueryInterface {
	return &chunkQuery{
		params: map[string]interface{}{},
	}
}

func (q *chunkQuery) ToSelectDataset(st *store) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if st == nil {
		return nil, []any{}, errors.New("store cannot be nil")
	}

	if err := q.Validate(); err != nil {
		return nil, []any{}, err
	}

	sql := goqu.Dialect(st.dbDriverName).From(st.tableDocumentChunk)

	// Document ID filter
	if q.IsDocumentIDSet() {
		sql = sql.Where(goqu.C(COLUMN_DOCUMENT_ID).Eq(q.GetDocumentID()))
	}

	// Document ID IN filter
	if q.IsDocumentIDInSet() {
		sql = sql.Where(goqu.C(COLUMN_DOCUMENT_ID).In(q.GetDocumentIDIn()))
	}

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

// Validate validates the query parameters
func (q *chunkQuery) Validate() error {

	if q.IsCreatedAtGteSet() && q.GetCreatedAtGte() == "" {
		return errors.New("chunk query: created_at_gte cannot be empty")
	}

	if q.IsCreatedAtLteSet() && q.GetCreatedAtLte() == "" {
		return errors.New("chunk query: created_at_lte cannot be empty")
	}

	if q.IsDocumentIDSet() && q.GetDocumentID() == "" {
		return errors.New("chunk query: document_id cannot be empty")
	}

	if q.IsDocumentIDInSet() && len(q.GetDocumentIDIn()) < 1 {
		return errors.New("chunk query: document_id_in cannot be empty array")
	}

	if q.IsIDSet() && q.GetID() == "" {
		return errors.New("chunk query: id cannot be empty")
	}

	if q.IsIDInSet() && len(q.GetIDIn()) < 1 {
		return errors.New("chunk query: id_in cannot be empty array")
	}

	if q.IsIDNotInSet() && len(q.GetIDNotIn()) < 1 {
		return errors.New("chunk query: id_not_in cannot be empty array")
	}

	if q.IsLimitSet() && q.GetLimit() < 0 {
		return errors.New("chunk query: limit cannot be negative")
	}

	if q.IsOffsetSet() && q.GetOffset() < 0 {
		return errors.New("chunk query: offset cannot be negative")
	}

	if q.IsOrderBySet() && q.GetOrderBy() == "" {
		return errors.New("chunk query: order_by cannot be empty")
	}

	if q.IsOrderDirectionSet() && q.GetOrderDirection() == "" {
		return errors.New("chunk query: order_direction cannot be empty")
	}

	if q.IsStatusInSet() && len(q.GetStatusIn()) < 1 {
		return errors.New("chunk query: status_in cannot be empty array")
	}

	return nil
}

// ============================================================================
// == Getters and Setters
// ============================================================================

func (q *chunkQuery) IsCountOnlySet() bool {
	return q.hasProperty("count_only")
}

func (q *chunkQuery) GetCountOnly() bool {
	if q.IsCountOnlySet() {
		return q.params["count_only"].(bool)
	}

	return false
}

func (q *chunkQuery) SetCountOnly(countOnly bool) ChunkQueryInterface {
	q.params["count_only"] = countOnly
	return q
}

func (q *chunkQuery) IsCreatedAtGteSet() bool {
	return q.hasProperty("created_at_gte")
}

func (q *chunkQuery) GetCreatedAtGte() string {
	if q.IsCreatedAtGteSet() {
		return q.params["created_at_gte"].(string)
	}

	return ""
}

func (q *chunkQuery) SetCreatedAtGte(createdAtGte string) ChunkQueryInterface {
	q.params["created_at_gte"] = createdAtGte
	return q
}

func (q *chunkQuery) IsCreatedAtLteSet() bool {
	return q.hasProperty("created_at_lte")
}

func (q *chunkQuery) GetCreatedAtLte() string {
	if q.IsCreatedAtLteSet() {
		return q.params["created_at_lte"].(string)
	}

	return ""
}

func (q *chunkQuery) SetCreatedAtLte(createdAtLte string) ChunkQueryInterface {
	q.params["created_at_lte"] = createdAtLte
	return q
}

func (q *chunkQuery) IsDocumentIDSet() bool {
	return q.hasProperty("document_id")
}

func (q *chunkQuery) GetDocumentID() string {
	if q.IsDocumentIDSet() {
		return q.params["document_id"].(string)
	}

	return ""
}

func (q *chunkQuery) SetDocumentID(chatID string) ChunkQueryInterface {
	q.params["document_id"] = chatID
	return q
}

func (q *chunkQuery) IsDocumentIDInSet() bool {
	return q.hasProperty("document_id_in")
}

func (q *chunkQuery) GetDocumentIDIn() []string {
	if q.IsDocumentIDInSet() {
		return q.params["document_id_in"].([]string)
	}

	return []string{}
}

func (q *chunkQuery) SetDocumentIDIn(documentIDIn []string) ChunkQueryInterface {
	q.params["document_id_in"] = documentIDIn
	return q
}

func (q *chunkQuery) IsIDSet() bool {
	return q.hasProperty("id")
}

func (q *chunkQuery) GetID() string {
	if q.IsIDSet() {
		return q.params["id"].(string)
	}

	return ""
}

func (q *chunkQuery) SetID(id string) ChunkQueryInterface {
	q.params["id"] = id
	return q
}

func (q *chunkQuery) IsIDInSet() bool {
	return q.hasProperty("id_in")
}

func (q *chunkQuery) GetIDIn() []string {
	if q.IsIDInSet() {
		return q.params["id_in"].([]string)
	}

	return []string{}
}

func (q *chunkQuery) SetIDIn(idIn []string) ChunkQueryInterface {
	q.params["id_in"] = idIn
	return q
}

func (q *chunkQuery) IsLimitSet() bool {
	return q.hasProperty("limit")
}

func (q *chunkQuery) GetLimit() int {
	if q.IsLimitSet() {
		return q.params["limit"].(int)
	}

	return 0
}

func (q *chunkQuery) IsIDNotInSet() bool {
	return q.hasProperty("id_not_in")
}

func (q *chunkQuery) GetIDNotIn() []string {
	if q.IsIDNotInSet() {
		return q.params["id_not_in"].([]string)
	}

	return []string{}
}

func (q *chunkQuery) SetIDNotIn(idNotIn []string) ChunkQueryInterface {
	q.params["id_not_in"] = idNotIn
	return q
}

func (q *chunkQuery) SetLimit(limit int) ChunkQueryInterface {
	q.params["limit"] = limit
	return q
}

func (q *chunkQuery) IsOffsetSet() bool {
	return q.hasProperty("offset")
}

func (q *chunkQuery) GetOffset() int {
	if q.IsOffsetSet() {
		return q.params["offset"].(int)
	}

	return 0
}

func (q *chunkQuery) SetOffset(offset int) ChunkQueryInterface {
	q.params["offset"] = offset
	return q
}

func (q *chunkQuery) IsOnlySoftDeletedSet() bool {
	return q.hasProperty("only_soft_deleted")
}

func (q *chunkQuery) GetOnlySoftDeleted() bool {
	if q.IsOnlySoftDeletedSet() {
		return q.params["only_soft_deleted"].(bool)
	}

	return false
}

func (q *chunkQuery) SetOnlySoftDeleted(onlySoftDeleted bool) ChunkQueryInterface {
	q.params["only_soft_deleted"] = onlySoftDeleted
	return q
}

func (q *chunkQuery) IsOrderDirectionSet() bool {
	return q.hasProperty("order_direction")
}

func (q *chunkQuery) GetOrderDirection() string {
	if q.IsOrderDirectionSet() {
		return q.params["order_direction"].(string)
	}

	return ""
}

func (q *chunkQuery) SetOrderDirection(orderDirection string) ChunkQueryInterface {
	q.params["order_direction"] = orderDirection
	return q
}

func (q *chunkQuery) IsOrderBySet() bool {
	return q.hasProperty("order_by")
}

func (q *chunkQuery) GetOrderBy() string {
	if q.IsOrderBySet() {
		return q.params["order_by"].(string)
	}

	return ""
}

func (q *chunkQuery) SetOrderBy(orderBy string) ChunkQueryInterface {
	q.params["order_by"] = orderBy
	return q
}

func (q *chunkQuery) IsSenderIDSet() bool {
	return q.hasProperty("sender_id")
}

func (q *chunkQuery) GetSenderID() string {
	if q.IsSenderIDSet() {
		return q.params["sender_id"].(string)
	}

	return ""
}

func (q *chunkQuery) SetSenderID(senderID string) ChunkQueryInterface {
	q.params["sender_id"] = senderID
	return q
}

func (q *chunkQuery) IsStatusSet() bool {
	return q.hasProperty("status")
}

func (q *chunkQuery) GetStatus() string {
	if q.IsStatusSet() {
		return q.params["status"].(string)
	}

	return ""
}

func (q *chunkQuery) SetStatus(status string) ChunkQueryInterface {
	q.params["status"] = status
	return q
}

func (q *chunkQuery) IsStatusInSet() bool {
	return q.hasProperty("status_in")
}

func (q *chunkQuery) GetStatusIn() []string {
	if q.IsStatusInSet() {
		return q.params["status_in"].([]string)
	}

	return []string{}
}

func (q *chunkQuery) SetStatusIn(statusIn []string) ChunkQueryInterface {
	q.params["status_in"] = statusIn
	return q
}

func (q *chunkQuery) IsUpdatedAtGteSet() bool {
	return q.hasProperty("updated_at_gte")
}

func (q *chunkQuery) GetUpdatedAtGte() string {
	if q.IsUpdatedAtGteSet() {
		return q.params["updated_at_gte"].(string)
	}

	return ""
}

func (q *chunkQuery) SetUpdatedAtGte(updatedAt string) ChunkQueryInterface {
	q.params["updated_at_gte"] = updatedAt
	return q
}

func (q *chunkQuery) IsUpdatedAtLteSet() bool {
	return q.hasProperty("updated_at_lte")
}

func (q *chunkQuery) GetUpdatedAtLte() string {
	if q.IsUpdatedAtLteSet() {
		return q.params["updated_at_lte"].(string)
	}

	return ""
}

func (q *chunkQuery) SetUpdatedAtLte(updatedAt string) ChunkQueryInterface {
	q.params["updated_at_lte"] = updatedAt
	return q
}

func (q *chunkQuery) IsWithSoftDeletedSet() bool {
	return q.hasProperty("with_soft_deleted")
}

func (q *chunkQuery) GetWithSoftDeleted() bool {
	if q.IsWithSoftDeletedSet() {
		return q.params["with_soft_deleted"].(bool)
	}

	return false
}

func (q *chunkQuery) SetWithSoftDeleted(withSoftDeleted bool) ChunkQueryInterface {
	q.params["with_soft_deleted"] = withSoftDeleted
	return q
}

func (q *chunkQuery) hasProperty(key string) bool {
	return q.params[key] != nil
}
