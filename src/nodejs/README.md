# Nodejs Backend

## Key points

- Create docker-compsoe service, Dockerfile and .dockerignore
- if existing app, read package.json to get the list of scripts and infer the build step if needed
- if no existing app, propose to create it with a list of boilerplates

## Questions

- [x] 1. Choose a framework
- [x] 2. Set app directory or name (default cwd)
- If no existing app
  - [x] 3. Choose js or ts
  - [x] 4. Decide wether to use a boilerplate to create the app or not
  - [x] 5. Choose name
  - [ ] 6. Infer build/start commands
- If existing app
  - [ ] 7. App directory
  - [ ] 8. Infer framework then ask for confirmation
  - [ ] 9. Infer build/start commands then ask for confirmation
- [ ] 10. Expose Ports
- [ ] 11. Configure Ports
- [ ] 12. Set restart policy
- [ ] 13. Add desired environment variable
