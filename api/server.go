package api

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingItems []Items
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		shoppingItems: []item{},
	} 
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/shoppding-items", s.listShoppdingItem().Methods("GET"))
	s.HandleFunc("/shoppding-items", s.createShoppingItem().Methods("POST"))
	s.HandleFunc("/shoppding-items/{id}", s.removeShoppingItem().Methods("DELETE"))
}

// clousure pattern
func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return 
		}
		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)


		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppdingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]  
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItems {
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[1+1:]...)
				break
			}
		}
	}
}
