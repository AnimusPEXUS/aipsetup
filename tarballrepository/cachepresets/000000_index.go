package cachepresets

import (
	"errors"

	"github.com/AnimusPEXUS/utils/cache01"
)

func Get(name string) (*cache01.Settings, error) {
	if t, ok := Index[name]; ok {
		return t, nil
	} else {
		return nil, errors.New("cache preset not found")
	}
}

var Index = map[string]*cache01.Settings{
	// nil value should make cache to use default (internal) settings
	"":    nil,
	"nil": nil,
}
