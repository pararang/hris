-- +goose Up
-- +goose StatementBegin
CREATE TABLE payslips (
  id UUID PRIMARY KEY DEFAULT (gen_random_uuid()),
  user_id UUID NOT NULL,
  payroll_period_id UUID NOT NULL,
  base_salary DECIMAL(18,2) NOT NULL,
  prorated_base_salary DECIMAL(18,2) NOT NULL,
  overtime_pay DECIMAL(18,2) NOT NULL,
  reimbursement_amount DECIMAL(18,2) NOT NULL,
  take_home_pay DECIMAL(18,2) NOT NULL,
  details_json JSONB,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255)
);

CREATE INDEX payslips_user ON payslips(user_id);
CREATE INDEX payslips_payroll_periods ON payslips(payroll_period_id);
CREATE UNIQUE INDEX ON payslips (user_id, payroll_period_id);
ALTER TABLE payslips ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE payslips ADD FOREIGN KEY ("payroll_period_id") REFERENCES "payroll_periods" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payslips;
-- +goose StatementEnd
