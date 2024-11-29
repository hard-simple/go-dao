package dao

import (
	"context"
	"fmt"
	"github.com/hard-simple/go-dao/pkg/contract/config"
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
	err := Register(dbName, userDao)
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

	err, createResponse := drefDao.Create(ctx, &CreateRequest[User]{
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
			&ReadRequest[Filter]{
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
	DAO[string, User, Filter]
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

func (u *InMemoryUserDAO) Create(ctx context.Context, request *CreateRequest[User]) (error, *CreateResponse[User]) {
	data := request.Data
	data.id = u.getAndIncrementId()
	u.users[data.id] = data
	return nil, &CreateResponse[User]{
		Data: &data,
	}
}

func (u *InMemoryUserDAO) BulkCreate(ctx context.Context, request *BulkCreateRequest[User]) (error, *BulkCreateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Read(ctx context.Context, request *ReadRequest[Filter]) (error, *ReadResponse[User]) {
	return nil, &ReadResponse[User]{
		Data: []User{u.users[request.Filter.id]},
	}
}

func (u *InMemoryUserDAO) ReadBulk(ctx context.Context, request *BulkReadRequest[Filter]) (error, *BulkReadResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) RangeRead(ctx context.Context, request *RangeReadRequest[Filter]) (error, *RangeReadResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Update(ctx context.Context, request *UpdateRequest[User]) (error, *UpdateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) BulkUpdate(ctx context.Context, request *BulkUpdateRequest[User]) (error, *BulkUpdateResponse[User]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) Delete(ctx context.Context, request *DeleteRequest[string]) (error, *DeleteResponse[string]) {
	panic("implement me")
}

func (u *InMemoryUserDAO) BulkDelete(ctx context.Context, request *BulkDeleteRequest[string]) (error, *BulkDeleteResponse[string]) {
	panic("implement me")
}
