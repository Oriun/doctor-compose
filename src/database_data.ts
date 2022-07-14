export default {
  supported_db: [
    "PostgreSQL",
    "MySQL",
    "MongoDB",
    "Redis",
    "MariaDB",
    "ElasticSearch"
  ],
  db_images: {
    PostgreSQL: "postgres",
    MySQL: "mysql",
    MongoDB: "mongo",
    Redis: "redis",
    MariaDB: "mariadb",
    ElasticSearch: "elasticsearch"
  },
  db_storage: {
    PostgreSQL: "postgres",
    MySQL: "/var/lib/mysql",
    MongoDB: "mongo",
    Redis: "redis",
    MariaDB: "mariadb",
    ElasticSearch: "elasticsearch"
  },
  persistModes: ["Disk", "Volume"],
  docker_hub_tags_url: {
    PostgreSQL:
      "https://hub.docker.com/v2/repositories/library/postgres/tags/?page_size=100&page=1",
    MySQL:
      "https://hub.docker.com/v2/repositories/library/mysql/tags/?page_size=100&page=1",
    MongoDB:
      "https://hub.docker.com/v2/repositories/library/mongo/tags/?page_size=100&page=1",
    Redis:
      "https://hub.docker.com/v2/repositories/library/redis/tags/?page_size=100&page=1",
    MariaDB:
      "https://hub.docker.com/v2/repositories/library/mariadb/tags/?page_size=100&page=1",
    ElasticSearch:
      "https://hub.docker.com/v2/repositories/library/elasticsearch/tags/?page_size=100&page=1"
  },
  default_ports: {
    PostgreSQL: 5432,
    MySQL: 3306,
    MongoDB: 27017,
    Redis: 6379,
    MariaDB: 3306,
    ElasticSearch: 9200
  },
  db_envs: {
    PostgreSQL: [
      {
        label: "Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "unsecure-password",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Password",
        var_name: "MYSQL_DATABASE",
        default: "${PROJECT_NAME}",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "Password",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ],
    MySQL: [
      {
        label: "Set Root Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Database",
        var_name: "MYSQL_DATABASE",
        default: "doctor-mysql",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "User",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ],
    MongoDB: [
      {
        label: "Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "unsecure-password",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Password",
        var_name: "MYSQL_DATABASE",
        default: "${PROJECT_NAME}",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "Password",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ],
    Redis: [
      {
        label: "Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "unsecure-password",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Password",
        var_name: "MYSQL_DATABASE",
        default: "${PROJECT_NAME}",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "Password",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ],
    MariaDB: [
      {
        label: "Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "unsecure-password",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Password",
        var_name: "MYSQL_DATABASE",
        default: "${PROJECT_NAME}",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "Password",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ],
    ElasticSearch: [
      {
        label: "Password",
        var_name: "MYSQL_ROOT_PASSWORD",
        default: "unsecure-password",
        mandatory: true,
        description:
          "This variable is mandatory and specifies the password that will be set for the MySQL root superuser account. In the above example, it was set to my-secret-pw"
      },
      {
        label: "Password",
        var_name: "MYSQL_DATABASE",
        default: "${PROJECT_NAME}",
        mandatory: false,
        description:
          "This variable is optional and allows you to specify the name of a database to be created on image startup. If a user/password was supplied (see below) then that user will be granted superuser access (corresponding to GRANT ALL) to this database."
      },
      {
        label: "Password",
        var_name: "MYSQL_USER",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable is optional, used in conjunction to create a new user. This user will be granted superuser permissions (see above) for the database specified by the MYSQL_DATABASE variable. Both variables are required for a user to be created."
      },
      {
        label: "Password",
        var_name: "MYSQL_PASSWORD",
        default: "${RANDOM_STRING}",
        mandatory: false,
        description:
          "This variable allows you to set a password for the user specified by the MYSQL_USER variable. This variable is required for a user to be created."
      }
    ]
  }
};
