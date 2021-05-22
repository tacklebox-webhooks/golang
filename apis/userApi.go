package apis

import "fmt"

type UserApi struct {
	ApiKey  string
	BaseUrl string
}

func (self UserApi) ListUsers(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The ListUsers method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("services/%s/users", serviceId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self UserApi) CreateUser(serviceId string, userData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The CreateUser method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidData(userData) {
		return "Error: The CreateUser method must be invoked with a map[string]interface{} containing a name key with a non-empty string value."
	}

	path := fmt.Sprintf("services/%s/users", serviceId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: userData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self UserApi) GetUser(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetUser method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The GetUser method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("services/%s/users/%s", serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self UserApi) DeleteUser(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The DeleteUser method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The DeleteUser method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("services/%s/users/%s", serviceId, userId)
	request := HttpRequest{method: "DELETE", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
