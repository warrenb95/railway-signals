# CRUD API Project - To-Do List

## Step 1: Initial Setup & Planning

- [x] Install Golang and dependencies
- [x] Install Docker and PostgreSQL
- [x] Initialize Go project
- [x] Set up folder structure (`cmd`, `pkg`, `internal`, `models`, `api`)
- [x] Set up version control (Git)
- [x] Create `.gitignore`
- [x] Add Go modules (`Echo`, `dockertest`, `pq`)
- [x] Define `Signal`, `Track`, `Mileage`, `TrackSignal` models
- [x] Design database schema for models
- [x] Create migrations for PostgreSQL schema
- [x] Create `docker-compose.yml` for PostgreSQL container

## Step 2: Develop the API's Core CRUD Endpoints

- [x] Set up Echo framework and basic server
- [ ] Create `/health` endpoint for health check
- [x] Implement `Create`, `Get`, `List`, `Update` and `Delete` for `Signal`
- [ ] Add validation logic for `Signal` data
- [ ] Write tests for `Signal` API endpoints
- [x] Implement `Create`, `Get`, `List`, `Update` and `Delete` for `Track`
- [ ] Add validation logic for `Track` data
- [ ] Write tests for `Track` API endpoints
- [ ] Implement data loading for `TrackSignals`
- [ ] Add validation logic for `TrackSignal` data
- [ ] Write tests for `TrackSignals` API endpoints

## Step 3: Database Integration and Dependency Injection

- [ ] Set up database connection pool
- [x] Write functions for interacting with PostgreSQL (`Create`, `Get`)
- [x] Use dependency injection to inject DB connection
- [x] Integrate structured logging (e.g., `logrus`)
- [x] Add logs for API requests and responses
- [ ] Write tests for DB integration and logging

## Step 4: Additional Features (Pagination, Filtering, Soft Deletes)

- [ ] Implement pagination for `Get` endpoints
- [ ] Write tests for pagination
- [ ] Implement filtering for `Get` endpoints
- [ ] Write tests for filtering
- [ ] Implement soft deletes for models
- [ ] Write tests for soft delete functionality

## Step 5: Security, Background Jobs, and Performance Enhancements

- [ ] Set up authentication (JWT)
- [ ] Add middleware for protected routes
- [ ] Write tests for authentication and authorization
- [ ] Implement background job processing (optional)
- [ ] Optimize database queries and add indexing
- [ ] Write performance tests

## Step 6: Testing and Final Integration

- [ ] Write unit tests for DB functions
- [ ] Write unit tests for handler functions
- [ ] Write end-to-end tests using `dockertest`
- [ ] Integrate all components and run tests in CI/CD pipeline

## Step 7: Documentation and Clean-Up

- [ ] Create API documentation (endpoints, data models)
- [ ] Final code review and refactor
- [ ] Ensure consistency in error handling and naming conventions
- [ ] Prepare for deployment
