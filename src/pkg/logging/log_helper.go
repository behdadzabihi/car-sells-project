package logging

func mapToZapParam(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{},0)
	for k,v := range extra {
		params = append(params, string(k), v)
	}
	return params
}

func logParamsToZeroParams(keys map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for k, v := range keys {
		params[string(k)] = v
	}

	return params
}