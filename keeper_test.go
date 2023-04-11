package main

import "testing"

func TestDummyKeeperSet(t *testing.T) {
	keeper := DummyKeeper{mem: make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.Set(key, value)
	if keeper.mem[key] != value {
		t.Error("bad memory value")
	}
}

func TestDummyKeeperGet(t *testing.T) {
	keeper := DummyKeeper{mem: make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.mem[key] = value
	value_from_get, _ := keeper.Get(key)
	if value_from_get != value {
		t.Error("bad value from get")
	}
}

func TestDummyKeeperClean(t *testing.T) {
	keeper := DummyKeeper{mem: make(map[string]string)}
	key := "foo"
	value := "bar"
	keeper.mem[key] = value
	keeper.Clean(key)
	_, ok := keeper.mem[key]
	if ok {
		t.Error("clean does not work")
	}
}
