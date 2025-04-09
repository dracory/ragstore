package ragstore

import (
	"testing"

	"github.com/gouniverse/sb"
	_ "modernc.org/sqlite"
)

func TestStore_DocumentCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Create multiple documents
	document1 := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	document2 := NewDocument().
		SetStatus(DOCUMENT_STATUS_INACTIVE).
		SetFileName("test2.txt").
		SetText("This is second test document.")

	document3 := NewDocument().
		SetStatus(DOCUMENT_STATUS_INACTIVE).
		SetFileName("test3.txt").
		SetText("This is third test document.")

	err = store.DocumentCreate(document1)
	if err != nil {
		t.Fatal("unexpected error creating document1:", err)
	}

	err = store.DocumentCreate(document2)
	if err != nil {
		t.Fatal("unexpected error creating document2:", err)
	}

	err = store.DocumentCreate(document3)
	if err != nil {
		t.Fatal("unexpected error creating document3:", err)
	}

	// Test counting all documents
	allCount, err := store.DocumentCount(DocumentQuery())
	if err != nil {
		t.Fatal("unexpected error counting all documents:", err)
	}

	if allCount != 3 {
		t.Fatalf("Expected count of 3 documents, got %d", allCount)
	}

	// Test counting by monitor ID
	// documentCount, err := store.DocumentCount(DocumentQuery().SetOwnerID(testUser_O1))
	// if err != nil {
	// 	t.Fatal("unexpected error counting monitor documents:", err)
	// }

	// if documentCount != 3 {
	// 	t.Fatalf("Expected count of 3 documents, got %d", documentCount)
	// }

	// Test counting by status
	activeCount, err := store.DocumentCount(DocumentQuery().SetStatus(DOCUMENT_STATUS_ACTIVE))
	if err != nil {
		t.Fatal("unexpected error counting active documents:", err)
	}

	if activeCount != 1 {
		t.Fatalf("Expected count of 1 active document, got %d", activeCount)
	}
}

func TestStore_DocumentCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStore_DocumentCreateDuplicate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error on first create:", err)
	}

	// Try to create the same document again
	err = store.DocumentCreate(document)
	if err == nil {
		t.Fatal("expected error for duplicate document, but got nil")
	}
}

func TestStore_DocumentFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = document.SetMetas(map[string]string{
		"severity":         "high",
		"affected_service": "web",
		"team":             "backend",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.DocumentCreate(document)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	documentFound, errFind := store.DocumentFindByID(document.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if documentFound == nil {
		t.Fatal("Document MUST NOT be nil")
	}

	if documentFound.ID() != document.ID() {
		t.Fatal("IDs do not match")
	}

	if documentFound.Status() != document.Status() {
		t.Fatal("Statuses do not match")
	}

	if documentFound.Status() != DOCUMENT_STATUS_ACTIVE {
		t.Fatal("Statuses do not match")
	}

	severity, err := document.Meta("severity")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	foundSeverity, err := documentFound.Meta("severity")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if foundSeverity != severity {
		t.Fatal("Metas do not match")
	}

	affectedService, err := document.Meta("affected_service")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	foundAffectedService, err := documentFound.Meta("affected_service")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if foundAffectedService != affectedService {
		t.Fatal("Metas do not match")
	}

	team, err := document.Meta("team")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	foundTeam, err := documentFound.Meta("team")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if foundTeam != team {
		t.Fatal("Metas do not match")
	}
}

func TestStore_DocumentFindByIDNotFound(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	documentFound, errFind := store.DocumentFindByID("non-existent-id")

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if documentFound != nil {
		t.Fatal("Document MUST be nil for non-existent ID")
	}
}

// func TestStoreDocumentUpdateNonExistent(t *testing.T) {
// 	store, err := initStore(":memory:")

// 	if err != nil {
// 		t.Fatal("unexpected error:", err)
// 	}

// 	document := NewDocument().
// 		SetStatus(DOCUMENT_STATUS_ACTIVE)

// 	// Try to update a non-existent document
// 	err = store.DocumentUpdate(document)
// 	if err == nil {
// 		t.Fatal("expected error for updating non-existent document, but got nil")
// 	}
// }

func TestStore_DocumentDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Delete the document
	err = store.DocumentDelete(document)
	if err != nil {
		t.Fatal("unexpected error on delete:", err)
	}

	// Verify the document is deleted
	deletedDocument, err := store.DocumentFindByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error finding deleted document:", err)
	}

	if deletedDocument != nil {
		t.Fatal("Document should be nil after deletion")
	}
}

func TestStore_DocumentDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Delete the document by ID
	err = store.DocumentDeleteByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error on delete by ID:", err)
	}

	// Verify the document is deleted
	deletedDocument, err := store.DocumentFindByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error finding deleted document:", err)
	}

	if deletedDocument != nil {
		t.Fatal("Document should be nil after deletion by ID")
	}
}

func TestStore_DocumentList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Create multiple documents
	document1 := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	document2 := NewDocument().
		SetStatus(DOCUMENT_STATUS_INACTIVE).
		SetFileName("test2.txt").
		SetText("This is second test document.")

	err = store.DocumentCreate(document1)
	if err != nil {
		t.Fatal("unexpected error creating document1:", err)
	}

	err = store.DocumentCreate(document2)
	if err != nil {
		t.Fatal("unexpected error creating document2:", err)
	}

	// Test listing all documents
	allDocuments, err := store.DocumentList(DocumentQuery())
	if err != nil {
		t.Fatal("unexpected error listing all documents:", err)
	}

	if len(allDocuments) != 2 {
		t.Fatalf("Expected 2 documents, got %d", len(allDocuments))
	}

	// Test filtering by owner ID
	// ownerDocuments, err := store.DocumentList(DocumentQuery().SetOwnerID(testUser_O1))
	// if err != nil {
	// 	t.Fatal("unexpected error listing owner documents:", err)
	// }

	// if len(ownerDocuments) != 2 {
	// 	t.Fatalf("Expected 2 documents for MONITOR_01, got %d", len(ownerDocuments))
	// }

	// Test filtering by status
	activeDocuments, err := store.DocumentList(DocumentQuery().SetStatus(DOCUMENT_STATUS_ACTIVE))
	if err != nil {
		t.Fatal("unexpected error listing active documents:", err)
	}

	if len(activeDocuments) != 1 {
		t.Fatalf("Expected 1 active document, got %d", len(activeDocuments))
	}

	// Test limit and offset
	limitedDocuments, err := store.DocumentList(DocumentQuery().SetLimit(1))
	if err != nil {
		t.Fatal("unexpected error listing limited documents:", err)
	}

	if len(limitedDocuments) != 1 {
		t.Fatalf("Expected 1 document with limit, got %d", len(limitedDocuments))
	}

	offsetDocuments, err := store.DocumentList(DocumentQuery().SetOffset(1).SetLimit(2))
	if err != nil {
		t.Fatal("unexpected error listing offset documents:", err)
	}

	if len(offsetDocuments) != 1 {
		t.Fatalf("Expected 1 document with offset, got %d", len(offsetDocuments))
	}
}

func TestStore_DocumentSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Soft delete the document
	err = store.DocumentSoftDelete(document)
	if err != nil {
		t.Fatal("unexpected error on soft delete:", err)
	}

	// Verify the document is soft deleted (not found by default)
	softDeletedDocument, err := store.DocumentFindByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error finding soft deleted document:", err)
	}

	if softDeletedDocument != nil {
		t.Fatal("Document should not be found after soft deletion")
	}

	// Verify the document can be found when including soft deleted
	query := DocumentQuery().
		SetWithSoftDeleted(true).
		SetID(document.ID()).
		SetLimit(1)

	documentFindWithDeleted, err := store.DocumentList(query)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(documentFindWithDeleted) == 0 {
		t.Fatal("Document should be found when including soft deleted")
	}

	if documentFindWithDeleted[0].SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Document should be soft deleted, but SoftDeletedAt is MAX_DATETIME:", documentFindWithDeleted[0].SoftDeletedAt())
	}

	if !documentFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Document should be marked as soft deleted")
	}
}

func TestStore_DocumentSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Soft delete the document by ID
	err = store.DocumentSoftDeleteByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error on soft delete by ID:", err)
	}

	// Verify the document is soft deleted (not found by default)
	softDeletedDocument, err := store.DocumentFindByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error finding soft deleted document:", err)
	}

	if softDeletedDocument != nil {
		t.Fatal("Document should not be found after soft deletion by ID")
	}

	// Verify the document can be found when including soft deleted
	query := DocumentQuery().
		SetWithSoftDeleted(true).
		SetID(document.ID()).
		SetLimit(1)

	documentFindWithDeleted, err := store.DocumentList(query)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(documentFindWithDeleted) == 0 {
		t.Fatal("Document should be found when including soft deleted")
	}

	if documentFindWithDeleted[0].SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Document should be soft deleted, but SoftDeletedAt is MAX_DATETIME:", documentFindWithDeleted[0].SoftDeletedAt())
	}

	if !documentFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Document should be marked as soft deleted")
	}
}

func TestStore_DocumentUpdate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	document := NewDocument().
		SetStatus(DOCUMENT_STATUS_ACTIVE).
		SetFileName("test1.txt").
		SetText("This is first test document.")

	err = store.DocumentCreate(document)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// Update the document
	document.SetStatus(DOCUMENT_STATUS_INACTIVE).
		SetMemo("Resolved by ops team")

	err = store.DocumentUpdate(document)
	if err != nil {
		t.Fatal("unexpected error on update:", err)
	}

	// Verify the update
	updatedDocument, err := store.DocumentFindByID(document.ID())
	if err != nil {
		t.Fatal("unexpected error finding updated document:", err)
	}

	if updatedDocument.Status() != DOCUMENT_STATUS_INACTIVE {
		t.Fatalf("Status not updated. Expected %s, got %s", DOCUMENT_STATUS_INACTIVE, updatedDocument.Status())
	}

	if updatedDocument.Memo() != "Resolved by ops team" {
		t.Fatalf("Memo not updated. Expected 'Resolved by ops team', got '%s'", updatedDocument.Memo())
	}
}
