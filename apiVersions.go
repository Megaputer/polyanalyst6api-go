package polyanalyst6api

var apiVersions = map[string][]string{
	"1.0": {
		"/project/nodes",
		"/project/execution-statistics",
		"/project/execute",
		"/project/global-abort",
		"/project/save",
		"/project/unload",
		"/project/repair",
		"/project/delete",
		"/dataset/preview",
		"/scheduler/run-task",
		"/project/tasks",
		"/server/info",
		"/parameters/configure",
		"/parameters/nodes",
	},
}

func checkPathSupported(v string, url string) bool {
	for _, existingURL := range apiVersions[v] {
		if existingURL == url {
			return true
		}
	}
	return false
}

func pathSupportedIn(url string) []string {
	var res []string

	allVersions := []string{}
	for key := range apiVersions {
		allVersions = append(allVersions, key)
	}

	for _, v := range allVersions {
		if checkPathSupported(v, url) {
			res = append(res, v)
		}
	}

	return res
}
