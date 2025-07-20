-- +goose Up
-- +goose StatementBegin
CREATE TABLE overtimes (
  id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  user_id UUID NOT NULL,
  date DATE NOT NULL,
  hours_taken SMALLINT NOT NULL,
  payroll_period_id UUID,
  status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, approved, rejected
  reason VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTamptz NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE overtimes;
-- +goose StatementEnd
