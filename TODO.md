# Directus SDK Go - Package Update and Service Enhancement TODO

## Phase 1: Update Package Import Paths
- [ ] Update go.mod file with new module path
- [ ] Update README.md installation instructions
- [ ] Update all import statements in source files
- [ ] Update examples and documentation

## Phase 2: Add New Service Support
- [ ] Create services.go - Service management service
- [ ] Create system.go - System configuration and info service
- [ ] Create settings.go - Settings management service
- [ ] Create flow.go - Flow automation service
- [ ] Create relations.go - Relations management service
- [ ] Update Client struct to include new services
- [ ] Update NewClient function to initialize new services

## Phase 3: Enhance Existing Services
- [ ] Add missing endpoints to existing services
- [ ] Add comprehensive error handling
- [ ] Add context support for all operations
- [ ] Add type safety improvements

## Phase 4: Documentation and Testing
- [ ] Update all examples with new import path
- [ ] Add examples for new services
- [ ] Update README.md with new service documentation
- [ ] Add comprehensive test coverage

## Files to Update
- [ ] README.md
- [ ] go.mod
- [ ] examples/basic_usage.go
- [ ] example_test.go
- [ ] All .go files with import statements

## New Files to Create
- [ ] services.go
- [ ] system.go
- [ ] settings.go
- [ ] flow.go
- [ ] relations.go
