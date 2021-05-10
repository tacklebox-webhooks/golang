package apis

import "fmt"

type EventTypeApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self EventTypeApi) ListEventTypes(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The listEventTypes method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types", self.Stage, serviceId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventTypeApi) CreateEventType(serviceId string, eventTypeData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The createEventType method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidData(eventTypeData) {
		return "Error: The createEventType method must be invoked with an eventTypeData object that contains a non-empty string name property."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types", self.Stage, serviceId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: eventTypeData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self EventTypeApi) GetEventType(serviceId, eventTypeId string) string {
	if !isValidId(serviceId) {
		return "Error: The getEventType method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(eventTypeId) {
		return "Error: The getEventType method must be invoked with a non-empty string eventTypeId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/event_types/%s", self.Stage, serviceId, eventTypeId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

// async deleteEventType(serviceId, eventTypeId) {
//   const url = `${this.baseUrl}/${serviceId}/event_types/${eventTypeId}`;
//   const request = new HttpRequest("DELETE", url);
//   return await httpClient.send(request);
// }
