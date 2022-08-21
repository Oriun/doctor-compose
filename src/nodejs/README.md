# Nodejs Backend

## Key points

- Create docker-compsoe service, Dockerfile and .dockerignore
- if existing app, read package.json to get the list of scripts and infer the build step if needed
- if no existing app, propose to create it with a list of boilerplates

## Questions

- [x] 1. Choose a framework
- [x] 2. Set app directory or name (default cwd)
- If no existing app
  - [x] 3. Choose a boilerplate to create app (or not)
  - [x] 4. Choose name
  - [ ] 5. Infer build/start commands
- If existing app
  - [ ] 6. App directory
  - [ ] 7. Infer framework then ask for confirmation
  - [ ] 8. Infer build/start commands then ask for confirmation
- [ ] 9. Expose Ports
- [ ] 10. Configure Ports
- [ ] 11. Set restart policy
- [ ] 12. Add desired environment variable
