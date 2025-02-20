# railway-signals

This API is designed to handle the creation and retrieval of entities related to **Signals**, **Tracks**, **Mileage**, and **TrackSignals**. The focus is on simplicity and extensibility, allowing for future enhancements, such as pagination, filtering, authentication, and advanced error handling.

> [!NOTE]
> I somehow messed up the database migrations and couldn't fix it, so testing is an issue.
> I also need to add a load more unit tests, added some simple ones to show how I would do it but there's not enough.

> [!NOTE]
> I tried to learn the echo framework and go-pg orm while writing this so my usage might be a little niave.

---

## **Entities**

### **Signal Model**

```go
type Signal struct {
    ID   int     `json:"id"`
    Name string  `json:"signal_name"`
    ELR  string  `json:"elr"`
}
```

### **Mileage Model**

```go
type Mileage struct {
    SignalID int      `json:"signal_id"`
    TrackID  int      `json:"track_id"`
    Mileage  *float64 `json:"mileage"`
}
```

### **Track Model**

```go
type Track struct {
    ID     int    `json:"id"`
    Source string `json:"source"`
    Target string `json:"target"`
}
```

### **TrackSignals Model**

```go
type TrackSignals struct {
    ID      int    `json:"track_id"`
    Source  string `json:"source"`
    Target  string `json:"target"`
    Signals []struct {
        ID      int      `json:"signal_id"`
        Name    string   `json:"signal_name"`
        ELR     string   `json:"elr"`
        Mileage float64  `json:"mileage"`
    }
}
```

---

## **API Endpoints**

### **1. Signal Endpoints**

- **Create Signal (POST /api/v1/signals)**
  - **Input**: JSON object representing the signal.
  - **Validation**:
    - `ELR` is optional.
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

## Service designed

I tried to use a hexagonal architecture along with DDD. Separating the business logic from the adapters as much as possible

- <https://www.geeksforgeeks.org/hexagonal-architecture-system-design/>
- <https://www.geeksforgeeks.org/domain-driven-design-ddd/?ref=header_outind>

---

## **Future Considerations**

- **Unit tests** Need to add a lot more tests to get better coverage.
- **Pagination** and **Filtering** for the `GET All` endpoints will be added later.
- **Soft Deletes** and **deletion endpoints** will be implemented in a future iteration.
- **Authentication and Authorization** (e.g., JWT, API keys) will be added later.
- **Rate Limiting** per user or more advanced rate limiting strategies can be considered as the API scales.
- **Database connection pooling** will be implemented for better performance.
- **Graceful shutdown** will be ensured to handle DB connection closures.
- **Error handling** will be simplified at first, but more detailed errors will be added later as necessary.
