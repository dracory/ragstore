package ragstore

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
)

type ChunkInterface interface {
	dataobject.DataObjectInterface

	IsSoftDeleted() bool

	DocumentID() string
	SetDocumentID(chatID string) ChunkInterface

	ChunkIndex() int
	SetChunkIndex(chunkIndex int) ChunkInterface

	Content() string
	SetContent(content string) ChunkInterface

	Embedding() []float32
	SetEmbedding(embedding []float32) ChunkInterface

	CreatedAt() string
	CreatedAtCarbon() *carbon.Carbon
	SetCreatedAt(createdAt string) ChunkInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() *carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) ChunkInterface

	UpdatedAt() string
	UpdatedAtCarbon() *carbon.Carbon
	SetUpdatedAt(updatedAt string) ChunkInterface
}
