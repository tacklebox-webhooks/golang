package apis

import (
	"fmt"
)

func isValidId(id string) bool {
	return len(id) > 0 && fmt.Sprintf("%T", id) == "string"
}

func isValidData(data map[string]interface{}) bool {
	return len(data) > 0
}

func isValidName(data map[string]string) bool {
	return len(data["name"]) > 0 && fmt.Sprintf("%T", data["name"]) == "string"
}

func isValidSubscriptionData(data map[string]interface{}) bool {
	return len(data["url"].(string)) > 0 &&
		fmt.Sprintf("%T", data["url"].(string)) == "string" &&
		len(data["eventTypes"].([]string)) > 0
}

func isValidEventData(data map[string]interface{}) bool {
	return len(data["event_type"].(string)) > 0 &&
		fmt.Sprintf("%T", data["event_type"].(string)) == "string" &&
		len(data["idempotency_key"].(string)) > 0 &&
		fmt.Sprintf("%T", data["idempotency_key"].(string)) == "string" &&
		len(data["payload"].(map[string]string)) > 0
}
