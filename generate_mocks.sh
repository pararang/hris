#!/bin/bash

cd /Users/ikhsan/Projects/techtest/hris-api

# Create mocks directory if it doesn't exist
mkdir -p mocks

# Generate mocks for each repository interface
mockery --srcpkg github.com/pararang/hris/domain/repository --name AttendanceRepository --output mocks
mockery --srcpkg github.com/pararang/hris/domain/repository --name OvertimeRepository --output mocks
mockery --srcpkg github.com/pararang/hris/domain/repository --name ReimbursementRepository --output mocks
mockery --srcpkg github.com/pararang/hris/domain/repository --name PayslipRepository --output mocks
mockery --srcpkg github.com/pararang/hris/domain/repository --name UserRepository --output mocks
mockery --srcpkg github.com/pararang/hris/domain/repository --name AuditRepository --output mocks
