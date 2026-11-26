package db

import (
	"context"
	"database/sql" // this is our USB port for connection
	"fmt"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib" // I know I'm not using any names from this package, but I need it loaded and initialized.
)

func Connect(databaseUrl string) (*sql.DB, error) {
	// sql.DB is not a single connection. It's a connection pool that manages multiple connections to your database.
	db, err:= sql.Open("pgx", databaseUrl) // pgx: is like our cable for connection, it tells our database/sql which registered driver to use.
	// sql.Open does not connects to our db. it just creates our pool object.
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err) // %w preserves the original error 
	}

	db.SetMaxOpenConns(25) // at most 25 total connections (in-use + idle) to Postgres.  If all 25 are busy, new requests block until one frees up
	db.SetMaxIdleConns(25) // Keep up to 25 idle connections warm in the pool. Since this equals MaxOpen, the pool never tears down a connection just because it's idle — it reuses it immediately on the next request.
	db.SetConnMaxLifetime(5 * time.Minute) // Any connection older than 5 minutes is closed 

	ctx, cancel:= context.WithTimeout(context.Background(), 5 * time.Second) // Creates a context that auto-cancels after 5 seconds.
	// context.Background(): The root context — an empty, never-cancelled starting point.
	// context.WithTimeout: Creates a child context that auto-cancels after 5 seconds.
	defer cancel() // ← guarantees cancel() runs when Connect() returns

	if err:= db.PingContext(ctx); err != nil { // Ping grabs a connection from the pool (creating one if needed), sends a trivial query, and returns it.
		return nil, fmt.Errorf("db.ping: %w", err)
	}
	return db, nil
}