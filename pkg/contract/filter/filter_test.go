package filter

import "testing"

// ==============================================================
// Defining custom filter
//==============================================================

type UserFilter struct {
	fes []FieldExpression
}

var _ Filter = (*UserFilter)(nil)

func NewUserFilter() Filter {
	return &UserFilter{
		fes: make([]FieldExpression, 0),
	}
}

func (f *UserFilter) AddField(fe ...FieldExpression) Filter {
	f.fes = append(f.fes, fe...)
	return f
}

func (f *UserFilter) And(filter ...Filter) Filter {
	return nil
}

func (f *UserFilter) Or(filter ...Filter) Filter {
	return nil
}

func (f *UserFilter) Not(filter ...Filter) Filter {
	return nil
}

func (f *UserFilter) GetAnd() []Filter {
	return nil
}

func (f *UserFilter) GetOr() []Filter {
	return nil
}

func (f *UserFilter) GetNot() []Filter {
	return nil
}

func (f *UserFilter) GetFields() []FieldExpression {
	return nil
}

//==============================================================

func TestFilterExpression(t *testing.T) {

	filter := NewUserFilter()
	filter.And(
		NewUserFilter().
			Or(
				NewUserFilter().
					AddField(NewFieldExpression("id", NewExpression(Eq, "USR1"))).
					AddField(NewFieldExpression("name", NewExpression(Contains, "John"))),

				NewUserFilter().
					AddField(NewFieldExpression("name", NewExpression(Eq, "John"))),
			),

		NewUserFilter().
			AddField(NewFieldExpression("status", NewExpression(Eq, "active"))),
	)
}
