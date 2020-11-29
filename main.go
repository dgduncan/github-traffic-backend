package main

import (
	"context"
	"fmt"
	"helloworld/internal/handler"
	"helloworld/internal/secret"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi"
)

// type GithubPageViewsResponse struct {
// 	Count   int          `json:"count"`
// 	Uniques int          `json:"uniques"`
// 	Views   []GithubView `json:"views"`
// }

// type GithubView struct {
// 	Timestamp string `json:"timestamp"`
// 	Count     int    `json:"count"`
// 	Uniques   int    `json:"uniques"`
// }

// type GithubClonesResponse struct {
// 	Count   int           `json:"count"`
// 	Uniques int           `json:"uniques"`
// 	Clones  []GithubClone `json:"clones"`
// }

// type GithubClone struct {
// 	Timestamp string `json:"timestamp"`
// 	Count     int    `json:"count"`
// 	Uniques   int    `json:"uniques"`
// }

func init() {
}

func main() {
	ctx := context.Background()

	port := os.Getenv("PORT")
	appID := os.Getenv("GAE_APPLICATION")
	sPath := os.Getenv("GITHUB_TOKEN_SECRET_PATH")

	token, err := secret.FetchSecret(ctx, sPath)
	if err != nil {
		fmt.Println(err)
	}

	gh := &handler.GithubHandler{
		Token: token,
	}

	r := chi.NewRouter()
	r.Get("/test_traffic", gh.TrafficHandle)
	r.Get("/query", func(w http.ResponseWriter, r *http.Request) {
		test := strings.Split(appID, "~")
		dsClient, err := datastore.NewClient(context.Background(), test[1])
		if err != nil {
			fmt.Println(err)
		}

		var entities []*MyEntity
		q := datastore.NewQuery("test")
		_, err = dsClient.GetAll(r.Context(), q, &entities)
		if err != nil {
			fmt.Println(err)
		}

		for _, entity := range entities {
			fmt.Println(entity)
		}

	})
	r.Post("/query_insert", func(w http.ResponseWriter, r *http.Request) {
		test := strings.Split(appID, "~")
		dsClient, err := datastore.NewClient(context.Background(), test[1])
		if err != nil {
			fmt.Println(err)
		}

		entity := new(MyEntity)
		k := datastore.IncompleteKey("test", nil)
		_, err = dsClient.Put(r.Context(), k, entity)
		if err != nil {
			fmt.Println(err)
		}

	})
	http.ListenAndServe(":"+port, r)
}

type MyEntity struct {
	A    int
	K    *datastore.Key `datastore:"__key__"`
	Date time.Time      `datastore:"date"`
}
