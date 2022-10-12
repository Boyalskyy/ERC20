package main

import (
	db2 "ERC20/db"
	"ERC20/events"
	"encoding/json"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	database, err := db2.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	go events.GetEvent(database)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", Recovery(GetFromDB(database)))
	http.ListenAndServe(":8080", r)
}
func Recovery(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
func GetFromDB(db *sqlx.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		more := r.URL.Query().Get("more")
		less := r.URL.Query().Get("less")
		event := r.URL.Query().Get("event")
		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "10"
		}
		cursor := r.URL.Query().Get("cursor")
		if cursor == "" {
			cursor = "0"
		}
		order := r.URL.Query().Get("order")
		if order == "" {
			order = "asc"
		}
		eventsQuery := db2.NewEvents(db)
		eventsQuery = eventsQuery.Page(limit, cursor, order)
		if from != "" {
			eventsQuery = eventsQuery.FilterByAddresFrom(from)
		}
		if to != "" {
			eventsQuery = eventsQuery.FilterByAddresTo(to)
		}
		if event != "" {
			eventsQuery = eventsQuery.FilterByEvents(event)
		}
		if more != "" {
			eventsQuery = eventsQuery.FilterByAmountMore(more)
		}
		if less != "" {
			eventsQuery = eventsQuery.FilterByAmountLess(less)
		}
		result, err := eventsQuery.Select()
		var resultDeteils []EventsDeteils

		for _, s := range result {
			x := EventsDeteils{
				ID:          s.ID,
				LogName:     s.LogName,
				AddressFrom: s.AddressFrom,
				AddressTo:   s.AddressTo,
				Amount:      s.Amount,
			}
			resultDeteils = append(resultDeteils, x)

		}
		if err != nil {
			log.Println(errors.Wrap(err, "Select error"))
			w.WriteHeader(500)
			return
		}

		err = json.NewEncoder(w).Encode(&resultDeteils)
		if err != nil {
			log.Println(errors.Wrap(err, "failed to render response"))
			w.WriteHeader(500)
		}

	})

}

type EventsDeteils struct {
	ID          int64  `json:"id"`
	LogName     string `json:"log_name"`
	AddressFrom string `json:"address_from"`
	AddressTo   string `json:"address_to"`
	Amount      string `json:"amount"`
}
