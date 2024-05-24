module github.com/TayyibChohan/SC_Distributed_Hash_Table/src/client

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Utils => ../server/Utils

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants => ../server/Constants

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/shared/ProtocolBuffers => ../shared/ProtocolBuffers


go 1.22.3

require github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants v0.0.0-00010101000000-000000000000
