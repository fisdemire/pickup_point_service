-- +goose Up
CREATE TABLE products (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seq          BIGINT GENERATED ALWAYS AS IDENTITY,
    reception_id UUID NOT NULL REFERENCES receptions (id),
    date_time    TIMESTAMPTZ NOT NULL DEFAULT now(),
    type         TEXT NOT NULL
);

CREATE INDEX products_reception_seq_idx ON products (reception_id, seq);

-- +goose Down
DROP TABLE products;
