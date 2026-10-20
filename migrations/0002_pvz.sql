-- +goose Up
CREATE TABLE pvz (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    city              TEXT NOT NULL
);

CREATE INDEX pvz_registration_date_id_idx ON pvz (registration_date DESC, id DESC);
-- +goose Down
DROP TABLE pvz;
