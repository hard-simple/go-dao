package dao

import "fmt"

type Filter[F any] struct {
	And    *Filter[F]
	Or     *Filter[F]
	Not    *Filter[F]
	Fields F
}

type UserFilter struct {
	And    *UserFilter
	Or     *UserFilter
	Not    *UserFilter
	Fields []FieldExpression
	ID     *StringFilterExpression
	Name   *StringFilterExpression
}

//===========================================================================

type FieldExpression interface {
	Name() string
	Expression() Expression
}

type SimpleFieldExpression struct {
	FieldExpression
	name       string
	expression Expression
}

func (s *SimpleFieldExpression) Name() string {
	return s.name
}

func (s *SimpleFieldExpression) Expression() Expression {
	return s.expression
}

func NewFieldExpression(name string, expression Expression) FieldExpression {
	return &SimpleFieldExpression{
		name:       name,
		expression: expression,
	}
}

//===========================================================================

type Operation int

var Eq Operation = 0
var Contains Operation = 1

type Expression interface {
	Op() Operation
	Value() any
}

type SimpleExpression struct {
	op    Operation
	value any
}

func (s *SimpleExpression) Op() Operation {
	return s.op
}

func (s *SimpleExpression) Value() any {
	return s.value
}

func NewSimpleExpression(op Operation, value any) Expression {
	return &SimpleExpression{
		op:    op,
		value: value,
	}
}

//===========================================================================

type FilterV2 interface {
	Expressions() []Expression
}

type BaseFilterV2 struct {
	expressions []Expression
}

func (b *BaseFilterV2) Expressions() []Expression {
	return b.expressions
}

type StringExpression[C comparable] interface {
	FilterBuilder[C]
	Eq(val string) StringExpression[C]
	Contains(val string) StringExpression[C]
	Build() FilterV2
}

type BaseStringExpression struct {
}

type IntExpression[C comparable] interface {
	FilterBuilder[C]
	Range(from int, to int) IntExpression[C]
	Eq(val int) IntExpression[C]
	Gt(val int) IntExpression[C]
	Gte(val int) IntExpression[C]
	Build() FilterV2
}

type FilterBuilder[C comparable] interface {
	And() FilterBuilder[C]
	Or() FilterBuilder[C]
	Not() FilterBuilder[C]
	ByInt(column C) IntExpression[C]
	ByString(column C) StringExpression[C]
	Build() FilterV2
}

func NewFilterBuilder[C comparable]() FilterBuilder[C] {
	return nil
}

type Column interface {
	ID() ColumnID
	Type() ColumnType
}

type ColumnType int

type ColumnID int

var id ColumnID = 0
var name ColumnID = 1

//===========================================================================

type StringFilterExpression struct {
	Expression
	Eq       *string
	Contains *string
}

func NewStringFieldExpression(eq *string, contains *string) Expression {
	return &StringFilterExpression{
		Eq:       eq,
		Contains: contains,
	}
}

func example() {

	filter1 := &UserFilter{
		And: &UserFilter{
			ID: &StringFilterExpression{
				Eq: ptr("asd"),
			},
			Name: &StringFilterExpression{
				Contains: ptr("bla"),
			},
		},
	}

	filter2 := &UserFilter{
		And: &UserFilter{
			Fields: []FieldExpression{
				NewFieldExpression("name", NewSimpleExpression(Contains, ptr("blabla"))),
				NewFieldExpression("id", NewSimpleExpression(Eq, ptr("123"))),
			},
		},
	}

	filter3 :=
		NewFilterBuilder[ColumnID]().
			And().
			ByInt(id).
			Gt(10).
			And().
			ByString(name).
			Contains("dummy").
			Build()

	fmt.Printf("%v : %v : %v", filter1, filter2, filter3)

}

func ptr[T any](t T) *T {
	return &t
}
