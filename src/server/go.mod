module github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers => ../ProtocolBuffers

go 1.22.3

require (
	github.com/TayyibChohan/SC_Distributed_Hash_Table/src/ProtocolBuffers v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.0
)

require google.golang.org/protobuf v1.34.1 // indirect
