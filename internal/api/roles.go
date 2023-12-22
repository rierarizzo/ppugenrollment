package api

func routesAllowedByRoles() map[string][]string {
	return map[string][]string{
		"/authentication":           {"M", "S", "A"},
		"/project":                  {"M", "S", "A"},
		"/project/insertNewProject": {"M"},
		"/enrollment":               {"S"},
		"/approval":                 {"A"},
	}
}
