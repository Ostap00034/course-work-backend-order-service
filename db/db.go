package db

import (
	"context"
	"log"

	"github.com/Ostap00034/course-work-backend-order-service/ent"
	_ "github.com/lib/pq"
)

func NewClient(dsn string) *ent.Client {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Выполняем автоматическую миграцию схемы
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
