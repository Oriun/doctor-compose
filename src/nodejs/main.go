package nodejs

import (
	"fmt"
	types "oriun/doctor-compose/src"
	nodejs_data "oriun/doctor-compose/src/nodejs/data"
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

	fmt.Println(getNames(nodejs_data.Data))
	/*
	 * Do things here
	 */

	return name, service, env_string
}
