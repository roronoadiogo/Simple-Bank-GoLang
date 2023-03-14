package models

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/roronoadiogo/Simple-Bank-GoLang/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(1, 2000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err, "Error in the creation testing of the Account")
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccountOnSuccess(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccountOnSuccess(t *testing.T) {
	accountToDb := createRandomAccount(t)
	accountFromDb, err := testQueries.GetAccount(context.Background(), accountToDb.ID)

	require.NoError(t, err, "Error in the testing retrieving the Account from database.")
	require.NotEmpty(t, accountFromDb)

	require.Equal(t, accountFromDb.ID, accountToDb.ID)
	require.Equal(t, accountFromDb.Balance, accountFromDb.Balance)
	require.Equal(t, accountFromDb.Owner, accountFromDb.Owner)
	require.Equal(t, accountFromDb.Currency, accountFromDb.Currency)
	require.WithinDuration(t, accountToDb.CreatedAt, accountFromDb.CreatedAt, time.Second)

}

func TestUpdateAccountOnSuccess(t *testing.T) {
	accountToDb := createRandomAccount(t)

	accountUpdated := UpdateAccountParams{
		ID:      accountToDb.ID,
		Balance: util.RandomMoney(),
	}

	accountFromDb, err := testQueries.UpdateAccount(context.Background(), accountUpdated)
	require.NoError(t, err, "Error in the update testing, the Account in database.")
	require.NotEmpty(t, accountFromDb)

	require.Equal(t, accountFromDb.ID, accountToDb.ID)
	require.Equal(t, accountUpdated.Balance, accountFromDb.Balance)
	require.Equal(t, accountFromDb.Owner, accountFromDb.Owner)
	require.Equal(t, accountFromDb.Currency, accountFromDb.Currency)
	require.WithinDuration(t, accountToDb.CreatedAt, accountFromDb.CreatedAt, time.Second)

}

func TestDeleteAccountOnSuccess(t *testing.T) {
	accountToDb := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), accountToDb.ID)
	require.NoError(t, err, "Error in delete account test, the account in the database.")

	accountFromDb, err := testQueries.GetAccount(context.Background(), accountToDb.ID)
	require.Error(t, err, "Error in delete account test, the value still present in the database")
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountFromDb)

}

func TestListAccountsOnSuccess(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	listAccounts := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), listAccounts)
	require.NoError(t, err, "Error in the list account testing")
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestListAccountsOnFailure(t *testing.T) {

	db, err := sql.Open("postgres", "root@account_bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
