package amqprpc

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go-scaffold/internal/entity"
	"go-scaffold/internal/usecase"
	"go-scaffold/pkg/rabbitmq/rmq_rpc/server"
)

type translationRoutes struct {
	translationUseCase usecase.Translation
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

func (r *translationRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		translations, err := r.translationUseCase.History(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - translationRoutes - getHistory - r.translationUseCase.History: %w", err)
		}
		response := historyResponse{translations}
		return response, nil
	}
}

func newTranslationRoutes(routes map[string]server.CallHandler, t usecase.Translation) {
	r := &translationRoutes{t}
	{
		routes["getHistory"] = r.getHistory()
	}
}
