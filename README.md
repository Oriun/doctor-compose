# Doctor-Compose

> The CLI that diagnose your app and allow you to create the best docker-compose solution.

This tool is solo-maintained by me [@Oriun](https://github.com/Oriun) with more than [![wakatime](https://wakatime.com/badge/user/83b7f2f1-ca36-4e49-86dc-82cc48d49d70/project/bf6a5af4-fb0c-4151-bfe8-a511666859a4.svg)](https://wakatime.com/badge/user/83b7f2f1-ca36-4e49-86dc-82cc48d49d70/project/bf6a5af4-fb0c-4151-bfe8-a511666859a4) passed on the project. I use it in my projects so maybe you can use it too. Tell me if you find it useful or if you would like new features.

## Installation

#### Installation script

```bash
curl -sSL https://raw.githubusercontent.com/Oriun/doctor-compose/master/install.sh | bash
```

#### Manual

Go to the [Release page](https://github.com/Oriun/doctor-compose/releases) and download the latest version. Unzip the archive and move the binary to your bin folder.

## Roadmap

The goal of this project is to create a CLI that scan a directory and generate a `docker-compose.yml`, `.env`, and several `Dockerfile`s depending on the detected tech stack. It is also possible to generate these files via cli and it is the first goal of this project.

- [x] Generate `docker-compose.yml`
- [x] Generate `.env`
- [x] Handle some database types
- [x] Increment existing stack
- [ ] Handle some Nodejs backend stacks (Ongoing :fire:)
- [ ] Be aware of other services when creating a service
- [ ] Handle some Python backend stacks
- [ ] Handle some Go backend stacks
- [ ] Handle some Javascript frontend stacks
- [ ] Handle some Proxies like Nginx, Traefik, Apache, etc.
- [ ] Handle advance networking, configs and secrets
- Next steps: Scan a directory...

## Usage

```bash
# Go to the root of the project
cd my-project
doctor-compose
```

## Outputs

- `docker-compose.yml` with version ^3.9 and services. Currently no support for `networks`, `configs` or `secrets`.
  - services with `image`, `volumes`, `restart`, `ports` and `environment`.
  - services have a comment that explains roughly the purpose of the service.
  - if `.env` is generated, `service.environment` remaps the variable to conform to the service.
- `.env` if selected in the CLI. Variables names are prepended with the service name in case of duplicate variables.

###### Incoming :

- `Dockerfile` depending on the detected tech stack
