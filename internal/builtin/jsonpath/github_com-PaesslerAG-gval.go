// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

// Code generated by 'yaegi extract github.com/PaesslerAG/gval'. DO NOT EDIT.

package jsonpath

import (
	"github.com/PaesslerAG/gval"
	"reflect"
)

func init() {
	Symbols["github.com/PaesslerAG/gval/gval"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Arithmetic":          reflect.ValueOf(gval.Arithmetic),
		"Base":                reflect.ValueOf(gval.Base),
		"Bitmask":             reflect.ValueOf(gval.Bitmask),
		"Constant":            reflect.ValueOf(gval.Constant),
		"Evaluate":            reflect.ValueOf(gval.Evaluate),
		"Full":                reflect.ValueOf(gval.Full),
		"Function":            reflect.ValueOf(gval.Function),
		"InfixBoolOperator":   reflect.ValueOf(gval.InfixBoolOperator),
		"InfixEvalOperator":   reflect.ValueOf(gval.InfixEvalOperator),
		"InfixNumberOperator": reflect.ValueOf(gval.InfixNumberOperator),
		"InfixOperator":       reflect.ValueOf(gval.InfixOperator),
		"InfixShortCircuit":   reflect.ValueOf(gval.InfixShortCircuit),
		"InfixTextOperator":   reflect.ValueOf(gval.InfixTextOperator),
		"JSON":                reflect.ValueOf(gval.JSON),
		"NewLanguage":         reflect.ValueOf(gval.NewLanguage),
		"PostfixOperator":     reflect.ValueOf(gval.PostfixOperator),
		"Precedence":          reflect.ValueOf(gval.Precedence),
		"PrefixExtension":     reflect.ValueOf(gval.PrefixExtension),
		"PrefixMetaPrefix":    reflect.ValueOf(gval.PrefixMetaPrefix),
		"PrefixOperator":      reflect.ValueOf(gval.PrefixOperator),
		"PropositionalLogic":  reflect.ValueOf(gval.PropositionalLogic),
		"Text":                reflect.ValueOf(gval.Text),
		"VariableSelector":    reflect.ValueOf(gval.VariableSelector),

		// type definitions
		"Evaluable":  reflect.ValueOf((*gval.Evaluable)(nil)),
		"Evaluables": reflect.ValueOf((*gval.Evaluables)(nil)),
		"Language":   reflect.ValueOf((*gval.Language)(nil)),
		"Parser":     reflect.ValueOf((*gval.Parser)(nil)),
	}
}
