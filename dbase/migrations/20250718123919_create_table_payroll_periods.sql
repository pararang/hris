-- +goose Up
-- +goose StatementBegin
CREATE TABLE payroll_periods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, processing, paid
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255)
);

CREATE INDEX idx_payroll_periods_status ON payroll_periods (status);
CREATE INDEX idx_payroll_periods_dates ON payroll_periods (start_date, end_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payroll_periods;
-- +goose StatementEnd
