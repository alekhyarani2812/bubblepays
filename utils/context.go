package utils

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string
const (
	TransactionIDKey ContextKey = "transaction_id"
)
func GenerateTransactionID() string {
	// Generate a UUID for the transaction ID
	id := uuid.New()
	return id.String()
}
func WithTransactionID(ctx context.Context) context.Context {
	// Generate a transaction ID and add it to the context
	transactionID := GenerateTransactionID()
	return context.WithValue(ctx, TransactionIDKey, transactionID)
}
func GetTransactionID(ctx context.Context) string {
	// Retrieve the transaction ID from the context
	if id, ok := ctx.Value(TransactionIDKey).(string); ok {
		return id
	}
	return ""
}
