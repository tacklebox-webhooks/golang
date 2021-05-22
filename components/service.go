package components

import "github.com/tacklebox-webhooks/golang/apis"

type Service struct {
	ApiKey  string
	BaseUrl string
}

func (self Service) List() string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.ListServices()
}

func (self Service) Create(serviceData map[string]interface{}) string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.CreateService(serviceData)
}

func (self Service) Get(serviceId string) string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.GetService(serviceId)
}

func (self Service) Delete(serviceId string) string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl}
	return api.DeleteService(serviceId)
}
