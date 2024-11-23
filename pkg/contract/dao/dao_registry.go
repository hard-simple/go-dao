package dao

import (
	"context"
	"fmt"
	"github.com/hard-simple/go-dao/pkg/contract/registry"
	"reflect"
)

var (
	r = registry.New[any](context.Background())
)

// Register function produces pluggable solution for DAO implementations.
func Register[D any](name string, instance D) error {
	a := toAny(instance)
	return r.Register(name, a)
}

// GetDAO returns DAO by name from the registry. It returns error if there is no an instance for the name.
// Also, it can return error if the result instance has unexpected type.
func GetDAO[T any](name string) (error, *T) {
	err, a := r.Get(name)
	if err != nil {
		return err, nil
	}

	dr := *a
	if cast, ok := dr.(T); ok {
		return nil, &cast
	}
	return fmt.Errorf(
		"dao instance has unexpected type %s, expected type %s",
		reflect.TypeOf(dr),
		reflect.TypeOf((*T)(nil)),
	), nil
}

func toAny[T any](value T) any {
	return value
}
