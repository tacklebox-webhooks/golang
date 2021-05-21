package components

import "github.com/tacklebox-webhooks/golang/apis"

type Subscription struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self Subscription) List(serviceId, userId string) string {
	api := apis.SubscriptionApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.ListSubscriptions(serviceId, userId)
}

func (self Subscription) Create(serviceId, userId string, subscriptionData map[string]interface{}) string {
	api := apis.SubscriptionApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.CreateSubscription(serviceId, userId, subscriptionData)
}

func (self Subscription) Get(serviceId, userId, subscriptionId string) string {
	api := apis.SubscriptionApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.GetSubscription(serviceId, userId, subscriptionId)
}

func (self Subscription) Delete(serviceId, userId, subscriptionId string) string {
	api := apis.SubscriptionApi{ApiKey: self.ApiKey, BaseUrl: self.BaseUrl, Stage: self.Stage}
	return api.DeleteSubscription(serviceId, userId, subscriptionId)
}
