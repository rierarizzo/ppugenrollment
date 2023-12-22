package api

func routesAllowedByRoles() map[string][]string {
	return map[string][]string{
		"/authentication": {"ALL"},
		"/project":        {"ALL"},
		"/enrollment":     {"S"},
		"/approval":       {"A"},
	}
}
