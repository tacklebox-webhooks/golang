package components

import "github.com/tacklebox-webhooks/golang/apis"

type Service struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self Service) List() string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListServices()
}

func (self Service) Create(serviceData map[string]interface{}) string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.CreateService(serviceData)
}

func (self Service) Get(serviceId string) string {
	api := apis.ServiceApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetService(serviceId)
}

// func (self Service) Delete(serviceId string) string {
// 	   api := apis.ServiceApi{apiKey: self.ApiKey, baseUrl: self.BaseUrl, stage: self.Stage}
//     return apis.DeleteService(self, serviceId)
// }
