package tests

import (
	"github.com/hard-simple/go-dao/pkg/contract/filter"
	"testing"
)

// ==============================================================
// Defining custom filter
//==============================================================

type UserFilter struct {
	fes []filter.FieldExpression
}

var _ filter.Filter = (*UserFilter)(nil)

func NewUserFilter() filter.Filter {
	return &UserFilter{
		fes: make([]filter.FieldExpression, 0),
	}
}

func (f *UserFilter) AddField(fe ...filter.FieldExpression) filter.Filter {
	f.fes = append(f.fes, fe...)
	return f
}

func (f *UserFilter) And(filter ...filter.Filter) filter.Filter {
	return nil
}

func (f *UserFilter) Or(filter ...filter.Filter) filter.Filter {
	return nil
}

func (f *UserFilter) Not(filter ...filter.Filter) filter.Filter {
	return nil
}

func (f *UserFilter) GetAnd() []filter.Filter {
	return nil
}

func (f *UserFilter) GetOr() []filter.Filter {
	return nil
}

func (f *UserFilter) GetNot() []filter.Filter {
	return nil
}

func (f *UserFilter) GetFields() []filter.FieldExpression {
	return nil
}

//==============================================================

func TestFilterExpression(t *testing.T) {

	sourceFilter := NewUserFilter()
	sourceFilter.And(
		NewUserFilter().
			Or(
				NewUserFilter().
					AddField(filter.NewFieldExpression("id", filter.NewExpression(filter.Eq, "USR1"))).
					AddField(filter.NewFieldExpression("name", filter.NewExpression(filter.Contains, "John"))),

				NewUserFilter().
					AddField(filter.NewFieldExpression("name", filter.NewExpression(filter.Eq, "John"))),
			),

		NewUserFilter().
			AddField(filter.NewFieldExpression("status", filter.NewExpression(filter.Eq, "active"))),
	)

	if sourceFilter == nil {
		t.Fail()
	}
}
