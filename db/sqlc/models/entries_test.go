package models

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/roronoadiogo/Simple-Bank-GoLang/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	account := CreateRandomAccount(t)

	entryGeneric := CreateEntriesParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), entryGeneric)
	require.NoError(t, err, "Error in the test, cannot create the entry")
	require.NotEmpty(t, entry)

	require.Equal(t, entryGeneric.AccountID, entryGeneric.AccountID)
	require.Equal(t, entryGeneric.Amount, entryGeneric.Amount)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestEntryCreateOnSuccess(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntryOnSuccess(t *testing.T) {

	entries := createRandomEntry(t)

	entriesFromDB, err := testQueries.GetEntry(context.Background(), entries.ID)
	require.NoError(t, err, "Error in the test, cannot get data in Entry table")
	require.NotEmpty(t, entriesFromDB)

	require.Equal(t, entries.Amount, entriesFromDB.Amount)
	require.Equal(t, entries.CreatedAt, entriesFromDB.CreatedAt)

}

func TestGetEntriesFromAccountOnSuccess(t *testing.T) {

	entries := createRandomEntry(t)

	entriesAdditional := CreateEntriesParams{
		AccountID: entries.AccountID,
		Amount:    util.RandomMoney(),
	}

	for i := 0; i < 5; i++ {
		testQueries.CreateEntries(context.Background(), entriesAdditional)
	}

	entriesDb := GetEntriesFromAccountParams{
		ID:     entries.AccountID.Int64,
		Limit:  5,
		Offset: 0,
	}

	getEntries, err := testQueries.GetEntriesFromAccount(context.Background(), entriesDb)
	require.NoError(t, err, "Error in the test GetEntriesFromAccount")
	require.NotEmpty(t, getEntries)

	require.Len(t, getEntries, 5)

}

func TestUpdateEntriesOnSuccess(t *testing.T) {
	entriesToDatabase := createRandomEntry(t)

	entriesToUpdate := UpdateEntriesParams{
		ID:     entriesToDatabase.ID,
		Amount: util.RandomMoney(),
	}

	entriesUpdated, err := testQueries.UpdateEntries(context.Background(), entriesToUpdate)
	require.NoError(t, err, "Error in the UpdateEntry in database")
	require.NotEmpty(t, entriesUpdated)

	require.Equal(t, entriesToDatabase.ID, entriesUpdated.ID)
	require.WithinDuration(t, entriesToDatabase.CreatedAt.Time, entriesUpdated.CreatedAt.Time, time.Second)
}

func TestDeleteEntryOnSuccess(t *testing.T) {

	entriesToDelete := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entriesToDelete.ID)
	require.NoError(t, err, "Error in delete the Entry")

	entriesFromDb, err := testQueries.GetEntry(context.Background(), entriesToDelete.ID)
	require.Error(t, err, "Error in delete entry test, the value still present in the database")
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entriesFromDb)

}
