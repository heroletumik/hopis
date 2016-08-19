package session

import (
	"bitbucket.org/atticlab/horizon/cache"
	"bitbucket.org/atticlab/horizon/db2"
	"bitbucket.org/atticlab/horizon/ingest/session/ingestion"
)

// Session represents a single attempt at ingesting data into the history
// database.
type Session struct {
	Cursor    *Cursor
	Ingestion *ingestion.Ingestion

	// ClearExisting causes the session to clear existing data from the horizon db
	// when the session is run.
	ClearExisting bool

	// Metrics is a reference to where the session should record its metric information
	Metrics *IngesterMetrics

	//
	// Results fields
	//

	// Err is the error that caused this session to fail, if any.
	Err error

	// Ingested is the number of ledgers that were successfully ingested during
	// this session.
	Ingested int

	accountCache *cache.HistoryAccount
}

// NewSession initialize a new ingestion session, from `first` to `last`
func NewSession(first, last int32, horizonDB *db2.Repo, coreDB *db2.Repo, metrics *IngesterMetrics, currentVersion int) *Session {
	hdb := horizonDB.Clone()

	return &Session{
		Ingestion:    ingestion.New(horizonDB, currentVersion),
		Cursor:       NewCursor(coreDB, first, last, metrics.LoadLedgerTimer),
		Metrics:      metrics,
		accountCache: cache.NewHistoryAccount(hdb),
	}
}