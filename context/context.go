package context

import (
	"context"
	"errors"
	"log"

	"github.com/bubblepays/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type contextKey string
const (
	//TransactionIDKey is used for represent http request for tracing an issue
	TransactionIDKey contextKey = "transaction_id"
)

type dbConnection string
const (
	//DBconnectionKey used for getting db connection
	DBconnectionKey dbConnection = "dbConnection"
)
// GenerateTransactionID to generate uuid ID
func GenerateTransactionID() string {
	// Generate a UUID for the transaction ID
	id := uuid.New()
	return id.String()
}
// WithTransactionID for with transactionIDKey and transactionID
func WithTransactionID(ctx context.Context) context.Context {
	// Generate a transaction ID and add it to the context
	transactionID := GenerateTransactionID()
	return context.WithValue(ctx, TransactionIDKey, transactionID)
}
// WithDBConnection for dbConnection with context 
func WithDBConnection(ctx context.Context) context.Context {
	dbConnection, dbErr := database.GetORMDbConnection()
	if dbErr != nil {
		log.Printf(" WithDBConnection err %v ", dbErr.Error())
		return nil
	}
	ctx = context.WithValue(ctx, DBconnectionKey, dbConnection)
	return ctx
}
// GetTransactionID to get transactionId
func GetTransactionID(ctx context.Context) string {
	// Retrieve the transaction ID from the context
	if id, ok := ctx.Value(TransactionIDKey).(string); ok {
		return id
	}
	return ""
}
// GetDBConnection to get db connection
func GetDBConnection(ctx context.Context) (*gorm.DB, error) {
	// Retrieve the transaction ID from the context
	if db, ok := ctx.Value(DBconnectionKey).(*gorm.DB); ok {
		return db, nil
	}
	return nil, errors.New("DB connection does not exist in context")
}
