package filter

//=================================================================================================

type simpleFieldExpression struct {
	FieldExpression
	name       string
	expression Expression
}

func (s *simpleFieldExpression) Name() string {
	return s.name
}

func (s *simpleFieldExpression) Expression() Expression {
	return s.expression
}

// NewFieldExpression creates a simple field expression. It is a container with getters.
func NewFieldExpression(name string, expression Expression) FieldExpression {
	return &simpleFieldExpression{
		name:       name,
		expression: expression,
	}
}

//===========================================================================

type simpleExpression struct {
	op    Operation
	value any
}

func (s *simpleExpression) Op() Operation {
	return s.op
}

func (s *simpleExpression) Value() any {
	return s.value
}

// NewExpression creates a simple expression. It is a container with getters.
func NewExpression(op Operation, value any) Expression {
	return &simpleExpression{
		op:    op,
		value: value,
	}
}

//===========================================================================
