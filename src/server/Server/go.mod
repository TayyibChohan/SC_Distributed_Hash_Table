module github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Server

go 1.22.3

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Utils => ../Utils

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes => ../Structures/nodes

replace github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants => ../Constants

require (
	github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Constants v0.0.0-00010101000000-000000000000
	github.com/TayyibChohan/SC_Distributed_Hash_Table/src/server/Structures/nodes v0.0.0-00010101000000-000000000000
)
