package components

import "juanpe.example.com/golang/apis"

type Message struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self Message) List(serviceId, userId string) string {
	api := apis.MessageApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListMessages(serviceId, userId)
}

func (self Message) Resend(serviceId, userId, messageId string) string {
	api := apis.MessageApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ResendMessage(serviceId, userId, messageId)
}

func (self Message) Get(serviceId, userId, messageId string) string {
	api := apis.MessageApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetMessage(serviceId, userId, messageId)
}
