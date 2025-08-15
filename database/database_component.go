package database

import (
	"context"
	"gorm.io/gorm"
)

// WithContext
//
// Get database definition with context.
func (pc *PostgresClient) WithContext(ctx context.Context) *gorm.DB {
	return pc.WithContext(ctx)
}
