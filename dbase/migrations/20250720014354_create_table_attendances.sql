-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendances (
  id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  user_id UUID NOT NULL,
  date DATE NOT NULL,
  clockin_at TIMESTAMPTZ NOT NULL,
  clockout_at TIMESTAMPTZ,
  payroll_period_id UUID,
  is_finalized BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);

CREATE UNIQUE INDEX ON attendances (user_id, date);
CREATE INDEX attendances_user ON attendances(user_id);
CREATE INDEX attendances_payroll_periods ON attendances(payroll_period_id);

ALTER TABLE attendances ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE attendances ADD FOREIGN KEY ("payroll_period_id") REFERENCES "payroll_periods" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attendances;
-- +goose StatementEnd
