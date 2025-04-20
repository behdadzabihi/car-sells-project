package logging

func mapToZapParam(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{},0)
	for k,v := range extra {
		params = append(params, string(k), v)
	}
	return params
}
