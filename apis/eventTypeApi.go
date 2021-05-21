package apis

import "fmt"

type EventTypeApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self EventTypeApi) ListEventTypes(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The ListEventTypes method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types", self.Stage, serviceId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventTypeApi) CreateEventType(serviceId string, eventTypeData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The CreateEventType method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidData(eventTypeData) {
		return "Error: The CreateEventType method must be invoked with a map[string]interface{} that contains a name key with non-empty string value."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types", self.Stage, serviceId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: eventTypeData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventTypeApi) GetEventType(serviceId, eventTypeId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetEventType method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(eventTypeId) {
		return "Error: The GetEventType method must be invoked with a non-empty string eventTypeId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types/%s", self.Stage, serviceId, eventTypeId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventTypeApi) DeleteEventType(serviceId, eventTypeId string) string {
	if !isValidId(serviceId) {
		return "Error: The DeleteEventType method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(eventTypeId) {
		return "Error: The DeleteEventType method must be invoked with a non-empty string eventTypeId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types/%s", self.Stage, serviceId, eventTypeId)
	request := HttpRequest{method: "DELETE", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

// async deleteEventType(serviceId, eventTypeId) {
//   const url = `${this.baseUrl}/${serviceId}/event_types/${eventTypeId}`;
//   const request = new HttpRequest("DELETE", url);
//   return await httpClient.send(request);
// }
