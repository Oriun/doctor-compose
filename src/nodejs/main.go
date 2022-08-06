package nodejs

import (
	types "oriun/doctor-compose/src"
)

func getNames(vs []types.SupportedNodeFrameworks) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v.Name
	}
	return vsm
}

func GetService() (string, types.Service, string) {
	service := types.Service{}
	var name = ""
	var env_string = ""

	/*
	 * Do things here
	 */

	return name, service, env_string
}
