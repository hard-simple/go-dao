package config

import (
	"context"
	"go-dao/pkg/contract/registry"
)

var (
	r = registry.New[Producer](context.Background())
)

// Register function produces pluggable solution for Config Producer implementations.
func Register(name string, instance Producer) error {
	return r.Register(name, instance)
}

// GetConfigProducer returns Producer by name from the registry. It returns error if there is no an instance for the name.
// Also, it can return error if the result instance has unexpected type.
func GetConfigProducer(name string) (error, *Producer) {
	return r.Get(name)
}
