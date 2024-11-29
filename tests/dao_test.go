package tests

import (
	"context"
	"fmt"
	"github.com/hard-simple/go-dao/pkg/contract/config"
	"github.com/hard-simple/go-dao/pkg/contract/dao"
	"github.com/hard-simple/go-dao/pkg/contract/factory"
	"os"
	"strconv"
	"sync"
	"testing"
)

func TestRunInMemoryUserDAO(t *testing.T) {

	// DAO info

	dbName := "in-memory"

	// Init instances

	userDao := NewInMemoryUserDAO()
	err := dao.Register(dbName, userDao)
	if err != nil {
		panic(err)
	}

	err = config.Register(dbName, func(ctx context.Context) config.Config {
		maxBatchSize := os.Getenv("MAX_BATCH_SIZE")
		if maxBatchSize == "" {
			maxBatchSize = "10"
		}
		confVal, castErr := strconv.Atoi(maxBatchSize)
		if castErr != nil {
			fmt.Println("error occurred while trying to fetch configuration")
			panic(castErr)
		}
		return &InMemoryConfig{
			MaxBatchSize: confVal,
		}
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	daoFactory := factory.NewSingletonDAOFactory[string, User, Filter](ctx)
	err, sourceDAO := daoFactory.Make(ctx, dbName)
	if err != nil {
		panic(err)
	}

	drefDao := *sourceDAO

	defer func() {
		e := drefDao.Close()
		if e != nil {
			panic(e)
		}
	}()

	userName := "Yev"

	err, createResponse := drefDao.Create(ctx, &dao.CreateRequest[User]{
		Data: User{
			name: userName,
		},
	})

	if err != nil {
		panic(err)
	}

	userId := createResponse.Data.id

	err, readResponse :=
		drefDao.Read(
			ctx,
			&dao.ReadRequest[Filter]{
				Filter: &Filter{
					id: userId,
				},
			},
		)

	if err != nil {
		panic(err)
	}

	if readResponse.Data[0].name != userName {
		t.Fail()
	}

}

type User struct {
	id      string
	name    string
	hobbies []string
	friends []string
}

type Filter struct {
	id   string
	name string
}

type InMemoryConfig struct {
	config.Config
	MaxBatchSize int
}

type UserDAO interface {
	dao.DAO[string, User, Filter]
}

type InMemoryUserDAO struct {
	users   map[string]User
	counter int
	mu      sync.Mutex
}

var _ UserDAO = (*InMemoryUserDAO)(nil)

func NewInMemoryUserDAO() UserDAO {
	return &InMemoryUserDAO{
		users:   map[string]User{},
		counter: 0,
	}
}

func (u *InMemoryUserDAO) getAndIncrementId() string {
	u.mu.Lock()
	defer u.mu.Unlock()
	id := u.counter
	u.counter = id + 1
	return strconv.Itoa(id)
}

func (u *InMemoryUserDAO) Configure(ctx context.Context, config config.Config) error {
	fmt.Printf("User DAO is configured by %v", config)
	return nil
}

func (u *InMemoryUserDAO) Close() error {
	fmt.Printf("User DAO is closed")
	return nil
}

func (u *InMemoryUserDAO) Create(ctx context.Context, request *dao.CreateRequest[User]) (error, *dao.CreateResponse[User]) {
	data := request.Data
	data.id = u.getAndIncrementId()
	u.users[data.id] = data
	return nil, &dao.CreateResponse[User]{
		Data: &data,
	}
}

func (u *InMemoryUserDAO) BulkCreate(ctx context.Context, request *dao.BulkCreateRequest[User]) (error, *dao.BulkCreateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Read(ctx context.Context, request *dao.ReadRequest[Filter]) (error, *dao.ReadResponse[User]) {
	return nil, &dao.ReadResponse[User]{
		Data: []User{u.users[request.Filter.id]},
	}
}

func (u *InMemoryUserDAO) BulkRead(ctx context.Context, request *dao.BulkReadRequest[Filter]) (error, *dao.BulkReadResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) RangeRead(ctx context.Context, request *dao.RangeReadRequest[Filter]) (error, *dao.RangeReadResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Update(ctx context.Context, request *dao.UpdateRequest[User]) (error, *dao.UpdateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) BulkUpdate(ctx context.Context, request *dao.BulkUpdateRequest[User]) (error, *dao.BulkUpdateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Delete(ctx context.Context, request *dao.DeleteRequest[string]) (error, *dao.DeleteResponse[string]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) BulkDelete(ctx context.Context, request *dao.BulkDeleteRequest[string]) (error, *dao.BulkDeleteResponse[string]) {
	panic("implement me")
}
