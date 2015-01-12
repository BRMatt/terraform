package lang

import (
	"fmt"
	"sync"

	"github.com/hashicorp/terraform/config/lang/ast"
)

// TypeVisitor implements ast.Visitor for type checking an AST tree.
// It requires some configuration to look up the type of nodes.
type TypeVisitor struct {
	Scope *Scope

	stack []ast.Type
	err   error
	lock  sync.Mutex
}

func (v *TypeVisitor) Visit(root ast.Node) error {
	v.lock.Lock()
	defer v.lock.Unlock()
	defer v.reset()
	root.Accept(v.visit)
	return v.err
}

func (v *TypeVisitor) visit(raw ast.Node) {
	if v.err != nil {
		return
	}

	switch n := raw.(type) {
	case *ast.Call:
		v.visitCall(n)
	case *ast.Concat:
		v.visitConcat(n)
	case *ast.LiteralNode:
		v.visitLiteral(n)
	case *ast.VariableAccess:
		v.visitVariableAccess(n)
	default:
		v.createErr(n, fmt.Sprintf("unknown node: %#v", raw))
	}
}

func (v *TypeVisitor) visitCall(n *ast.Call) {
	// Look up the function in the map
	function, ok := v.Scope.LookupFunc(n.Func)
	if !ok {
		v.createErr(n, fmt.Sprintf("unknown function called: %s", n.Func))
		return
	}

	// The arguments are on the stack in reverse order, so pop them off.
	args := make([]ast.Type, len(n.Args))
	for i, _ := range n.Args {
		args[len(n.Args)-1-i] = v.stackPop()
	}

	// Verify the args
	for i, expected := range function.ArgTypes {
		if args[i] != expected {
			v.createErr(n, fmt.Sprintf(
				"%s: argument %d should be %s, got %s",
				n.Func, i+1, expected, args[i]))
			return
		}
	}

	// Return type
	v.stackPush(function.ReturnType)
}

func (v *TypeVisitor) visitConcat(n *ast.Concat) {
	types := make([]ast.Type, len(n.Exprs))
	for i, _ := range n.Exprs {
		types[len(n.Exprs)-1-i] = v.stackPop()
	}

	// All concat args must be strings, so validate that
	for i, t := range types {
		if t != ast.TypeString {
			v.createErr(n, fmt.Sprintf(
				"argument %d must be a sting", n, i+1))
			return
		}
	}

	// This always results in type string
	v.stackPush(ast.TypeString)
}

func (v *TypeVisitor) visitLiteral(n *ast.LiteralNode) {
	v.stackPush(n.Type)
}

func (v *TypeVisitor) visitVariableAccess(n *ast.VariableAccess) {
	// Look up the variable in the map
	variable, ok := v.Scope.LookupVar(n.Name)
	if !ok {
		v.createErr(n, fmt.Sprintf(
			"unknown variable accessed: %s", n.Name))
		return
	}

	// Add the type to the stack
	v.stackPush(variable.Type)
}

func (v *TypeVisitor) createErr(n ast.Node, str string) {
	v.err = fmt.Errorf("%s: %s", n.Pos(), str)
}

func (v *TypeVisitor) reset() {
	v.stack = nil
	v.err = nil
}

func (v *TypeVisitor) stackPush(t ast.Type) {
	v.stack = append(v.stack, t)
}

func (v *TypeVisitor) stackPop() ast.Type {
	var x ast.Type
	x, v.stack = v.stack[len(v.stack)-1], v.stack[:len(v.stack)-1]
	return x
}