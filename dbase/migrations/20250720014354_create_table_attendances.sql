-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendances (
  id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  user_id UUID NOT NULL,
  date DATE NOT NULL,
  clockin_at TIMESTAMPTZ NOT NULL,
  clockout_at TIMESTAMPTZ,
  payroll_period_id UUID,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);

CREATE UNIQUE INDEX ON attendances (user_id, date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attendances;
-- +goose StatementEnd
