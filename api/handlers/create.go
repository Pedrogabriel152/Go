package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Pedrogabriel152/Go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"message": fmt.Sprintf("Ocorreu um erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"message": fmt.Sprintf("todo certo: ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
