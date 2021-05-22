package apis

import "fmt"

type MessageApi struct {
	ApiKey  string
	BaseUrl string
}

func (self MessageApi) ListMessages(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The ListMessages method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The ListMessages method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("services/%s/users/%s/messages", serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self MessageApi) ResendMessage(serviceId, userId, messageId string) string {
	if !isValidId(serviceId) {
		return "Error: The ResendMessage method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The ResendMessage method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(messageId) {
		return "Error: The ResendMessage method must be invoked with a non-empty string messageId argument."
	}

	path := fmt.Sprintf("services/%s/users/%s/messages/%s/resend", serviceId, userId, messageId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self MessageApi) GetMessage(serviceId, userId, messageId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetMessage method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The GetMessage method must be invoked with a non-empty string userId argument."
	}
	if !isValidId(messageId) {
		return "Error: The GetMessage method must be invoked with a non-empty string messageId argument."
	}

	path := fmt.Sprintf("services/%s/users/%s/messages/%s", serviceId, userId, messageId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
