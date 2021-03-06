package handlers

import (
	"context"
	"fmt"

	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/internal/env"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/internal/mongodb"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/internal/rabbitmq"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/pkg/model"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/pkg/usecase"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/shoppinglist"
	"go.opentelemetry.io/otel"
)

type RpcServer struct {
	mongoClient                       *mongodb.MongoClient
	getCustomerShoppingListUseCase    *usecase.GetCustomerShoppingListUseCase
	removeCustomerShoppingListUseCase *usecase.RemoveCustomerShoppingListUseCase
	changeCustomerShoppingListUseCase *usecase.ChangeCustomerShoppingListUseCase
	rabbitmq                          rabbitmq.RabbitMqClient
}

func InitRpc(ctx context.Context) (*RpcServer, error) {
	shutdown := initPrivder(ctx)
	defer shutdown()
	client, err := mongodb.NewTracedClient(context.TODO(), env.GetEnvOrDefault("MONGODB_CONNECTION", "mongodb://db:27017"), otel.GetTracerProvider())
	if err != nil {
		return nil, fmt.Errorf("error when trying connect to mongo, ERR: %w", err)
	}
	repo := mongodb.NewMongoShoppingListsRepository(client)
	rabbitmqclient, err := rabbitmq.NewRabbitMqClient(env.GetEnvOrDefault("RABBITMQ_CONNECTION", "amqp://guest:guest@rabbitmq:5672/"))
	if err != nil {
		return nil, fmt.Errorf("error when trying connect to rabbitmq, ERR: %w", err)
	}
	publisher := rabbitmq.NewRabbitmMqCustomerBasketChangedEventPublisher(rabbitmqclient, &rabbitmq.RabbitMqConfig{ExchangeName: "basket", Topic: "changed"}, &rabbitmq.RabbitMqConfig{ExchangeName: "basket", Topic: "removed"})
	getCustomerShoppingListUseCase := usecase.NewGetCustomerShoppingListUseCase(repo)
	changeCustomerShoppingListUseCase := usecase.NewChangeCustomerShoppingListUseCase(repo, publisher)
	removeCustomerShoppingListUseCase := usecase.NewRemoveCustomerShoppingListUseCase(repo, publisher)
	err = initData(repo)
	if err != nil {
		return nil, fmt.Errorf("error when trying save initial data, ERR: %w", err)
	}
	return &RpcServer{mongoClient: client, removeCustomerShoppingListUseCase: removeCustomerShoppingListUseCase, getCustomerShoppingListUseCase: getCustomerShoppingListUseCase, changeCustomerShoppingListUseCase: changeCustomerShoppingListUseCase, rabbitmq: rabbitmqclient}, nil
}

func (serv *RpcServer) GetCustomerShoppingList(ctx context.Context, req *shoppinglist.GetCustomerShoppingListRequest) (*shoppinglist.GetCustomerShoppingListResponse, error) {
	res, err := serv.getCustomerShoppingListUseCase.Execute(ctx, int(req.CustomerId))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &shoppinglist.GetCustomerShoppingListResponse{Empty: true}, nil
	}
	items := make([]*shoppinglist.CustomerShoppingList_Item, len(res.Items))
	for i, item := range res.Items {
		items[i] = &shoppinglist.CustomerShoppingList_Item{ItemId: int32(item.ItemID), ItemQuantity: int32(item.ItemQuantity)}
	}
	return &shoppinglist.GetCustomerShoppingListResponse{Empty: false, CustomerShoppingList: &shoppinglist.CustomerShoppingList{CustomerId: req.CustomerId, Items: items}}, nil
}

func (s *RpcServer) ChangeCustomerShoppingList(ctx context.Context, basket *shoppinglist.ChangeCustomerShoppingListRequest) (*shoppinglist.ChangeCustomerShoppingListResponse, error) {
	items := make([]model.Item, len(basket.ShoppingList.Items))
	for i, item := range basket.ShoppingList.Items {
		items[i] = *model.NewItem(int(item.ItemId), int(item.ItemQuantity))
	}
	model := model.NewCustomerShoppingList(int(basket.CustomerId), items)
	err := s.changeCustomerShoppingListUseCase.Execute(ctx, model)
	if err != nil {
		return nil, err
	}
	return &shoppinglist.ChangeCustomerShoppingListResponse{Success: true}, nil
}

func (s *RpcServer) RemoveCustomerShoppingList(ctx context.Context, req *shoppinglist.RemoveCustomerShoppingListRequest) (*shoppinglist.RemoveCustomerShoppingListResponse, error) {
	err := s.removeCustomerShoppingListUseCase.Execute(ctx, int(req.CustomerId))
	if err != nil {
		return nil, err
	}

	return &shoppinglist.RemoveCustomerShoppingListResponse{Success: true}, nil
}

func (ser *RpcServer) Close(ctx context.Context) {
	ser.mongoClient.Close(ctx)
	ser.rabbitmq.Close()
}
