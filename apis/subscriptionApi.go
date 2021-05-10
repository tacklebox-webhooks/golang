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
		return "Error: The listSubscriptions method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The listSubscriptions method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions", self.Stage, serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self SubscriptionApi) CreateSubscription(serviceId, userId string, subscriptionData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The createSubscription method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The createSubscription method must be invoked with a non-empty string userId argument."
	}
	if !isValidSubscriptionData(subscriptionData) {
		return "Error: The createSubscription method must be invoked with an subscriptionData argument containing non-empty event_type_id and payload properties."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions", self.Stage, serviceId, userId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: subscriptionData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self SubscriptionApi) GetSubscription(serviceId, userId, subscriptionId string) string {
	if !isValidId(serviceId) {
		return "Error: The getSubscription method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The getSubscription method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(subscriptionId) {
		return "Error: The getSubscription method must be invoked with a non-empty string subscriptionId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/subscriptions/$s", self.Stage, serviceId, userId, subscriptionId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
