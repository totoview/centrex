// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "pb/CentrexMsg.pb.go", nil, 0)
	if err != nil {
		panic(err)
	}
	// ast.Print(fset, f)

	reqTypes := []string{}
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			for _, spec := range x.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if !strings.HasPrefix(ts.Name.Name, "CentrexMsg") && !strings.HasSuffix(ts.Name.Name, "Rsp") {
						reqTypes = append(reqTypes, ts.Name.Name)
					}
				}
			}
		}
		return true
	})

	httpCodec, err := os.Create("http/codec.go")
	if err != nil {
		panic(err)
	}
	defer httpCodec.Close()

	fmt.Fprintf(httpCodec, "// Code generated DO NOT EDIT.\n\n")
	fmt.Fprintf(httpCodec, "package http\n\n")

	fmt.Fprintf(httpCodec, "import(\n")
	fmt.Fprintf(httpCodec, "\t\"context\"\n")
	fmt.Fprintf(httpCodec, "\t\"encoding/json\"\n")
	fmt.Fprintf(httpCodec, "\t\"net/http\"\n")
	fmt.Fprintf(httpCodec, "\t\"github.com/totoview/centrex/pb\"\n")
	fmt.Fprintf(httpCodec, ")\n\n")

	fmt.Fprintf(httpCodec,
		`func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeHttpError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}`+"\n")

	for _, reqType := range reqTypes {
		fmt.Fprintf(httpCodec, "\nfunc Decode%sRequest(_ context.Context, r *http.Request) (interface{}, error) {\n", reqType)
		fmt.Fprintf(httpCodec, "\tvar req pb.%s\n", reqType)
		fmt.Fprintf(httpCodec, "\tif err := json.NewDecoder(r.Body).Decode(&req); err != nil {\n")
		fmt.Fprintf(httpCodec, "\t\treturn nil, err\n")
		fmt.Fprintf(httpCodec, "\t}\n")
		fmt.Fprintf(httpCodec, "\treturn req, nil\n")
		fmt.Fprintf(httpCodec, "}\n")
	}
}
