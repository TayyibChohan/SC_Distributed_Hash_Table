package constants

// "time"
import (
	"os"
)

var PID = os.Getpid()

const (
	// Server constants
	MAX_KEY_SIZE          = 32       // 32 bytes
	MAX_VALUE_SIZE        = 10000    // 10KB
	MAXIMUM_TOTAL_KV_SIZE = 64000000 // 64MB
	MAX_MESSAGE_SIZE      = 16000    // 16KB

	LOCALHOST          = "127.0.0.1" // Localhost IP
	SHUTDOWN_EXIT_CODE = 1           // Exit code for shutdown

	// Cache constants
	EXPIRATION_TIME             = 2000    // 2 seconds
	CACHE_CLEANUP_RETRY_TIMEOUT = 50      // 1 second
	MAXIMUM_CACHE_CAPACITY      = 1000000 // 10 MB

	// // Constants for Replication
	// REPLICATION_FACTOR    = 3
	// MAX_KEYS_PER_REQUEST  = 50

)
