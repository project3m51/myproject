package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// LINE webhook event structure (simplified)
type Event struct {
	Events []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Message    struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Println("Received event:", event)

	// Respond to each event
	for _, e := range event.Events {
		replyMessage(e.ReplyToken, "Hello from Go webhook!")
	}

	w.WriteHeader(http.StatusOK)
}

// Function to send reply back to LINE
func replyMessage(replyToken string, message string) {
	url := "https://api.line.me/v2/bot/message/reply"
	reqBody := map[string]interface{}{
		"replyToken": replyToken,
		"messages": []map[string]string{
			{
				"type": "text",
				"text": message,
			},
		},
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, ioutil.NopCloser(
		bytes.NewReader(jsonBody)),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearerpackage main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// LINE webhook event structure (simplified)
type Event struct {
	Events []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Message    struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Println("Received event:", event)

	// Respond to each event
	for _, e := range event.Events {
		replyMessage(e.ReplyToken, "Hello from Go webhook!")
	}

	w.WriteHeader(http.StatusOK)
}

// Function to send reply back to LINE
func replyMessage(replyToken string, message string) {
	url := "https://api.line.me/v2/bot/message/reply"
	reqBody := map[string]interface{}{
		"replyToken": replyToken,
		"messages": []map[string]string{
			{
				"type": "text",
				"text": message,
			},
		},
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, ioutil.NopCloser(
		bytes.NewReader(jsonBody)),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearerpackage main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// LINE webhook event structure (simplified)
type Event struct {
	Events []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Message    struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Println("Received event:", event)

	// Respond to each event
	for _, e := range event.Events {
		replyMessage(e.ReplyToken, "Hello from Go webhook!")
	}

	w.WriteHeader(http.StatusOK)
}

// Function to send reply back to LINE
func replyMessage(replyToken string, message string) {
	url := "https://api.line.me/v2/bot/message/reply"
	reqBody := map[string]interface{}{
		"replyToken": replyToken,
		"messages": []map[string]string{
			{
				"type": "text",
				"text": message,
			},
		},
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, ioutil.NopCloser(
		bytes.NewReader(jsonBody)),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "BearerMCvgvIB3PWzxi6WjIK+hdQgG/6pCExJyUZlFhEZjzT2LKqLALkGx2YFKoHPBu6NVwNCiiCS12h/BXLOJbdv6/s2uX4nWbzLuqz4DFfZ5L1q9zsAME+cPallPjmfGCSQJqybkJ7cat4qbGJKRkqE5qwdB04t89/1O/w1cDnyilFU=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending reply:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("LINE reply status:", resp.Status)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending reply:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("LINE reply status:", resp.Status)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending reply:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("LINE reply status:", resp.Status)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
