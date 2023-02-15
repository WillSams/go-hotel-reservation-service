package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	apiPort := os.Getenv("API_PORT")
	fmt.Println("Listening on http://localhost/" + apiPort)
	http.ListenAndServe(":"+apiPort, nil)
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

	playgroundPort := os.Getenv("PLAYGROUND_PORT")
	fmt.Println("Listening on http://localhost:" + playgroundPort)
	http.Handle("/playground", h)
	http.ListenAndServe(":"+playgroundPort, nil)
}
