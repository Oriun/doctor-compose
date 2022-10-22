package nodejs_data_express

import (
	"fmt"

	types "oriun/doctor-compose/src"
)

var jsPackageJson = `{
	  \"name\": \"${APP_NAME}\",
	  \"version\": \"1.0.0\",
	  \"description\": \"nodejs express backend bootstrapped with doctor-compose\",
	  \"main\": \"index.js\",
	  \"scripts\": {
		\"dev\": \"nodemon index.js\",
		\"start\": \"node index.js\"
	  },
	  \"keywords\": [],
	  \"author\": \"\",
	  \"license\": \"ISC\",
	  \"dependencies\": {
		\"dotenv\": \"^16.0.2\",
		\"express\": \"^4.17.1\"
	  },
	  \"devDependencies\": {
		\"nodemon\": \"^2.0.4\"
	  }
	}`

var tsPackageJson = `{
	  \"name\": \"${APP_NAME}\",
	  \"version\": \"1.0.0\",
	  \"description\": \"nodejs express backend bootstrapped with doctor-compose\",
	  \"main\": \"dist/index.js\",
	  \"scripts\": {
		\"dev\": \"nodemon --exec ts-node index.ts\",
		\"start\": \"node dist/index.js\",
		\"build\": \"tsc --outDir dist\",
		\"prestart\": \"npm run build\"
	  },
	  \"keywords\": [],
	  \"author\": \"\",
	  \"license\": \"ISC\",
	  \"dependencies\": {
		\"express\": \"^4.17.1\",
		\"dotenv\": \"^16.0.2\"
	  },
	  \"devDependencies\": {
		\"@types/dotenv\": \"^8.2.0\",
		\"@types/express\": \"^4.17.14\",
		\"@types/node\": \"^18.7.23\",
		\"nodemon\": \"^2.0.20\",
		\"ts-node\": \"^10.9.1\",
		\"typescript\": \"^4.8.4\"
	  }
	}`

var tsConfigJson = `{
	\"complierOptions\": {
		\"target\": \"es2016\",
		\"module\": \"commonjs\",
		\"strict\": true,
		\"esModuleInterop\": true,
		\"skipLibCheck\": true,
		\"forceConsistentCasingInFileNames\": true,
    	\"outDir\": \"dist\"
		\"allowSyntheticDefaultImports\": true,
		\"baseUrl\": \".\",
		\"paths\": {
		\"@/*\": [\"src/*\"]
		}
	},
	\"exclude\": [\"node_modules\", \"dist\"]
}`

var jsIndex = `const express = require('express')
require('dotenv').config()
const app = express()
const port = process.env.API_PORT || 3000

app.get('/', (req, res) => {
	  res.send('Hello World!')
})

app.listen(port, () => {
	  console.log(\"Example app listening at http://localhost:\" + port)
})`

var tsIndex = `import * as Express from 'express'
import * as DotEnv from 'dotenv'
DotEnv.config()
const app = Express()
const port = process.env.API_PORT || 3000

app.get('/', (req, res) => {
	  res.send('Hello World!')
})

app.listen(port, () => {
	  console.log(\"Example app listening at http://localhost:\" + port)
})`

var Data = []types.SupportedNodeFrameworks{
	{
		Name:    "Express",
		Package: "express",
		Options: map[string]types.NodeFrameWorkOption{
			"js": {
				RunCommand: "npm run start",
				Envs: []types.Env{
					{
						Label:       "NODE_ENV",
						VarName:     "NODE_ENV",
						Default:     "production",
						Mandatory:   false,
						Description: "The environment to run the application in",
					},
				},
				ManualConfig: types.ManualConfigOption{
					RunCommand: "npm run start",
					InstallCommand: []string{
						"mkdir -p ${APP_NAME} && cd ${APP_NAME}",
						fmt.Sprintf("echo \"%s\" > package.json", jsPackageJson),
						fmt.Sprintf("echo \"%s\" > index.js", jsIndex),
						"npm install",
					},
				},
				BoilerPlate: types.BoilerPlate{
					Name: "RESTful Api Boilerplate",
					Url:  "https://github.com/hagopj13/node-express-boilerplate.git",
					Tags: []string{
						"mongodb",
					},
					RunCommand:   "npm run docker:prod",
					CloneCommand: "git clone --depth 1 https://github.com/hagopj13/node-express-boilerplate.git ${APP_NAME} && cd ${APP_NAME} && npm install",
					Envs: []types.Env{
						{
							Label:       "Port number",
							VarName:     "PORT",
							Default:     "3000",
							Mandatory:   true,
							Description: "The port number to run the application on",
						},
						{
							Label:       "MongoDB URI",
							VarName:     "MONGODB_URL",
							Default:     "${DATABASE_URI}",
							Mandatory:   true,
							Description: "The URI of the MongoDB database",
						},
						{
							Label:       "JWT Secret",
							VarName:     "JWT_SECRET",
							Default:     "${RANDOM_STRING}",
							Mandatory:   true,
							Description: "The secret key used to sign the JWT",
						},
						{
							Label:       "JWT Access Token Expiration (minutes)",
							VarName:     "JWT_ACCESS_EXPIRATION_MINUTES",
							Default:     "120",
							Mandatory:   false,
							Description: "The number of minutes the access token will be valid for",
						},
						{
							Label:       "Email Server Host",
							VarName:     "SMTP_HOST",
							Mandatory:   false,
							Description: "The host of the email server. An email will be sent to verify the user's email address",
						},
						{
							Label:       "Email Server Port",
							VarName:     "SMTP_PORT",
							Mandatory:   false,
							Description: "The port of the email server. An email will be sent to verify the user's email address",
						},
						{
							Label:       "Email Server Username",
							VarName:     "SMTP_USERNAME",
							Mandatory:   false,
							Description: "The username of the email server.",
						},
						{
							Label:       "Email Server Password",
							VarName:     "SMTP_PASSWORD",
							Mandatory:   false,
							Description: "The password of the email server.",
						},
						{
							Label:       "Email Sender Address",
							VarName:     "EMAIL_FROM",
							Mandatory:   false,
							Description: "The email address from which the email will be sent.",
						},
					},
				},
			},
			"ts": {
				BuildCommand: "npm run build",
				RunCommand:   "npm run start",
				Envs: []types.Env{
					{
						Label:       "NODE_ENV",
						VarName:     "NODE_ENV",
						Default:     "development",
						Mandatory:   false,
						Description: "The environment to run the application in",
					},
				},
				ManualConfig: types.ManualConfigOption{
					RunCommand:   "npm run start",
					BuildCommand: "npm run build",
					InstallCommand: []string{
						"mkdir -p ${APP_NAME} && cd ${APP_NAME}",
						fmt.Sprintf("echo \"%s\" > package.json", tsPackageJson),
						fmt.Sprintf("echo \"%s\" > index.ts", tsIndex),
						fmt.Sprintf("echo \"%s\" > tsconfig.json", tsConfigJson),
						"npm install",
					},
				},
				BoilerPlate: types.BoilerPlate{
					Name:         "Express Typescript Boilerplate",
					Url:          "https://github.com/w3tecch/express-typescript-boilerplate.git",
					BuildCommand: "yarn start build",
					RunCommand:   "yarn start",
					CloneCommand: "git clone --depth 1 -b develop https://github.com/w3tecch/express-typescript-boilerplate.git ${APP_NAME} && cd ${APP_NAME} && npm install",
					Envs: []types.Env{
						{
							Label:       "Port number",
							VarName:     "API_PORT",
							Default:     "3000",
							Mandatory:   true,
							Description: "The port number to run the application on",
							Static:      true,
						},
						/*
							Continue to add more environment variables here from the template :
							https://github.com/w3tecch/express-typescript-boilerplate/blob/develop/.env.example
						*/
					},
				},
			},
		},
	},
}
