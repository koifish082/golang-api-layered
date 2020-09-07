package library

// FilterNotEmptyValue filter not empty value
func FilterNotEmptyValue(m map[string]string) map[string]string {
	res := map[string]string{}
	for k, v := range m {
		if v != "" {
			res[k] = v
		}
	}
	return res
}
