// Package lsf implements a log-structured file.
package lsf

import "fmt"

// File is a log-structured file.
type File struct {
	// Name of the file
	name string
	// Current version
	version int64
	// Max chunk size in bytes
	maxChunkSize int64
	// Chunks
	chunks []*chunk
	// Current writable chunk
	writable    *chunk
	dirtyChunks map[chunkID]*chunk
}

// chunkID uniquely represents a chunk
type chunkID struct {
	offset  int64
	version int64
}

// chunk is a versioned portion of a file.
type chunk struct {
	offset  int64
	version int64
	data    []byte
}

func (id chunkID) String() string {
	return fmt.Sprintf("%016x-%016x", id.offset, id.version)
}
