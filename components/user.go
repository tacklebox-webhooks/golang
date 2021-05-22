package components

import "github.com/tacklebox-webhooks/golang/apis"

type User struct {
	ApiKey  string
	BaseUrl string
}

func (self User) List(serviceId string) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.ListUsers(serviceId)
}

func (self User) Create(serviceId string, userData map[string]interface{}) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.CreateUser(serviceId, userData)
}

func (self User) Get(serviceId, userId string) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.GetUser(serviceId, userId)
}

func (self User) Delete(serviceId, userId string) string {
	api := apis.UserApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.DeleteUser(serviceId, userId)
}
