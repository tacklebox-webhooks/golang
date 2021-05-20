package components

import "github.com/tacklebox-webhooks/golang/apis"

type EventType struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self EventType) List(serviceId string) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListEventTypes(serviceId)
}

func (self EventType) Create(serviceId string, eventTypeData map[string]interface{}) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.CreateEventType(serviceId, eventTypeData)
}

func (self EventType) Get(serviceId, eventTypeId string) string {
	api := apis.EventTypeApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetEventType(serviceId, eventTypeId)
}

// func (self EventType) delete(serviceId, eventTypeId string) string {
// 	api := apis.EventTypeApi{apiKey: self.ApiKey, baseUrl: self.BaseUrl, stage: self.Stage}
// 	return api.deleteEventType(serviceId, eventTypeId)
// }
