package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoshikiShibata/gpl/ch07/ex16/eval"
)

func main() {
	fmt.Printf("For + sign, use '%%2b'\n")
	fmt.Printf("ex) http://localhost:8000/calc?expr=(1.1 * 2 %%2b 3)*3&expr=355/113")
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values, ok := r.Form["expr"]
	if !ok {
		http.Error(w, "no expr", http.StatusBadRequest)
		return
	}

	for _, v := range values {
		expr, err := eval.Parse(v)
		if err != nil {
			http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
			return
		}

		result := expr.Eval(eval.Env{})
		fmt.Fprintf(w, "%s = %g\n", v, result)
	}
}
