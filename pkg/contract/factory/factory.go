package factory

import (
	"context"
	"errors"
	"go-dao/pkg/contract/config"
	"go-dao/pkg/contract/dao"
	"sync"
)

type DAOFactory[K any, T any, F any] interface {
	Make(ctx context.Context, name string) (error, *dao.DAO[K, T, F])
}

type singletonDAOFactory[K any, T any, F any] struct {
	DAOFactory[K, T, F]
	instance *dao.DAO[K, T, F]
	mutex    sync.Mutex
}

func NewSingletonDAOFactory[K any, T any, F any](ctx context.Context) DAOFactory[K, T, F] {
	return &singletonDAOFactory[K, T, F]{}
}

func (s *singletonDAOFactory[K, T, F]) Make(ctx context.Context, name string) (error, *dao.DAO[K, T, F]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.instance != nil {
		return nil, s.instance
	}

	err, producer := config.GetConfigProducer(name)
	if err != nil {
		return err, nil
	}

	err, targetDao := dao.GetDAO[dao.DAO[K, T, F]](name)
	if err != nil {
		return err, nil
	}

	if targetDao == nil {
		return errors.New(name + " dao wasn't found"), nil
	}

	err = (*targetDao).Configure(ctx, (*producer)(ctx))
	if err != nil {
		return err, nil
	}

	s.instance = targetDao

	return nil, s.instance
}
