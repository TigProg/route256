-- +goose Up
-- +goose StatementBegin
-- public.bus_station
CREATE TABLE IF NOT EXISTS public.bus_station (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(64),
    city VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS public.route (
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(10) NOT NULL UNIQUE,
    departure_id   INTEGER NOT NULL REFERENCES bus_station(id),
    destination_id INTEGER NOT NULL REFERENCES bus_station(id),
    CHECK (departure_id != destination_id)
);

CREATE TABLE IF NOT EXISTS public.booking (
    id       SERIAL PRIMARY KEY,
    route_id INTEGER NOT NULL REFERENCES route(id),
    date     DATE NOT NULL,
    seat     INTEGER NOT NULL,
    CONSTRAINT booking_uniq UNIQUE(route_id, date, seat)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.booking;
DROP TABLE IF EXISTS public.route;
DROP TABLE IF EXISTS public.bus_station;
-- +goose StatementEnd
