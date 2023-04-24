package models

import (
	"context"
	"database/sql"
	"testing"

	"github.com/roronoadiogo/Simple-Bank-GoLang/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfers(t *testing.T) Transfer {

	transferGeneric := CreateTransfersParams{
		FromAccountID: sql.NullInt64{Int64: 1, Valid: true},
		ToAccountID:   sql.NullInt64{Int64: 2, Valid: true},
		Amount:        util.RandomMoney(),
	}

	createdTransfer, err := testQueries.CreateTransfers(context.Background(), transferGeneric)
	require.NoError(t, err, "Error in the test, cannot create the transfer")
	require.NotEmpty(t, createdTransfer)

	require.Equal(t, createdTransfer.FromAccountID, transferGeneric.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, transferGeneric.ToAccountID)
	require.Equal(t, createdTransfer.Amount, transferGeneric.Amount)

	require.NotZero(t, createdTransfer.ID)
	require.NotZero(t, createdTransfer.CreatedAt)

	return createdTransfer
}

func TestTransferCreateOnSuccess(t *testing.T) {

	t.Run("Create the Transfer with Success", func(t *testing.T) {
		createRandomTransfers(t)
	})
}
