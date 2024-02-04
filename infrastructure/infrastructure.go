package infrastructure

import "test/infrastructure/handler"

type Infrastructure struct{
	HTTP handler.HTTP
}


func NewInfrastructure() Infrastructure {
	httpHandler := handler.NewHTTP()


	return Infrastructure{
		HTTP: httpHandler,
	}

}