package filter

import (
	"fmt"
	"sync"
)

// Filter is an interface designing a possibility to filter data by a set of complex expressions. Basically, it is
// a container with well-structured set of filtration expressions.
type Filter interface {

	// And adds a list of filters that will be joined by AND operator.
	// It should compose all Not filters if it is called several times. The operation call guaranty ordering of
	// all And filters in the current Filter instance.
	// It returns the current filter instance.
	//
	And(f ...Filter) Filter

	// Or adds a list of filters that will be joined by OR operator.
	// It should compose all Or filters if it is called several times. The operation call guaranty ordering of
	// all Or filters in the current Filter instance.
	// It returns the current filter instance.
	//
	Or(f ...Filter) Filter

	// Not adds a list of filters that will be joined by NOT operator.
	// It should compose all Not filters if it is called several times. The operation call guaranty ordering of
	// all Not filters in the current Filter instance.
	// It returns the current filter instance.
	//
	Not(f ...Filter) Filter

	// AddField adds a list of fields into the filter.
	// It should compose all field expressions if it is called several times. The operation call guaranty ordering of
	// all fields in the current Filter instance.
	// It returns the current filter instance.
	//
	AddField(fe ...FieldExpression) Filter

	// GetAnd returns a list of And filters that were added before this call. If there are no And filters
	// then empty array should be returned.
	//
	GetAnd() []Filter

	// GetOr returns a list of Or filters that were added before this call. If there are no Or filters
	// then empty array should be returned.
	//
	GetOr() []Filter

	// GetNot returns a list of Not filters that were added before this call. If there are no Not filters
	// then empty array should be returned.
	//
	GetNot() []Filter

	// GetFields returns a list of field expressions which were added before this call. If there are no
	// fields then empty array should be returned.
	//
	GetFields() []FieldExpression
}

// FieldExpression represents an expression on some field.
type FieldExpression interface {

	// Name of the field used by expression.
	Name() string

	// Expression of the field.
	Expression() Expression
}

// Expression describes operation with a value.
type Expression interface {

	// Op is Operation of the expression.
	Op() Operation

	// Value is a value of the expression. It is optional.
	Value() any
}

// Operation is an operation that will be applied on a field expression.
type Operation int

// Undefined is undefined.
var Undefined Operation = 0

// Eq is equals.
var Eq Operation = 1

// Contains is contains.
var Contains Operation = 2

// StartWith is checking out whether the value is started with some prefix. Usually, it is used on string based fields.
var StartWith Operation = 3

// Gt is greater than.
var Gt Operation = 4

// Gte is greater than equals.
var Gte Operation = 5

// Lt is less than.
var Lt Operation = 6

// Lte is less than equals.
var Lte Operation = 7

// RegEx is regular expression.
var RegEx Operation = 8

// Between is a range expression. Usually, it uses together with a range of values.
var Between Operation = 9

var ops = []Operation{Undefined, Eq, Contains, StartWith, Gt, Gte, Lt, Lte, RegEx, Between}
var opsMap = map[int]Operation{}
var mu sync.Mutex
var opsInz = false

func initOps() {
	mu.Lock()
	defer mu.Unlock()
	if !opsInz {
		for _, op := range ops {
			opsMap[FromOperation(op)] = op
		}
		opsInz = true
	}
}

// ToOperation converts `int` data type representation of the operation into Operation instance. If there is no such
// Operation for the input `int` value then error will be returned.
func ToOperation(op int) (error, Operation) {
	if !opsInz {
		initOps()
	}
	if op, ok := opsMap[op]; ok {
		return nil, op
	}
	return fmt.Errorf("there is no operation for code %d", op), Undefined
}

// FromOperation converts Operation into appropriate `int` representation.
func FromOperation(op Operation) int {
	return int(op)
}
