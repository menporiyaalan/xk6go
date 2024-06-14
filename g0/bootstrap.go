// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"github.com/menporiyaalan/xk6go/internal/builtin/jetstream"
	"github.com/menporiyaalan/xk6go/internal/builtin/nats"
	"github.com/menporiyaalan/xk6go/internal/builtin/gjson"
	"github.com/menporiyaalan/xk6go/internal/builtin/gofakeit"
	"github.com/menporiyaalan/xk6go/internal/builtin/goquery"
	"github.com/menporiyaalan/xk6go/internal/builtin/jsonpath"
	"github.com/menporiyaalan/xk6go/internal/builtin/jsonschema"
	"github.com/menporiyaalan/xk6go/internal/builtin/logrus"
	"github.com/menporiyaalan/xk6go/internal/builtin/resty"
	"github.com/menporiyaalan/xk6go/internal/builtin/stdlib"
	"github.com/menporiyaalan/xk6go/internal/builtin/testify"
	"go.k6.io/k6/js/modules"
)

func registerBuiltins() {
	RegisterExports(jetstream.Exports)
	RegisterExports(nats.Exports)
	RegisterExports(stdlib.Exports)
	RegisterExports(logrus.Exports)
	RegisterExports(resty.Exports)
	RegisterExports(testify.Exports)
	RegisterExports(goquery.Exports)
	RegisterExports(gjson.Exports)
	RegisterExports(jsonpath.Exports)
	RegisterExports(jsonschema.Exports)
	RegisterExports(gofakeit.Exports)
}

func registerExtension() {
	modules.Register("k6/x/g0", New())
}

func Bootstrap() {
	redirectStdin()
	registerBuiltins()
	registerExtension()
}
