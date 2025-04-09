package ragstore

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(enabled bool)

	DocumentCount(options DocumentQueryInterface) (int64, error)
	DocumentCreate(chat DocumentInterface) error
	DocumentDelete(chat DocumentInterface) error
	DocumentDeleteByID(id string) error
	DocumentFindByID(id string) (DocumentInterface, error)
	DocumentList(options DocumentQueryInterface) ([]DocumentInterface, error)
	DocumentSoftDelete(chat DocumentInterface) error
	DocumentSoftDeleteByID(id string) error
	DocumentUpdate(chat DocumentInterface) error

	ChunkCount(options ChunkQueryInterface) (int64, error)
	ChunkCreate(message ChunkInterface) error
	ChunkDelete(message ChunkInterface) error
	ChunkDeleteByID(id string) error
	ChunkFindByID(id string) (ChunkInterface, error)
	ChunkList(options ChunkQueryInterface) ([]ChunkInterface, error)
	ChunkSoftDelete(message ChunkInterface) error
	ChunkSoftDeleteByID(id string) error
	ChunkUpdate(message ChunkInterface) error
}
