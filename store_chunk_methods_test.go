package ragstore

import (
	"testing"

	"github.com/gouniverse/sb"
	_ "modernc.org/sqlite"
)

func TestStore_ChunkCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Create multiple chunks
	chunk1 := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	chunk2 := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(2).
		SetContent("chunk 2").
		SetEmbedding([]float32{4.0, 5.0, 6.0})

	chunk3 := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(3).
		SetContent("chunk 3").
		SetEmbedding([]float32{7.0, 8.0, 9.0})

	err = store.ChunkCreate(chunk1)
	if err != nil {
		t.Fatal("unexpected error creating chunk1:", err)
	}

	err = store.ChunkCreate(chunk2)
	if err != nil {
		t.Fatal("unexpected error creating chunk2:", err)
	}

	err = store.ChunkCreate(chunk3)
	if err != nil {
		t.Fatal("unexpected error creating chunk3:", err)
	}

	// Test counting all chunks
	allCount, err := store.ChunkCount(ChunkQuery())
	if err != nil {
		t.Fatal("unexpected error counting all chunks:", err)
	}

	if allCount != 3 {
		t.Fatalf("Expected count of 3 documents, got %d", allCount)
	}

	// Test counting by monitor ID
	chunkCount, err := store.ChunkCount(ChunkQuery().SetDocumentID(testDocument_O1))
	if err != nil {
		t.Fatal("unexpected error counting monitor chunks:", err)
	}

	if chunkCount != 3 {
		t.Fatalf("Expected count of 3 chunks, got %d", chunkCount)
	}
}

func TestStore_ChunkCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStore_ChunkCreateDuplicate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error on first create:", err)
	}

	// Try to create the same chunk again
	err = store.ChunkCreate(chunk)
	if err == nil {
		t.Fatal("expected error for duplicate chunk, but got nil")
	}
}

func TestStore_ChunkFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	chunkFound, errFind := store.ChunkFindByID(chunk.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if chunkFound == nil {
		t.Fatal("Chunk MUST NOT be nil")
	}

	if chunkFound.ID() != chunk.ID() {
		t.Fatal("IDs do not match")
	}
}

func TestStore_ChunkFindByIDNotFound(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunkFound, errFind := store.ChunkFindByID("non-existent-id")

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if chunkFound != nil {
		t.Fatal("Chunk MUST be nil for non-existent ID")
	}
}

// func TestStoreDocumentUpdateNonExistent(t *testing.T) {
// 	store, err := initStore(":memory:")

// 	if err != nil {
// 		t.Fatal("unexpected error:", err)
// 	}

// 	document := NewDocument().
// 		SetStatus(CHAT_STATUS_ACTIVE)

// 	// Try to update a non-existent document
// 	err = store.DocumentUpdate(document)
// 	if err == nil {
// 		t.Fatal("expected error for updating non-existent document, but got nil")
// 	}
// }

func TestStore_ChunkDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Delete the chunk
	err = store.ChunkDelete(chunk)
	if err != nil {
		t.Fatal("unexpected error on delete:", err)
	}

	// Verify the chunk is deleted
	deletedChunk, err := store.ChunkFindByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error finding deleted chunk:", err)
	}

	if deletedChunk != nil {
		t.Fatal("Chunk should be nil after deletion")
	}
}

func TestStore_ChunkDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Delete the chunk by ID
	err = store.ChunkDeleteByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error on delete by ID:", err)
	}

	// Verify the chunk is deleted
	deletedChunk, err := store.ChunkFindByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error finding deleted chunk:", err)
	}

	if deletedChunk != nil {
		t.Fatal("Chunk should be nil after deletion by ID")
	}
}

func TestStore_ChunkList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Create multiple chunks
	chunk1 := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	chunk2 := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(2).
		SetContent("chunk 2").
		SetEmbedding([]float32{4.0, 5.0, 6.0})

	err = store.ChunkCreate(chunk1)
	if err != nil {
		t.Fatal("unexpected error creating chunk1:", err)
	}

	err = store.ChunkCreate(chunk2)
	if err != nil {
		t.Fatal("unexpected error creating chunk2:", err)
	}

	// Test listing all chunks
	allChunks, err := store.ChunkList(ChunkQuery())
	if err != nil {
		t.Fatal("unexpected error listing all chunks:", err)
	}

	if len(allChunks) != 2 {
		t.Fatalf("Expected 2 chunks, got %d", len(allChunks))
	}

	// Test filtering by document ID
	documentChunks, err := store.ChunkList(ChunkQuery().SetDocumentID(testDocument_O1))
	if err != nil {
		t.Fatal("unexpected error listing document chunks:", err)
	}

	if len(documentChunks) != 2 {
		t.Fatalf("Expected 2 chunks for document %s, got %d", testDocument_O1, len(documentChunks))
	}

	// Test limit and offset
	limitedChunks, err := store.ChunkList(ChunkQuery().SetLimit(1))
	if err != nil {
		t.Fatal("unexpected error listing limited chunks:", err)
	}

	if len(limitedChunks) != 1 {
		t.Fatalf("Expected 1 chunk with limit, got %d", len(limitedChunks))
	}

	offsetChunks, err := store.ChunkList(ChunkQuery().SetOffset(1).SetLimit(2))
	if err != nil {
		t.Fatal("unexpected error listing offset chunks:", err)
	}

	if len(offsetChunks) != 1 {
		t.Fatalf("Expected 1 chunk with offset, got %d", len(offsetChunks))
	}
}

func TestStore_ChunkSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Soft delete the chunk
	err = store.ChunkSoftDelete(chunk)
	if err != nil {
		t.Fatal("unexpected error on soft delete:", err)
	}

	// Verify the chunk is soft deleted (not found by default)
	softDeletedChunk, err := store.ChunkFindByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error finding soft deleted chunk:", err)
	}

	if softDeletedChunk != nil {
		t.Fatal("Chunk should not be found after soft deletion")
	}

	// Verify the chunk can be found when including soft deleted
	query := ChunkQuery().
		SetWithSoftDeleted(true).
		SetID(chunk.ID()).
		SetLimit(1)

	chunkFindWithDeleted, err := store.ChunkList(query)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(chunkFindWithDeleted) == 0 {
		t.Fatal("Chunk should be found when including soft deleted")
	}

	if chunkFindWithDeleted[0].SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Chunk should be soft deleted, but SoftDeletedAt is MAX_DATETIME:", chunkFindWithDeleted[0].SoftDeletedAt())
	}

	if !chunkFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Chunk should be marked as soft deleted")
	}
}

func TestStore_ChunkSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Soft delete the chunk by ID
	err = store.ChunkSoftDeleteByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error on soft delete by ID:", err)
	}

	// Verify the chunk is soft deleted (not found by default)
	softDeletedChunk, err := store.ChunkFindByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error finding soft deleted chunk:", err)
	}

	if softDeletedChunk != nil {
		t.Fatal("Chunk should not be found after soft deletion by ID")
	}

	// Verify the chunk can be found when including soft deleted
	query := ChunkQuery().
		SetWithSoftDeleted(true).
		SetID(chunk.ID()).
		SetLimit(1)

	chunkFindWithDeleted, err := store.ChunkList(query)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(chunkFindWithDeleted) == 0 {
		t.Fatal("Chunk should be found when including soft deleted")
	}

	if chunkFindWithDeleted[0].SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Chunk should be soft deleted, but SoftDeletedAt is MAX_DATETIME:", chunkFindWithDeleted[0].SoftDeletedAt())
	}

	if !chunkFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Chunk should be marked as soft deleted")
	}
}

func TestStore_ChunkUpdate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	chunk := NewChunk().
		SetDocumentID(testDocument_O1).
		SetChunkIndex(1).
		SetContent("chunk 1").
		SetEmbedding([]float32{1.0, 2.0, 3.0})

	err = store.ChunkCreate(chunk)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Update the chunk
	chunk.SetDocumentID(testDocument_O2)

	err = store.ChunkUpdate(chunk)
	if err != nil {
		t.Fatal("unexpected error on update:", err)
	}

	// Verify the update
	updatedChunk, err := store.ChunkFindByID(chunk.ID())
	if err != nil {
		t.Fatal("unexpected error finding updated chunk:", err)
	}

	if updatedChunk.DocumentID() != testDocument_O2 {
		t.Fatalf("Text not updated. Expected 'Chunk 2', got '%s'", updatedChunk.DocumentID())
	}
}
