-- +goose Up
-- +goose StatementBegin
CREATE TABLE reimbursements (
  id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  user_id UUID NOT NULL,
  amount DECIMAL(18,2) NOT NULL,
  description TEXT,
  transaction_date DATE NOT NULL DEFAULT (current_date),
  payroll_period_id UUID,
  status VARCHAR(20) NOT NULL DEFAULT 'pending',
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reimbursements;
-- +goose StatementEnd
