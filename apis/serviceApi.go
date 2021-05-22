package apis

import "fmt"

type ServiceApi struct {
	ApiKey  string
	BaseUrl string
}

func (self ServiceApi) ListServices() string {
	path := fmt.Sprintf("services")
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self ServiceApi) CreateService(serviceData map[string]interface{}) string {
	if !isValidData(serviceData) {
		return "Error: The CreateService method must be invoked with a map[string]interface{} containing a name key with a non-empty string value."
	}

	path := fmt.Sprintf("services")
	request := HttpRequest{method: "POST", baseUrl: self.BaseUrl, path: path, attempt: 1, data: serviceData}
	return send(request, self.ApiKey)
}

func (self ServiceApi) GetService(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The GetService method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("services/%s", serviceId)
	request := HttpRequest{method: "GET", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}

func (self ServiceApi) DeleteService(serviceId string) string {
	if !isValidId(serviceId) {
		return "Error: The DeleteService method must be invoked with a non-empty string serviceId argument."
	}

	path := fmt.Sprintf("services/%s", serviceId)
	request := HttpRequest{method: "DELETE", baseUrl: self.BaseUrl, path: path, attempt: 1}
	return send(request, self.ApiKey)
}
