package ast

import "github.com/goby-lang/goby/compiler/token"

// Variable interface represents assignable nodes in Goby, currently are Identifier, InstanceVariable and Constant
type Variable interface {
	variableNode()
	ReturnValue() string
	Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) variableNode() {}
func (i *Identifier) ReturnValue() string {
	return i.Value
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type InstanceVariable struct {
	Token token.Token
	Value string
}

func (iv *InstanceVariable) variableNode() {}
func (iv *InstanceVariable) ReturnValue() string {
	return iv.Value
}
func (iv *InstanceVariable) expressionNode() {}
func (iv *InstanceVariable) TokenLiteral() string {
	return iv.Token.Literal
}
func (iv *InstanceVariable) String() string {
	return iv.Value
}

type Constant struct {
	Token       token.Token
	Value       string
	IsNamespace bool
}

func (c *Constant) variableNode() {}
func (c *Constant) ReturnValue() string {
	return c.Value
}
func (c *Constant) expressionNode() {}
func (c *Constant) TokenLiteral() string {
	return c.Token.Literal
}
func (c *Constant) String() string {
	return c.Value
}
