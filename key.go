package main

const DummyTestKey = "test_key"

type KeyBuilder interface {
	Get() string
}

type DummyKeyBuilder struct {
}

func (k DummyKeyBuilder) Get() string {
	return DummyTestKey
}

func getKeyBuilder() KeyBuilder {
	return DummyKeyBuilder{}
}

var keyBuilder = getKeyBuilder()
