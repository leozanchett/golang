package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (*Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setUp()
	if err != nil {
		return &Consumer{}, err
	}

	return &consumer, nil
}

func (consumer *Consumer) setUp() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}
	for _, s := range topics {
		if err := ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil); err != nil {
			return err
		}
	}
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)
			go handlePayload(payload)
		}
	}()

	fmt.Printf(" [*] Waiting for messages [Exchange: logs_topic, Queue: %s]. To exit press CTRL+C\n", q.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		// log whatever we get
		err := logEvent(payload)
		if err != nil {
			fmt.Println(err)
		}
	case "auth":
		// authenticate the user

	default:
		err := logEvent(payload)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func logEvent(entry Payload) error {
	// create some json we'll send to the log microservice
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	const logServceURL = "http://logger-service:3000/log"

	request, err := http.NewRequest("POST", logServceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
