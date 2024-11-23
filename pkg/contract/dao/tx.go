package dao

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

	TxConfig map[string]interface{}

	TxProducer[T any] func(config *TxConfig) (error, *Tx[T])

	TxKey string
)

var txKey = TxKey("db-tx")

func NewTx[T any](ctx context.Context, config *TxConfig, producer *TxProducer[T]) (error, context.Context) {
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