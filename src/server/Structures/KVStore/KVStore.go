package KVStore

import (
	"sync"
	pb "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
)

type KVStore struct {
	// Fields
	kvMap              map[string]*pb.KVResponse
	totalItemSize      int
	maxKVStoreCapacity int
	mu                 sync.Mutex
}

// NewKVStore creates a new KVStore
func NewKVStore() *KVStore {
	return &KVStore{
		kvMap:              make(map[string]*pb.KVResponse, constants.MAXIMUM_TOTAL_KV_SIZE),
		totalItemSize:      0,
		maxKVStoreCapacity: constants.MAXIMUM_TOTAL_KV_SIZE*constants.MAX_VALUE_SIZE,
	}
}

// Put inserts the given key-value pair
func (kv *KVStore) Put(key []byte, value []byte, version int32) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	val := &pb.KVResponse{
		Value:   value,
		Version: &version,
	}

	keyStr := string(key) // Convert key to string

	itemSizeDiff := len(value) + len(keyStr)
	if v, ok := kv.kvMap[keyStr]; ok {
		itemSizeDiff -= len(v.Value) + len(keyStr)
	}

	if kv.totalItemSize+itemSizeDiff > kv.maxKVStoreCapacity {
		return false // Assuming you have an error code for capacity exceeded
	}

	kv.kvMap[keyStr] = val
	kv.totalItemSize += itemSizeDiff

	return true // Assuming you have an error code for success
}

func (kv *KVStore) ContainsKey(key string) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	keyStr := string(key)
	_, ok := kv.kvMap[keyStr]
	return ok
}

func (kv *KVStore) GetKeys() []string {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	keys := make([]string, 0, len(kv.kvMap))
	for k := range kv.kvMap {
		keys = append(keys, k)
	}
	return keys
}

func (kv *KVStore) Get(key string) *pb.KVResponse {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if v, ok := kv.kvMap[key]; ok {
		return v
	}
	return nil // replace with your non-existent key response
}

func (kv *KVStore) Remove(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if v, ok := kv.kvMap[key]; ok {
		kv.totalItemSize -= len(v.Value) + len(key) + 4 // 4 is the size of the version integer
		delete(kv.kvMap, key)
	}
}

func (kv *KVStore) Wipeout() {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.kvMap = make(map[string]*pb.KVResponse, len(kv.kvMap))
	kv.totalItemSize = 0
}

