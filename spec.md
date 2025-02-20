# **API Specification for CRUD Operations on Signals, Tracks, Mileage, and TrackSignals**

## **Overview**

This API is designed to handle the creation and retrieval of entities related to **Signals**, **Tracks**, **Mileage**, and **TrackSignals**. The focus is on simplicity and extensibility, allowing for future enhancements, such as pagination, filtering, authentication, and advanced error handling.

---

## **Entities**

### **Signal Model**

```go
type Signal struct {
    ID   int     `json:"signal_id"`
    Name *string `json:"signal_name"`
    ELR  string  `json:"elr"`
}
```

- **ID**: Unique identifier (auto-incremented by DB).
- **Name**: Optional field, can be null.
- **ELR**: Required field.

### **Mileage Model**

```go
type Mileage struct {
    SignalID int      `json:"signal_id"`
    TrackID  int      `json:"track_id"`
    Mileage  *float64 `json:"mileage"`
}
```

- **SignalID**: Foreign key reference to `Signal`.
- **TrackID**: Foreign key reference to `Track`.
- **Mileage**: Optional field, can be null.

### **Track Model**

```go
type Track struct {
    ID     int    `json:"track_id"`
    Source string `json:"source"`
    Target string `json:"target"`
}
```

- **ID**: Unique identifier (auto-incremented by DB).
- **Source**: Required field.
- **Target**: Required field.

### **TrackSignals Model**

```go
type TrackSignals struct {
    ID      int    `json:"track_id"`
    Source  string `json:"source"`
    Target  string `json:"target"`
    Signals []struct {
        ID      int      `json:"signal_id"`
        Name    *string  `json:"signal_name"`
        ELR     string   `json:"elr"`
        Mileage *float64 `json:"mileage"`
    }
}
```

- **Signals**: A nested array of signal details linked to the track.

---

## **API Endpoints**

### **1. Signal Endpoints**

- **Create Signal (POST /api/v1/signals)**
  - **Input**: JSON object representing the signal.
  - **Validation**:
    - `ELR` is required.
    - `Name` is optional.
  - **Response**:
    - Returns the full Signal object with the assigned `signal_id`.
    - Status Code: `201 Created`.

- **Get Signal by ID (GET /api/v1/signals/{id})**
  - **Response**:
    - Returns the Signal object with the requested `signal_id`.
    - Status Code: `200 OK`.
    - Returns `404 Not Found` if signal doesn't exist.

- **Get All Signals (GET /api/v1/signals)**
  - **Response**:
    - Returns a list of all signals in the database.
    - Status Code: `200 OK`.
    - No pagination for now; all signals are returned.

### **2. Track Endpoints**

- **Create Track (POST /api/v1/tracks)**
  - **Input**: JSON object representing the track, with optional nested `signals`.
  - **Validation**:
    - `Source` and `Target` are required.
    - Nested signals are automatically created if they don’t exist (with `ELR` required).
  - **Response**:
    - Returns the full Track object with assigned `track_id` and associated signals.
    - Status Code: `201 Created`.

- **Get Track by ID (GET /api/v1/tracks/{id})**
  - **Response**:
    - Returns the Track object with the requested `track_id` and nested signals.
    - Status Code: `200 OK`.
    - Returns `404 Not Found` if track doesn’t exist.

- **Get All Tracks (GET /api/v1/tracks)**
  - **Response**:
    - Returns a list of all tracks in the database, with no nested signals.
    - Status Code: `200 OK`.

---

## **Data Handling**

- **Create Operation**
  - If signals are nested during Track creation, missing signals will be **auto-created** with a default `NULL` mileage.
  - If a signal already exists during creation, it will **use the existing record** rather than creating a duplicate.
  
- **Get Operation**
  - For **Get by ID** endpoints, the API will return the requested entity and **include nested data** (e.g., signals for Track by ID).
  - **Get All** endpoints will return results in the database’s **natural order** without pagination.

---

## **Error Handling**

- **Response Format**:
  - Errors will be returned in the format:

    ```json
    {
      "error": "Invalid input",
      "code": 1001
    }
    ```

- **Error Codes**:
  - **400**: Bad Request (e.g., missing required fields).
  - **404**: Not Found (e.g., entity with given ID does not exist).
  - **409**: Conflict (e.g., duplicate entity detected).
  
- **Validation Errors**: For required fields and incorrect data formats, the API will return a **simple error message**.

---

## **Testing Strategy**

- **Unit Tests**:
  - Use **`testify` for assertions**.
  - **Test cases will be created using a map** for organization.
  
- **Integration Tests**:
  - Tests will run against a **real PostgreSQL database** managed by **`dockertest`**.
  - **Tests will create their own data as needed** and clean up automatically after each test.
  
- **Test Execution**:
  - Tests will be run **sequentially** to avoid conflicts in the test database.
  - Tests will **automatically clean up** data after execution using transaction rollbacks or table truncation.

---

## **Architecture and Framework Choices**

- **Framework**: Golang's **Echo framework** will be used for routing and middleware.
- **Database**: PostgreSQL will be used, with **`dockertest`** to manage test databases.
- **Validation**: **`go-playground/validator`** will handle input validation for required fields.
- **Dependency Injection**: **Dependency injection** will be used for better testability and modularity.
- **Logging**: **Structured logging** in JSON format will be enabled.
- **Database Migrations**: **`golang-migrate`** will manage schema migrations.

---

## **Future Considerations**

- **Pagination** and **Filtering** for the `GET All` endpoints will be added later.
- **Soft Deletes** and **deletion endpoints** will be implemented in a future iteration.
- **Authentication and Authorization** (e.g., JWT, API keys) will be added later.
- **Rate Limiting** per user or more advanced rate limiting strategies can be considered as the API scales.

---

## **Other Considerations**

- **Database connection pooling** will be implemented for better performance.
- **Graceful shutdown** will be ensured to handle DB connection closures.
- **Error handling** will be simplified at first, but more detailed errors will be added later as necessary.

---
