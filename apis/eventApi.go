package apis

import "fmt"

type EventApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self EventApi) ListEvents(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The listEvents method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The listEvents method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events", self.Stage, serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventApi) CreateEvent(serviceId, userId string, eventData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The createEvent method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The createEvent method must be invoked with a non-empty string userId argument."
	}
	if !isValidEventData(eventData) {
		return "Error: The createEvent method must be invoked with an eventData argument containing non-empty event_type_id and payload properties."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events", self.Stage, serviceId, userId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: eventData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventApi) GetEvent(serviceId, userId, eventId string) string {
	if !isValidId(serviceId) {
		return "Error: The getEvent method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The getEvent method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(eventId) {
		return "Error: The getEvent method must be invoked with a non-empty string eventId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s/events/$s", self.Stage, serviceId, userId, eventId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
