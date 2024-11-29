package dao

import (
	"context"
	"github.com/hard-simple/go-dao/pkg/contract/config"
	"io"
)

// DAO is an abstract Data Access Object interface for accessing data objects. The interface represents CRUD
// operations for working with single entity/object.
//
// The interface has two generics:
//
// - K represents a key of a target object. Key should be unique in scope of all entities/objects.
//
// - T represents a target entity/object.
//
// - F represents a filter type. It provides flexibility for implementer.
//
// The interface doesn't support transactions since it could be implemented on the context layer. Feel free to use
// Tx interface for this purpose.
type DAO[K any, T any, F any] interface {

	// Configurable uses for set up a proper configuration for the implementation before use it.
	// Configure function should be firstly called after definition instantiation.
	config.Configurable

	// Closer - a good example of an IO resource usage is closing the resource after all operations are completed.
	io.Closer

	// Create creates an entity, otherwise returns an error.
	Create(ctx context.Context, request *CreateRequest[T]) (error, *CreateResponse[T])

	// BulkCreate creates a number of entities, otherwise returns an error.
	BulkCreate(ctx context.Context, request *BulkCreateRequest[T]) (error, *BulkCreateResponse[T])

	// Read returns a list of data according to the request.
	Read(ctx context.Context, request *ReadRequest[F]) (error, *ReadResponse[T])

	// BulkRead returns a map of data where key is unique value of the entity and value is the found value.
	// If there is no entity for unique value then map shouldn't contain that key.
	BulkRead(ctx context.Context, request *BulkReadRequest[F]) (error, *BulkReadResponse[T])

	// RangeRead returns a list of values from the range, otherwise returns an error.
	RangeRead(ctx context.Context, request *RangeReadRequest[F]) (error, *RangeReadResponse[T])

	// Update updates properties of entity, otherwise returns an error.
	Update(ctx context.Context, request *UpdateRequest[T]) (error, *UpdateResponse[T])

	// BulkUpdate updates a list of entities, otherwise returns an error. It also supports partial update design.
	BulkUpdate(ctx context.Context, request *BulkUpdateRequest[T]) (error, *BulkUpdateResponse[T])

	// Delete removes entity, otherwise returns an error.
	Delete(ctx context.Context, request *DeleteRequest[K]) (error, *DeleteResponse[K])

	// BulkDelete removes a list of entities, otherwise returns an error. The design support partial remove.
	BulkDelete(ctx context.Context, request *BulkDeleteRequest[K]) (error, *BulkDeleteResponse[K])
}

// Metadata represents a set of additional meta info that could be useful for downstream side.
type Metadata map[string]interface{}

type CreateRequest[T any] struct {

	// Data for creation.
	Data T

	// If it's true then the entity will be updated in case it already exists.
	// If it's false then the error will be returned in case it already exists.
	// It's optional. By default, it's not supposed to be used.
	Upsert *bool
}

type CreateResponse[T any] struct {

	// Data after create operation. It's optional.
	Data *T

	// Whether the entity was updated or not.
	Updated *bool

	// Whether the entity was created or not.
	Created *bool

	// Create operation Metadata
	Metadata Metadata
}

type BulkCreateRequest[T any] struct {

	// Data for creation.
	Data []T

	// If it's true then the entity will be updated in case it already exists.
	// If it's false then the error will be returned in case it already exists.
	// It's optional. By default, it's not supported.
	Upsert *bool

	// Whether the operation could be completed partially or not.
	// It's optional. By default, it's not supposed to be used.
	Partial *bool
}

type BulkCreateResponse[T any] struct {

	// Data after create operation. It's optional.
	// If it was partial request and Data isn't nil then the size of Data is a number of
	// successfully created/updated entities.
	Data *[]T

	// Bulk create operation Metadata
	Metadata Metadata
}

type UpdateRequest[T any] struct {

	// Source data
	Data T

	// Upsert mode.
	// If it's true then the entity will be updated in case it already exists.
	// If it's false then the error will be returned in case it already exists.
	// It's optional. By default, it's not supposed to be used.
	Upsert *bool
}

type UpdateResponse[T any] struct {

	// Updated data. If it is nil then an implementation doesn't support updated entity in the response.
	Data *T

	// Whether the entity was created or not.
	Created *bool

	// Whether the entity was updated or not.
	Updated *bool
}

type BulkUpdateRequest[T any] struct {

	// Data for update.
	Data []T

	// Upsert mode.
	// If it's true then the entity will be updated in case it already exists.
	// If it's false then the error will be returned in case it already exists.
	// It's optional. By default, it's not supposed to be used.
	Upsert *bool

	// Whether the operation could be completed partially or not.
	// It's optional. By default, it's not supposed to be used.
	Partial *bool
}

type BulkUpdateResponse[T any] struct {

	// Data after update operation. It's optional.
	// If it was partial request and Data isn't nil then the size of Data is a number of
	// successfully created/updated entities.
	Data *[]T

	// Bulk update operation Metadata
	Metadata Metadata
}

type ReadRequest[F any] struct {

	// Filtration info. It's optional.
	Filter *F

	// Page by page process info.
	Pagination *Pagination
}

type ReadResponse[T any] struct {

	// Found data according to the request filter.
	Data []T

	// Page by page process info.
	Pagination *Pagination
}

type BulkReadRequest[F any] struct {

	// Filtration info. It's optional.
	Filter *F

	// Page by page process info.
	Pagination *Pagination
}

type BulkReadResponse[T any] struct {

	// Found data according to the request filter.
	Data []T

	// Page by page process info.
	Pagination *Pagination
}

type RangeReadRequest[F any] struct {

	// Filtration info. It's optional.
	Filer *F

	// Page by page process info.
	Pagination *Pagination
}

type RangeReadResponse[T any] struct {

	// Found data according to the request filter.
	Data []T

	// Page by page process info.
	Pagination *Pagination
}

type DeleteRequest[K any] struct {

	// The key of entity for deletion.
	Key K
}

type DeleteResponse[K any] struct {

	// The key of deleted entity.
	Key K
}

type BulkDeleteRequest[K any] struct {

	// Keys of entities to be deleted.
	Keys []K

	// Whether the operation could be completed partially or not.
	// It's optional. By default, it's not supposed to be used.
	Partial *bool
}

type BulkDeleteResponse[K any] struct {

	// A set of keys were successfully deleted. If it was partial operation, and it isn't nil then the size of the
	// set will be equal to the number of deleted entities. It should be equal to the request key size
	// if request wasn't partial or default.
	Deleted *[]K
}
