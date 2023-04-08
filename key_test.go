package main

import "testing"


func TestDummyKeyBuilder(t *testing.T) {
	dummy_key_builder := DummyKeyBuilder{}
	if dummy_key_builder.Get() != DummyTestKey {
		t.Error("bad dummy key")
	}
}