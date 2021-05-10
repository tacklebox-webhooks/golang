package apis

import "fmt"

type UserApi struct {
	ApiKey  string
	BaseUrl string
	Stage   string
}

func (self UserApi) ListUsers(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The listUsers method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users", self.Stage, serviceId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self UserApi) CreateUser(serviceId string, userData map[string]interface{}) string {
	if !isValidId(serviceId) {
		return "Error: The createUser method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidData(userData) {
		return "Error: The createUser method must be invoked with an userData object that contains a non-empty string name property."
	}

	path := fmt.Sprintf("/%s/services/%s/users", self.Stage, serviceId)
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, data: userData, attempt: 1}
	return send(request, self.ApiKey)
}

func (self UserApi) GetUser(serviceId, userId string) string {
	if !isValidId(serviceId) {
		return "Error: The getUser method must be invoked with a non-empty string serviceId argument."
	}
	if !isValidId(userId) {
		return "Error: The getUser method must be invoked with a non-empty string userId argument."
	}

	path := fmt.Sprintf("/%s/services/%s/users/%s", self.Stage, serviceId, userId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
