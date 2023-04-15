package main

import "testing"

func TestDummyKeyBuilder(t *testing.T) {
	dummyKeyBuilder := DummyKeyBuilder{}
	key, _ := dummyKeyBuilder.Get()
	if key != DummyTestKey {
		t.Error("bad dummy key")
	}
}
