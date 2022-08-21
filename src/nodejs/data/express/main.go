package nodejs_data_express

import types "oriun/doctor-compose/src"

var Data = []types.SupportedNodeFrameworks{
	{
		Name:    "Express",
		Package: "express",
		Version: map[string]types.NodeFrameWorkVerions{
			"^4.0.0": {
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
				BoilerPlate: []types.BoilerPlate{
					{
						Name: "RESTful Api Boilerplate",
						Url:  "https://github.com/hagopj13/node-express-boilerplate.git",
						Tags: []string{
							"mongodb",
						},
						RunCommand:   "npm run docker:prod",
						CloneCommand: "git clone --depth 1 https://github.com/hagopj13/node-express-boilerplate.git ${APP_NAME}",
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
			},
		},
	},
	{
		Name:    "Express - typescript",
		Package: "express",
		Version: map[string]types.NodeFrameWorkVerions{
			"^4.0.0": {
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
				BoilerPlate: []types.BoilerPlate{
					{
						Name:         "Express Typescript Boilerplate",
						Url:          "https://github.com/w3tecch/express-typescript-boilerplate.git",
						BuildCommand: "yarn start build",
						RunCommand:   "yarn start",
						CloneCommand: "git clone -b develop https://github.com/w3tecch/express-typescript-boilerplate.git ${APP_NAME} && rm -rf ${APP_NAME}/.git",
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
	},
}
