package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerEmployee struct {
	st map[string]string
}

// Mapper
// DTO - Data Transfer Object

// *int => nil, inteiro
// *struct => nil, struct

type Employee struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ResponseGetByIdEmployee struct {
	Message string    `json:"message"`
	Data    *Employee `json:"data"`
	Error   bool      `json:"error"`
}

// /api/employees/:id

// Get returns the employee with the given id
func (c *HandlerEmployee) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// path params
		id := chi.URLParam(r, "id")
		// query params
		category := r.URL.Query().Get("category")
		fmt.Println("category", category)

		// process
		// -> get employee
		employee, ok := c.st[id]
		if !ok {
			code := http.StatusNotFound
			body := ResponseGetByIdEmployee{Message: "employee not found", Data: nil, Error: true}

			w.WriteHeader(code)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&body)
			return
		}

		code := http.StatusOK
		body := &ResponseGetByIdEmployee{
			Message: "Employee found",
			Data: &Employee{
				Id:   id,
				Name: employee,
			},
			Error: false,
		}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}
}
