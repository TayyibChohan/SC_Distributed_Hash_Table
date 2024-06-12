package KVStore

import (
	errCodes "github.com/omkz/golang-assignment/src/server/Constants/ErrorCodes"
)

type KVStore struct {
	// Fields
}

// NewKVStore creates a new KVStore
func NewKVStore() *KVStore {
	// Your code here
	return nil
}

// Get retrieves the value for the given key
func (kv *KVStore) Get(key string) string {
	// Your code here
	return ""
}

// Put inserts the given key-value pair
func (kv *KVStore) Put(key byte[], value byte[], version int32) errCodes.ErrorCode {
	// Your code here
	return errCodes.OPERATION_SUCCESSFUL
}

// Delete removes the key-value pair for the given key
func (kv *KVStore) Delete(key string) {
	// Your code here
}
