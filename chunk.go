package ragstore

import (
	"encoding/json"

	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/spf13/cast"
)

const CHAT_MESSAGE_STATUS_ACTIVE = "active"
const CHAT_MESSAGE_STATUS_INACTIVE = "inactive"
const CHAT_MESSAGE_STATUS_DELETED = "deleted"

type Chunk struct {
	dataobject.DataObject
}

func NewChunk() ChunkInterface {
	o := &Chunk{}
	o.SetID(uid.HumanUid())
	// REQUIRED: o.SetDocumentID("")
	o.SetCreatedAt(carbon.Now(carbon.UTC).Format("Y-m-d H:i:s"))
	o.SetUpdatedAt(carbon.Now(carbon.UTC).Format("Y-m-d H:i:s"))
	o.SetSoftDeletedAt(sb.MAX_DATETIME)
	return o
}

func NewChunkFromExistingData(data map[string]string) ChunkInterface {
	o := &Chunk{}
	o.Hydrate(data)
	return o
}

// IsSoftDeleted checks if the chunk is soft deleted
func (o *Chunk) IsSoftDeleted() bool {
	return o.SoftDeletedAt() != sb.MAX_DATETIME
}

func (o *Chunk) ChunkIndex() int {
	return cast.ToInt(o.Get(COLUMN_CHUNK_INDEX))
}

func (o *Chunk) SetChunkIndex(chunkIndex int) ChunkInterface {
	o.Set(COLUMN_CHUNK_INDEX, cast.ToString(chunkIndex))
	return o
}

func (o *Chunk) Content() string {
	return o.Get(COLUMN_CONTENT)
}

func (o *Chunk) SetContent(content string) ChunkInterface {
	o.Set(COLUMN_CONTENT, content)
	return o
}

func (o *Chunk) DocumentID() string {
	return o.Get(COLUMN_DOCUMENT_ID)
}

func (o *Chunk) SetDocumentID(id string) ChunkInterface {
	o.Set(COLUMN_DOCUMENT_ID, id)
	return o
}

func (o *Chunk) Embedding() []float32 {
	jsonStr := o.Get(COLUMN_EMBEDDING)
	var embedding []float32
	_ = json.Unmarshal([]byte(jsonStr), &embedding)
	return embedding
}

func (o *Chunk) SetEmbedding(embedding []float32) ChunkInterface {
	jsonStr, _ := json.Marshal(embedding)
	o.Set(COLUMN_EMBEDDING, string(jsonStr))
	return o
}

func (o *Chunk) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *Chunk) CreatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_CREATED_AT), carbon.UTC)
}

func (o *Chunk) SetCreatedAt(createdAt string) ChunkInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *Chunk) SetSoftDeletedAt(softDeletedAt string) ChunkInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, softDeletedAt)
	return o
}

func (o *Chunk) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *Chunk) SoftDeletedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *Chunk) Text() string {
	return o.Get(COLUMN_TEXT)
}

func (o *Chunk) SetText(text string) ChunkInterface {
	o.Set(COLUMN_TEXT, text)
	return o
}

func (o *Chunk) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *Chunk) UpdatedAtCarbon() *carbon.Carbon {
	return carbon.Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *Chunk) SetUpdatedAt(updatedAt string) ChunkInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
