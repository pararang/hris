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
  is_finalized BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);

CREATE INDEX reimbursements_status ON reimbursements(status);
CREATE INDEX reimbursements_user ON reimbursements(user_id);
CREATE INDEX reimbursements_payroll_periods ON reimbursements(payroll_period_id);
ALTER TABLE reimbursements ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE reimbursements ADD FOREIGN KEY ("payroll_period_id") REFERENCES "payroll_periods" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reimbursements;
-- +goose StatementEnd
