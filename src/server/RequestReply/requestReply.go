package server

import (
	"errors"
	"os"

	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	commands "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants/Commands"
	errorCodes "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants/ErrorCodes"

	. "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/Hasher"
	. "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/KVStore"
)

// CommandCodes, ErrorCodes, and CommonResponses should be defined according to your application's requirements.

type RequestReply struct {
	PrimaryKVStore   KVStore
	SecondaryKVStore KVStore
	Hasher           Hasher
}

func NewRequestReply(hasher Hasher, primaryKVStore, secondaryKVStore KVStore) *RequestReply {
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
	// Add other cases here
	default:
		return nil, errors.New("invalid command")
	}
}

func (rr *RequestReply) put(key, value []byte, version int32, isPrimary bool) *ProtocolBuffers.KVResponse {
	// Implement put operation here
	// This is a simplified example
	errCode := rr.PrimaryKVStore.Put(key, value, version)
	if errCode != errorCodes.OPERATION_SUCCESSFUL {
		return &ProtocolBuffers.KVResponse{ErrCode: errCode}
	}
	return &ProtocolBuffers.KVResponse{ErrCode: errorCodes.OPERATION_SUCCESSFUL}
}

// Implement other methods (get, remove, wipeout, etc.) similarly

func (rr *RequestReply) shutdown() {
	os.Exit(constants.SHUTDOWN_EXIT_CODE)
}
