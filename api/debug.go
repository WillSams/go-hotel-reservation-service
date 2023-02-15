package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func DebubGraphQlApiHandler() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		db := dbConnect()
		defer db.Close()

		queryParameters := make(map[string]interface{})
		for key, value := range r.URL.Query() {
			queryParameters[key] = value[0]
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  r.URL.RawQuery,
			VariableValues: queryParameters,
		})

		response, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	fmt.Println("Listening on http://localhost/api:8080")
	http.ListenAndServe(":8080", nil)
}

func PlaygroundHandler() {
	db := dbConnect()
	defer db.Close()

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	fmt.Println("Listening on http://localhost/playground:8080")
	http.Handle("/playground", h)
	http.ListenAndServe(":8080", nil)
}
