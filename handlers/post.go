package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"schedulearn/docker"
)

type Post struct {
	Logger *log.Logger
}

type Message struct {
	Message string `json:"message"`
}

func NewPost(logger *log.Logger) *Post {
	return &Post{logger}
}

func (post *Post) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200 OK
	post.Logger.Printf("Received 'POST' request")

	// open json file
	jsonFile, err := os.Open("message.json")
	if err != nil {
		post.Logger.Println(err)
		return
	}
	fmt.Println("Successfully Opened message.json")

	// defer the closing of our jsonFile
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var message Message

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'message' which we defined above.
	json.Unmarshal(byteValue, &message)

	// we then print out the content of Message which is
	// a slice of type Message.
	post.Logger.Println(message.Message)

	err = docker.Run(message.Message)
	if err != nil {
		post.Logger.Println("Error running docker:", err)
		return
	}
}