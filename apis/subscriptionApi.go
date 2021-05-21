package apis

import (
	"fmt"
)

type SubscriptionApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self SubscriptionApi) ListSubscriptions(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The ListSubscriptions method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The ListSubscriptions method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions", self.Stage, serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self SubscriptionApi) CreateSubscription(serviceId, userId string, subscriptionData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The CreateSubscription method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The CreateSubscription method must be invoked with a non-empty string userId argument."
	}
	if !isValidSubscriptionData(subscriptionData) {
		return "Error: The CreateSubscription method must be invoked with a map[string]interface{} containing an eventTypes key with a non-empty []string value and a url key with a non-empty string value."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions", self.Stage, serviceId, userId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: subscriptionData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self SubscriptionApi) GetSubscription(serviceId, userId, subscriptionId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetSubscription method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The GetSubscription method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(subscriptionId) {
		return "Error: The GetSubscription method must be invoked with a non-empty string subscriptionId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions/$s", self.Stage, serviceId, userId, subscriptionId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self SubscriptionApi) DeleteSubscription(serviceId, userId, subscriptionId string) string {
	if !isValidId(serviceId) {
		return "Error: The DeleteSubscription method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The DeleteSubscription method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(subscriptionId) {
		return "Error: The DeleteSubscription method must be invoked with a non-empty string subscriptionId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions/$s", self.Stage, serviceId, userId, subscriptionId)
	request := HttpRequest{method: "DELETE", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
