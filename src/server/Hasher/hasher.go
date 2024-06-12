package Hasher

import (
	"crypto/sha1"
	"encoding/hex"
)

// Hasher is a struct that contains a SHA1 hash
type Hasher struct {
	Hash []byte
}

// NewHasher creates a new Hasher with the given key
func NewHasher(key string) *Hasher {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return &Hasher{Hash: hasher.Sum(nil)}
}

// GetHash returns the hash as a string
func (h *Hasher) GetHash() string {
	return hex.EncodeToString(h.Hash)
}