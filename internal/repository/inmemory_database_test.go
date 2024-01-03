package repository

import (
	"testing"
)

func TestSet(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	if db.Get("key1") != "value1" {
		t.Errorf("Set method failed")
	}
}

func TestGet(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	if db.Get("key1") != "value1" {
		t.Errorf("Get method failed")
	}
}

func TestDelete(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.Delete("key1")
	if db.Get("key1") != "" {
		t.Errorf("Delete method failed")
	}
}

func TestCommitTransaction(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	db.Commit()
	if db.Get("key1") != "value2" {
		t.Errorf("Commit method failed")
	}
}

func TestCommitTransactionWithoutTransaction(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.Commit()
	if db.Get("key1") != "value1" {
		t.Errorf("Commit method failed")
	}
}

func TestRollbackTransaction(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.StartTransaction()
	if db.Get("key1") != "value1" {
		t.Errorf("Get method failed before setting new value in transaction")
	}
	db.Set("key1", "value2")
	if db.Get("key1") != "value2" {
		t.Errorf("Get method failed after setting new value in transaction")
	}
	db.Rollback()
	if db.Get("key1") != "value1" {
		t.Errorf("Rollback method failed")
	}
}

func TestRollbackTransactionWithoutTransaction(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.Rollback()
	if db.Get("key1") != "value1" {
		t.Errorf("Rollback method failed")
	}
}

func TestNestedTransactionsCommit(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if db.Get("key1") != "value2" {
		t.Errorf("Get method failed after setting new value in first transaction")
	}
	db.StartTransaction()
	if db.Get("key1") != "value2" {
		t.Errorf("Get method failed in nested transaction before delete")
	}
	db.Delete("key1")
	db.Commit()
	if db.Get("key1") != "" {
		t.Errorf("Commit method failed in nested transaction")
	}
	db.Commit()
	if db.Get("key1") != "" {
		t.Errorf("Commit method failed in outer transaction")
	}
}

func TestNestedTransactionsRollback(t *testing.T) {
	db := NewInMemoryDatabase[string, string]()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if db.Get("key1") != "value2" {
		t.Errorf("Get method failed after setting new value in first transaction")
	}
	db.StartTransaction()
	if db.Get("key1") != "value2" {
		t.Errorf("Get method failed in nested transaction before delete")
	}
	db.Delete("key1")
	db.Rollback()
	if db.Get("key1") != "value2" {
		t.Errorf("Rollback method failed in nested transaction")
	}
	db.Commit()
	if db.Get("key1") != "value2" {
		t.Errorf("Commit method failed in outer transaction after nested rollback")
	}
}
