package main

import (
	"broker/event"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayLoad  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayLoad struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "log":
		app.logEventViaRabbit(w, requestPayload.Log)
	case "mail":
		app.sendMail(w, requestPayload.Mail)
	default:
		app.errorJSON(w, errors.New("invalid action"))
	}
}

func (app *Config) logItem(w http.ResponseWriter, entry LogPayLoad) {
	// create some json we'll send to the log microservice
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	const logServceURL = "http://logger-service:3000/log"

	request, err := http.NewRequest("POST", logServceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("log service returned an error"))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Logged"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) authenticate(w http.ResponseWriter, auth AuthPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(auth, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authenticator-service:3000/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("authentication service error"))
		return
	}

	// create a variable we'll read response.Body into
	var responsePayload jsonResponse
	// decode
	err = json.NewDecoder(response.Body).Decode(&responsePayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if responsePayload.Error {
		app.errorJSON(w, errors.New(responsePayload.Message))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = responsePayload.Data

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) sendMail(w http.ResponseWriter, msg MailPayload) {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	// call the mail service
	const MAIL_SERVICE_URL = "http://mailer-service:3000/send"
	request, err := http.NewRequest("POST", MAIL_SERVICE_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("mail service returned an error"))
		return
	}

	// send back json
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Mail sent to " + msg.To
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) logEventViaRabbit(w http.ResponseWriter, l LogPayLoad) {
	err := app.pushToQueue(l.Name, l.Data)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Logged via RabbitMQ"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}

	payload := LogPayLoad{
		Name: name,
		Data: msg,
	}

	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}
