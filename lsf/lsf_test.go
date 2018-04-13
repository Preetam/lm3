package lsf

import "testing"

func TestChunkIDString(t *testing.T) {
	suffix := chunkID{}.String()
	if len(suffix) != 16+16+1 {
		t.Errorf("Chunk ID %s length is not %d", suffix, 16+16+1)
	}
}
