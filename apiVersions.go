package polyanalyst6api

var apiVersions = map[string][]string{
	"v0.1": {
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
	},
}

func checkAPIURL(v string, url string) bool {
	for _, existingURL := range apiVersions[v] {
		if existingURL == url {
			return true
		}
	}
	return false
}
