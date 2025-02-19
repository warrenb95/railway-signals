CREATE TABLE signals (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    elr VARCHAR(4) NOT NULL
);

CREATE TABLE tracks (
    id VARCHAR(36) PRIMARY KEY,
    source VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL
);

CREATE TABLE mileages (
    signal_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (signal_id) REFERENCES signals (id),
    track_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE,
    mileage REAL NOT NULL,
    UNIQUE (signal_id, track_id)
);
