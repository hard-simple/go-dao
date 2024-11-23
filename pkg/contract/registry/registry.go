package registry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Registry[T any] struct {
	instances map[string]any
	mu        sync.Mutex
}

func New[T any](ctx context.Context) *Registry[T] {
	return &Registry[T]{
		instances: map[string]any{},
	}
}

// Register function produces pluggable solution.
func (r *Registry[T]) Register(name string, instance T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.instances[name]; ok {
		return errors.New(name + " instance is already registered")
	}

	r.instances[name] = instance
	return nil
}

// Get returns an instance by name from the registry. It returns error if there is no an instance for the name.
// Also, it can return error if the result instance has unexpected type.
func (r *Registry[T]) Get(name string) (error, *T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if instance, ok := r.instances[name]; ok {

		if targetInst, casted := instance.(T); casted {
			return nil, &targetInst
		}
		return fmt.Errorf(
			"%s instance has unexpected type %s, expected type %s",
			name,
			reflect.TypeOf(instance),
			reflect.TypeOf((*T)(nil)),
		), nil
	}
	return fmt.Errorf("%s instance hasn't found in registry", name), nil
}
