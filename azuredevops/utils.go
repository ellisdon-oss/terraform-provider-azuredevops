package azuredevops

func convertInterfaceToStringMap(val map[string]interface{}) map[string]string {
	resultMap := make(map[string]string)

	for key, value := range val {
		strKey := key
		strValue := value.(string)

		resultMap[strKey] = strValue
	}

	return resultMap
}
