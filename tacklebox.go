package tacklebox

import (
	"github.com/tacklebox-webhooks/golang/components"
)

type Webhooks struct {
	ApiKey  string
	BaseUrl string
}

func (self Webhooks) Service() components.Service {
	service := components.Service{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return service
}

func (self Webhooks) EventType() components.EventType {
	eventType := components.EventType{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return eventType
}

func (self Webhooks) User() components.User {
	user := components.User{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return user
}

func (self Webhooks) Subscription() components.Subscription {
	subscription := components.Subscription{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return subscription
}

func (self Webhooks) Event() components.Event {
	event := components.Event{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return event
}

func (self Webhooks) Message() components.Message {
	message := components.Message{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return message
}
