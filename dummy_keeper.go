package main  

import "errors"

type DummyKeeper struct {
	mem map[string]string
}

func (k DummyKeeper) Get(key string) (string, error) {
	value, ok := k.mem[key]
	if !ok {
		return "", errors.New(NotFoundError)
	}
	return value, nil
}

func (k DummyKeeper) Set(key string, message string) error {
	k.mem[key] = message
	return nil
}

func (k DummyKeeper) Clean(key string) error {
	delete(k.mem, key)
	return nil
}
