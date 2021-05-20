package components

import "github.com/tacklebox-webhooks/golang/apis"

type Event struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self Event) List(serviceId, userId string) string {
	api := apis.EventApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListEvents(serviceId, userId)
}

func (self Event) Create(serviceId, userId string, eventData map[string]interface{}) string {
	api := apis.EventApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.CreateEvent(serviceId, userId, eventData)
}

func (self Event) Get(serviceId, userId, eventId string) string {
	api := apis.EventApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetEvent(serviceId, userId, eventId)
}
