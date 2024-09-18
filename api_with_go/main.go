package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items      []Item
	itemsMutex sync.Mutex
	idCounter  int
)

func main() {
	http.HandleFunc("/items", listItems)
	http.HandleFunc("/items/add", addItem)
	http.HandleFunc("/items/update", updateItem)
	http.HandleFunc("/items/delete", deleteItem)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func listItems(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idCounter++
	newItem.ID = idCounter
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found := false
	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	found := false
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
}
