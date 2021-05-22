package components

import "github.com/tacklebox-webhooks/golang/apis"

type EventType struct {
	ApiKey  string
	BaseUrl string
}

func (self EventType) List(serviceId string) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.ListEventTypes(serviceId)
}

func (self EventType) Create(serviceId string, eventTypeData map[string]interface{}) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.CreateEventType(serviceId, eventTypeData)
}

func (self EventType) Get(serviceId, eventTypeId string) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.GetEventType(serviceId, eventTypeId)
}

func (self EventType) Delete(serviceId, eventTypeId string) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.DeleteEventType(serviceId, eventTypeId)
}
