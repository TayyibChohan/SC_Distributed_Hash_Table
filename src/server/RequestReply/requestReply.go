package server

import (
	"errors"
	"os"

	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	commands "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants/Commands"
	errorCodes "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants/ErrorCodes"

	. "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Hasher"
	. "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/KVStore"
)

type RequestReply struct {
	PrimaryKVStore   *KVStore
	SecondaryKVStore *KVStore
	Hasher           *Hasher
}

func NewRequestReply(hasher *Hasher, primaryKVStore *KVStore, secondaryKVStore *KVStore) *RequestReply {
	return &RequestReply{
		PrimaryKVStore:   primaryKVStore,
		SecondaryKVStore: secondaryKVStore,
		Hasher:           hasher,
	}
}

func (rr *RequestReply) HandleRequest(kvRequest *ProtocolBuffers.KVRequest) (*ProtocolBuffers.KVResponse, error) {
	// Implement request handling logic here
	// This is a simplified example
	switch kvRequest.Command {

	case commands.PUT:
		return rr.put(kvRequest.Key, kvRequest.Value, *kvRequest.Version, true), nil

	case commands.TRANSFER_KEYS:
		return rr.put(kvRequest.Key, kvRequest.Value, *kvRequest.Version, true), nil

	case commands.TRANSFER_REPLICA_KEYS:
		return rr.put(kvRequest.Key, kvRequest.Value, *kvRequest.Version, false), nil
	case commands.SHUTDOWN:
		os.Exit(constants.SHUTDOWN_EXIT_CODE)
		return nil, errors.New("shutdown - should not reach here")

	case commands.GET:
		return rr.get(kvRequest.Key), nil

	case commands.REMOVE:
		return rr.remove(kvRequest.Key, true), nil

	case commands.REMOVE_REPLICA_KEYS:
		return rr.remove(kvRequest.Key, false), nil

	case commands.WIPEOUT:
		return rr.wipeout(), nil

	// Add other cases here
	case commands.IS_ALIVE:
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}, nil

	case commands.GET_PID:
		// Implement the logic for GET_PID
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL, Pid: &constants.PID}, nil

	case commands.GET_MEMBERSHIP_COUNT:
		// Implement the logic for GET_MEMBERSHIP_COUNT
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	case commands.GET_MEMBERSHIP_LIST:
		// Implement the logic for GET_MEMBERSHIP_LIST
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	case commands.SHARE_MEMBERSHIP:
		// Implement the logic for SHARE_MEMBERSHIP
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	case commands.WIPEOUT_REPLICAS:
		// Implement the logic for WIPEOUT_REPLICAS
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	case commands.REBUILD_SECONDARY:
		// Implement the logic for REBUILD_SECONDARY
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	case commands.REBUILD_PRIMARY:
		// Implement the logic for REBUILD_PRIMARY
		// Placeholder for actual implementation
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.UNIMPLEMENTED}, nil

	default:
		return nil, errors.New("invalid command")
	}

}
func (rr *RequestReply) put(key, value []byte, version int32, isPrimary bool) *ProtocolBuffers.KVResponse {
	//TODO: Update to handle replication
	if isPrimary {
		success := rr.PrimaryKVStore.Put(key, value, version)
		if !success {
			return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OUT_OF_SPACE}
		}
	} else {
		success := rr.SecondaryKVStore.Put(key, value, version)
		if !success {
			return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OUT_OF_SPACE}
		}
	}
	return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
}

func (rr *RequestReply) get(key []byte) *ProtocolBuffers.KVResponse {
	return rr.PrimaryKVStore.Get(string(key))
}

// func (rr *RequestReply) transferKeys() *ProtocolBuffers.KVResponse {
// 	keys := rr.PrimaryKVStore.GetKeys()
// 	for _, key := range keys {
// 		val := rr.PrimaryKVStore.Get(key)
// 		rr.SecondaryKVStore.Put([]byte(key), val.Value, *val.Version)
// 	}
// 	return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
// }

// TODO: Handle failure case
func (rr *RequestReply) remove(key []byte, isPrimary bool) *ProtocolBuffers.KVResponse {
	if isPrimary {
		if !rr.PrimaryKVStore.ContainsKey(string(key)) {
			return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.NONEXISTENT_KEY}
		}
		
		rr.PrimaryKVStore.Remove(string(key))
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
	} else {
		if !rr.SecondaryKVStore.ContainsKey(string(key)) {
			return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.NONEXISTENT_KEY}
		}

		rr.SecondaryKVStore.Remove(string(key))
		return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
	}

}

func (rr *RequestReply) wipeout() *ProtocolBuffers.KVResponse {
	rr.PrimaryKVStore.Wipeout()
	return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
}
