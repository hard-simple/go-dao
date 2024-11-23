package tx

import (
	"context"
	"fmt"
	"reflect"
)

type (
	Tx[T any] interface {
		ID() T
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}

	Config map[string]interface{}

	Producer[T any] func(config *Config) (error, *Tx[T])

	Key string
)

var txKey = Key("db-tx")

func NewTx[T any](ctx context.Context, config *Config, producer *Producer[T]) (error, context.Context) {
	err, tx := (*producer)(config)
	if err != nil {
		return err, nil
	}
	return nil, context.WithValue(ctx, txKey, tx)
}

func GetTx[T any](ctx context.Context) (error, *Tx[T]) {
	rawTx := ctx.Value(txKey)
	if rawTx == nil {
		return nil, nil
	}
	if casted, ok := rawTx.(*Tx[T]); ok {
		return nil, casted
	}
	return fmt.Errorf("incorrect transcation type %s", reflect.TypeOf(rawTx)), nil
}
