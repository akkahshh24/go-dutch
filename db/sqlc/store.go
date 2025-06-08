package db

type Store struct {
	*Queries
	db DBTX
}

func NewStore(db DBTX) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
