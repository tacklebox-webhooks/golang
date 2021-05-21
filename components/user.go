package components

import "github.com/tacklebox-webhooks/golang/apis"

type User struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self User) List(serviceId string) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListUsers(serviceId)
}

func (self User) Create(serviceId string, userData map[string]interface{}) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.CreateUser(serviceId, userData)
}

func (self User) Get(serviceId, userId string) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetUser(serviceId, userId)
}

func (self User) Delete(serviceId, userId string) string {
	api := apis.UserApi{apiKey: self.ApiKey, baseUrl: self.BaseUrl, stage: self.Stage}
	return api.DeleteUser(serviceId, userId)
}
