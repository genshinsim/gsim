package eval

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/gcs/ast"
)

type letStmtEvalNode struct {
	*ast.LetStmt
	env  *Env
	node evalNode
	res  Obj
}

func (l *letStmtEvalNode) nextAction() (Obj, bool, error) {
	// functionally, FnStmt is just a special type of let statement
	// add ident to env, then create a new fnval
	_, exist := l.env.varMap[l.Ident.Val]
	if exist {
		return nil, false, fmt.Errorf("function %v already exists; cannot redeclare", l.Ident.Val)
	}
	if l.res == nil {
		res, done, err := l.node.nextAction()
		if err != nil {
			return nil, false, err
		}
		if !done {
			// the only time it's not done is if the res is an action
			if res.Typ() == typAction {
				return res, false, nil
			}
			return nil, false, fmt.Errorf("unexpected error; let stmt stopped at non action: %v", l.LetStmt.String())
		}
		l.res = res
	}
	l.env.varMap[l.Ident.Val] = &l.res
	return &null{}, true, nil
}

func letStmtEval(n *ast.LetStmt, env *Env) evalNode {
	return &letStmtEvalNode{
		LetStmt: n,
		env:     env,
		node:    evalFromExpr(n.Val, env),
	}
}
