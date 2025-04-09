package ragstore

import "github.com/dromara/carbon/v2"

type DocumentInterface interface {
	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool

	ID() string
	SetID(id string) DocumentInterface

	Status() string
	SetStatus(status string) DocumentInterface

	FileName() string
	SetFileName(fileName string) DocumentInterface

	Text() string
	SetText(text string) DocumentInterface

	Memo() string
	SetMemo(memo string) DocumentInterface

	Meta(key string) (string, error)
	SetMeta(key string, value string) error

	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	UpsertMetas(metas map[string]string) error

	CreatedAt() string
	CreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) DocumentInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) DocumentInterface

	UpdatedAt() string
	UpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) DocumentInterface
}
