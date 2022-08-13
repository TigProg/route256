-- +goose Up
-- +goose StatementBegin
INSERT INTO public.bus_station (id, name, city) VALUES
    (1, 'Ufa_main', 'Ufa'),
    (2, 'SPb_1', 'SPb'),
    (3, 'SPb_2', 'SPb'),
    (4, 'Moscow_1', 'Moscow'),
    (5, 'Moscow_2', 'Moscow'),
    (6, 'Moscow_3', 'Moscow'),
    (7, 'nowhere', null);

INSERT INTO public.route (id, name, departure_id, destination_id) VALUES
    (1, 'ufaspb', 1, 2),
    (2, 'spbufa', 2, 1),
    (3, 'ufamsk', 1, 4),
    (4, 'mskufa', 4, 1),
    (5, 'spbmsk', 3, 6),
    (6, 'mskspb', 6, 3),
    (7, 'ufanowhere', 1, 7),
    (8, 'nowhereufa', 7, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM public.route
WHERE id in (1, 2, 3, 4, 5, 6, 7, 8);

DELETE
FROM public.bus_station
WHERE id in (1, 2, 3, 4, 5, 6, 7);
-- +goose StatementEnd
