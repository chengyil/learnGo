package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"regexp"
	"strconv"
	"strings"
	"webtutorial/dto"
)

var (
	todos = struct {
		storage []*dto.Todo
		index   int64
	}{
		storage: []*dto.Todo{
			{
				ID:     1,
				Author: "CY",
				Title:  "Buy Milk",
			},
			{
				ID:     2,
				Author: "Wendy",
				Title:  "Buy Milk",
			},
		},
		index: 3,
	}
	todoPattern = regexp.MustCompile(`/todo/(?P<ID>\d+)`)
)

func main() {
	startServer()
}

func startServer() {
	registerEndpoint()
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err == nil {
		fmt.Printf("Error while serving %v", err)
	}
}

type worldHandler struct{}

func (t *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func registerEndpoint() {
	http.Handle("/world", new(worldHandler))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("header", r.Header)
		fmt.Println("header Get", r.Header.Get("Header1"))
		fmt.Println("header Access", r.Header["Header1"])
		fmt.Println("content-type", r.Header["content-type"], "content-type")
		fmt.Println("content-type parsed", r.Header[textproto.CanonicalMIMEHeaderKey("content-type")], textproto.CanonicalMIMEHeaderKey("content-type"))
		message := []byte("world")
		w.Write(message)
	})
	http.Handle("/todo", new(todoController))
	http.HandleFunc("/todo/", handleTodo)
}

func handleTodo(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi((todoPattern.ReplaceAllString(r.URL.Path, "${ID}")))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if id >= len(todos.storage) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	todo := todos.storage[id]

	if todo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch {
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	case strings.EqualFold(r.Method, "GET"):
		getTodo(todo, w, r)
		return
	case strings.EqualFold(r.Method, "PUT"):
		updateTodo(todo, w, r)
		return
	}
}

func updateTodo(todo *dto.Todo, w http.ResponseWriter, r *http.Request) {
	payload := new(dto.Todo)

	err := json.Unmarshal(readAll(r.Body), payload)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	todo.Author = payload.Author
	todo.Done = payload.Done
	todo.Title = payload.Title

	w.WriteHeader(http.StatusNoContent)
}

func readAll(r io.Reader) []byte {
	data := make([]byte, 0, 1)

	for {
		if len(data) == cap(data) {
			data = append(data, 0)[:len(data)]
		}

		n, err := r.Read(data[len(data):cap(data)])
		data = data[:n+len(data)]
		if err != nil {
			break
		}
	}

	return data
}

func getTodo(todo *dto.Todo, w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

type todoController struct {
}

func (t *todoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	default:
		w.Write([]byte("UNKNOWN METHOD"))
	case strings.EqualFold(r.Method, "GET"):
		t.GetTodos(w, r)
	case strings.EqualFold(r.Method, "POST"):
		t.NewTodos(w, r)
	}
}

func (t *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(todos.storage)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to get todo"))
		return
	}

	w.Write(payload)
}

func (t *todoController) NewTodos(w http.ResponseWriter, r *http.Request) {

	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read payload"))
		return
	}

	var request dto.NewTodoRequest
	if err = json.Unmarshal(payload, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid payload %s", string(payload))))
		return
	}

	if request.Author == "" || request.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid payload %s", string(payload))))
		return
	}

	todo := &dto.Todo{
		ID:     todos.index,
		Author: request.Author,
		Title:  request.Title,
	}

	todos.storage = append(todos.storage, todo)
	todos.index += 1

	w.WriteHeader(204)
}
