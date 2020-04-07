package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GetGoing/section-4/todo-api/model"
	"github.com/GetGoing/section-4/todo-api/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// take data & save it to DB
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Delete error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item deleted"})
		}
	}
}