package Utils

import (
	"hash/crc32"
	"github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers"
	constants "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants"
	ErrorCodes "github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants/ErrorCodes"
)


func SerializeMessage(messageId, payload []byte) *ProtocolBuffers.Msg {
	checksum := ComputeChecksum(messageId, payload)

	return &ProtocolBuffers.Msg{
		MessageID: messageId,
		Payload:   payload,
		CheckSum:  uint64(checksum), //TODO: Check if this is correct if checksum fails
	}
}

func CreateMessageID() []byte {
	return []byte("MessageID Placeholder")   
}

func ComputeChecksum(messageId, payload []byte) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum(messageId, table)
	checksum = crc32.Update(checksum, table, payload)

	return checksum
}

func SerializeKVRequest(command uint32) *ProtocolBuffers.KVRequest {
	return &ProtocolBuffers.KVRequest{
		Command: command,
	}
}

func SerializeKVRequestWithKey(command uint32, key []byte) *ProtocolBuffers.KVRequest {
	return &ProtocolBuffers.KVRequest{
		Command: command,
		Key:     key,
	}
}

func SerializeKVRequestWithKeyValue(command uint32, key, value []byte) *ProtocolBuffers.KVRequest {
	return &ProtocolBuffers.KVRequest{
		Command: command,
		Key:     key,
		Value:   value,
	}
}

func SerializeKVRequestWithKeyValueVersion(command uint32, key, value []byte, version int32) *ProtocolBuffers.KVRequest {
	return &ProtocolBuffers.KVRequest{
		Command: command,
		Key:     key,
		Value:   value,
		Version: &version,
	}
}

func SerializeKVResponse(errCode uint32) *ProtocolBuffers.KVResponse {
	return &ProtocolBuffers.KVResponse{
		ErrCode: errCode,
	}
}

func SerializeOverloadKVResponse(overloadWaitTime int32) *ProtocolBuffers.KVResponse {
	return &ProtocolBuffers.KVResponse{
		ErrCode:          ErrorCodes.TEMP_SYSTEM_OVERLOAD,
		OverloadWaitTime: &overloadWaitTime,
	}
}

func SerializeGetKVResponse(value []byte, version int32) *ProtocolBuffers.KVResponse {
	return &ProtocolBuffers.KVResponse{
		ErrCode: ErrorCodes.OPERATION_SUCCESSFUL,
		Value:   value,
		Version: &version,
	}
}

func SerializePIDKVResponse() *ProtocolBuffers.KVResponse {
	pid := int32(constants.PID)
	return &ProtocolBuffers.KVResponse{
		ErrCode: ErrorCodes.OPERATION_SUCCESSFUL,
		Pid:     &pid,
	}
}

func SerializeMembershipCountKVResponse(membershipCount int32) *ProtocolBuffers.KVResponse {
	return &ProtocolBuffers.KVResponse{
		ErrCode:         ErrorCodes.OPERATION_SUCCESSFUL,
		MembershipCount: &membershipCount,
	}
}
