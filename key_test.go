package main

import "testing"

func TestDummyKeyBuilder(t *testing.T) {
	dummyKeyBuilder := DummyKeyBuilder{}
	if dummyKeyBuilder.Get() != DummyTestKey {
		t.Error("bad dummy key")
	}
}
