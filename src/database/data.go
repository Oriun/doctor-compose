package database

import types "oriun/doctor-compose/src"

var Data = []types.SupportedDatabase{
	{
		Name:    "PostgreSQL",
		Image:   "postgres",
		Storage: "/var/lib/postgresql/data",
		TagUrl:  "https://hub.docker.com/v2/repositories/library/postgres/tags/?page_size=100&page=1",
		Port:    "5432",
		Envs: []types.Env{
			{
				Label:       "Root Password",
				VarName:     "POSTGRES_PASSWORD",
				Default:     "${RANDOM_STRING}",
				Mandatory:   true,
				Description: "This environment variable is required for you to use the PostgreSQL image. It must not be empty or undefined. This environment variable sets the superuser password for PostgreSQL. The default superuser is defined by the POSTGRES_USER environment variable.",
			},
			{
				Label:       "Set Root User",
				VarName:     "POSTGRES_USER",
				Default:     "${RANDOM_STRING}",
				Mandatory:   false,
				Description: "This optional environment variable is used in conjunction with POSTGRES_PASSWORD to set a user and its password. This variable will create the specified user with superuser power and a database with the same name. If it is not specified, then the default user of postgres will be used.",
			},
			{
				Label:       "Set Default Database Name",
				VarName:     "POSTGRES_DB",
				Default:     "doctor-postgres",
				Mandatory:   false,
				Description: "This optional environment variable can be used to define a different name for the default database that is created when the image is first started. If it is not specified, then the value of POSTGRES_USER will be used.",
			},
		},
	},
	{
		Name:    "MySQL",
		Image:   "mysql",
		Storage: "/var/lib/mysql",
		TagUrl:  "https://hub.docker.com/v2/repositories/library/mysql/tags/?page_size=100&page=1",
		Port:    "3306",
		Envs: []types.Env{
			{
				Label:       "Set Root Password",
				VarName:     "MYSQL_ROOT_PASSWORD",
				Default:     "${RANDOM_STRING}",
				Mandatory:   true,
				Description: "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw",
			},
			{
				Label:       "Set Default Database Name",
				VarName:     "MYSQL_DATABASE",
				Default:     "doctor-mysql",
				Mandatory:   false,
				Description: "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database.",
			},
			{
				Label:       "Set User",
				VarName:     "MYSQL_USER",
				Default:     "${RANDOM_STRING}",
				Mandatory:   false,
				Description: "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created.",
			},
			{
				Label:       "Set User's Password",
				VarName:     "MYSQL_PASSWORD",
				Default:     "${RANDOM_STRING}",
				Mandatory:   false,
				Description: "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created.",
			},
		},
	},
	{
		Name:    "MongoDB",
		Image:   "mongo",
		Storage: "/data/db",
		TagUrl:  "https://hub.docker.com/v2/repositories/library/mongo/tags/?page_size=100&page=1",
		Port:    "27017",
		Envs: []types.Env{
			{
				Label:       "Set Root Username",
				VarName:     "MONGO_INITDB_ROOT_USERNAME",
				Default:     "${RANDOM_STRING}",
				Mandatory:   false,
				Description: "This variable is optional and specifies the username for the root user. If it is not specified, then 'root' will be used.",
			},
			{
				Label:       "Set Root Password",
				VarName:     "MONGO_INITDB_ROOT_PASSWORD",
				Default:     "${RANDOM_STRING}",
				Mandatory:   false,
				Description: "This variable is optional and specifies the username for the root user. If it is not specified, then 'root' will be used.",
			},
			{
				Label:       "Set Default Database Name",
				VarName:     "MONGO_INITDB_DATABASE",
				Default:     "doctor-mongodb",
				Mandatory:   false,
				Description: "This variable allows you to specify the name of a database to be used for creation scripts in /docker-entrypoint-initdb.d/*.js (see Initializing a fresh instance below). MongoDB is fundamentally designed for 'create on first use', so if you do not insert data with your JavaScript files, then no database is created.",
			},
		},
	},
	{
		Name:    "Redis",
		Image:   "redis",
		Storage: "/data",
		TagUrl:  "https://hub.docker.com/v2/repositories/library/redis/tags/?page_size=100&page=1",
		Port:    "6379",
		Envs:    []types.Env{},
	},
}
