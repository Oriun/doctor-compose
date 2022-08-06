package nodejs_data

import (
	types "oriun/doctor-compose/src"
	nodejs_data_express "oriun/doctor-compose/src/nodejs/data/express"
)

var Data = append([]types.SupportedNodeFrameworks{}, nodejs_data_express.Data...)
