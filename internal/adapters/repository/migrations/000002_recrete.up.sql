DROP TABLE mileage;
DROP TABLE signals;
DROP TABLE tracks;

CREATE TABLE signals (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    elr VARCHAR(4)
);

CREATE TABLE tracks (
    id INTEGER PRIMARY KEY,
    source VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL
);

CREATE TABLE mileages (
    signal_id INTEGER NOT NULL,
    FOREIGN KEY (signal_id) REFERENCES signals (id),
    track_id INTEGER NOT NULL,
    FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE,
    mileage REAL NOT NULL,
    UNIQUE (signal_id, track_id)
);
