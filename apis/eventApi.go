package apis

import "fmt"

type EventApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self EventApi) ListEvents(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The ListEvents method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The ListEvents method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events", self.Stage, serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventApi) CreateEvent(serviceId, userId string, eventData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The CreateEvent method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The CreateEvent method must be invoked with a non-empty string userId argument."
	}
	if !isValidEventData(eventData) {
		return "Error: The CreateEvent method must be invoked with a map[string]interface{} containing an event_type and idempotency_key keys with non-empty string values and a payload key with a non-empty map[string]interface{} value."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events", self.Stage, serviceId, userId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: eventData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventApi) GetEvent(serviceId, userId, eventId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetEvent method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The GetEvent method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(eventId) {
		return "Error: The GetEvent method must be invoked with a non-empty string eventId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events/$s", self.Stage, serviceId, userId, eventId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
