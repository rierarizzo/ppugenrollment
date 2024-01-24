package api

func routesAllowedByRoles() map[string][]string {
	return map[string][]string{
		"/authentication":                        {"M", "S", "A"},
		"/project":                               {"M", "S", "A"},
		"/project/insertNewProject":              {"M"},
		"/enrollment/getEnrollmentApplications":  {"A"},
		"/enrollment/enrollToProject":            {"S"},
		"/approval/approveEnrollmentApplication": {"A"},
	}
}
