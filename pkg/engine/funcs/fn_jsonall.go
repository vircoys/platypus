// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"github.com/GuanceCloud/ppl/pkg/ast"
	"github.com/GuanceCloud/ppl/pkg/engine/runtime"
)

func JSONAllChecking(ctx *runtime.Context, funcExpr *ast.CallExpr) error {
	l.Debugf("warning: json_all() is disabled")
	return nil
}

func JSONAll(ctx *runtime.Context, funcExpr *ast.CallExpr) runtime.PlPanic {
	l.Debugf("warning: json_all() is disabled")
	return nil
}
