package ragstore

import (
	"maps"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

// ============================================================================
// == TYPE
// ============================================================================

type documentImplementation struct {
	dataobject.DataObject
}

// ============================================================================
// == CONSTRUCTORS
// ============================================================================

func NewDocument() DocumentInterface {
	o := &documentImplementation{}

	o.SetID(uid.HumanUid()).
		// REQUIRED:SetOwnerID("").
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).Format("Y-m-d H:i:s")).
		SetUpdatedAt(carbon.Now(carbon.UTC).Format("Y-m-d H:i:s")).
		SetSoftDeletedAt(sb.MAX_DATE)

	o.SetMetas(map[string]string{})

	return o
}

func NewDocumentFromExistingData(data map[string]string) DocumentInterface {
	o := &documentImplementation{}
	o.Hydrate(data)
	return o
}

// ============================================================================
// == METHODS
// ============================================================================

func (o *documentImplementation) IsActive() bool {
	return o.Get(COLUMN_STATUS) == DOCUMENT_STATUS_ACTIVE
}

func (o *documentImplementation) IsDeleted() bool {
	return o.Get(COLUMN_STATUS) == DOCUMENT_STATUS_DELETED
}

func (o *documentImplementation) IsInactive() bool {
	return o.Get(COLUMN_STATUS) == DOCUMENT_STATUS_INACTIVE
}

func (o *documentImplementation) IsSoftDeleted() bool {
	return o.Get(COLUMN_SOFT_DELETED_AT) != ""
}

// ============================================================================
// == GETTERS / SETTERS
// ============================================================================

func (o *documentImplementation) FileName() string {
	return o.Get(COLUMN_FILE_NAME)
}

func (o *documentImplementation) SetFileName(fileName string) DocumentInterface {
	o.Set(COLUMN_FILE_NAME, fileName)
	return o
}

func (o *documentImplementation) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *documentImplementation) CreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_CREATED_AT), carbon.UTC)
}

func (o *documentImplementation) SetCreatedAt(createdAt string) DocumentInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *documentImplementation) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *documentImplementation) SetID(id string) DocumentInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *documentImplementation) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *documentImplementation) SetMemo(memo string) DocumentInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *documentImplementation) Meta(key string) (string, error) {
	metas, err := o.Metas()
	if err != nil {
		return "", err
	}
	return metas[key], nil
}

func (o *documentImplementation) SetMeta(key string, value string) error {
	return o.UpsertMetas(map[string]string{
		key: value,
	})
}

func (o *documentImplementation) Metas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson, errJson := utils.FromJSON(metasStr, map[string]string{})
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return maputils.MapStringAnyToMapStringString(metasJson.(map[string]any)), nil
}

func (o *documentImplementation) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}

	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *documentImplementation) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	maps.Copy(currentMetas, metas)

	return o.SetMetas(currentMetas)
}

func (o *documentImplementation) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *documentImplementation) SoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_SOFT_DELETED_AT), carbon.UTC)
}

func (o *documentImplementation) SetSoftDeletedAt(softDeletedAt string) DocumentInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, softDeletedAt)
	return o
}

func (o *documentImplementation) Status() string {
	return o.Get(COLUMN_STATUS)
}

func (o *documentImplementation) SetStatus(status string) DocumentInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *documentImplementation) Text() string {
	return o.Get(COLUMN_TEXT)
}

func (o *documentImplementation) SetText(text string) DocumentInterface {
	o.Set(COLUMN_TEXT, text)
	return o
}

func (o *documentImplementation) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *documentImplementation) UpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *documentImplementation) SetUpdatedAt(updatedAt string) DocumentInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
