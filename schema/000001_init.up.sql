CREATE TABLE events
(
    id SERIAL PRIMARY KEY,
    log_name TEXT,
    address_from TEXT,
    address_to TEXT,
    amount TEXT

);