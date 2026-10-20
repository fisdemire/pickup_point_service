-- +goose Up
CREATE TABLE receptions (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pvz_id    UUID NOT NULL REFERENCES pvz (id),
    date_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    status    TEXT NOT NULL DEFAULT 'in_progress' CHECK (status IN ('in_progress', 'close'))
);

CREATE UNIQUE INDEX one_open_reception ON receptions (pvz_id) WHERE status = 'in_progress';

CREATE INDEX receptions_pvz_date_idx ON receptions (pvz_id, date_time DESC);

-- +goose Down
DROP TABLE receptions;
