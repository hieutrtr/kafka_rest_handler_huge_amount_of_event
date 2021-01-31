package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"github.com/kafka_rest_handler_huge_amount_of_event/entity"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (h *Handler) RegisterEventHandler(g *echo.Group) {
	g.POST("", h.registerEvent)
}

func (h *Handler) registerEvent(c echo.Context) error {
	var events []entity.Event
	if err := c.Bind(&events); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "Event data is invalid")
	}
	fmt.Println(events)

	if err := h.produceKafkaEvents(events); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, events)
}

func (h *Handler)produceKafkaEvents(events []entity.Event) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return err
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "hieu_test_event"
	for _, event := range events {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("%v", event)), // serialization
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	return nil
}
