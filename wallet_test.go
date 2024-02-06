package jupitergo_test

import (
	"testing"

	jupitergo "github.com/ilkamo/jupiter-go"

	"github.com/test-go/testify/require"
)

func TestNewWalletFromPrivateKeyBase58(t *testing.T) {
	testPk := "5473ZnvEhn35BdcCcPLKnzsyP6TsgqQrNFpn4i2gFegFiiJLyWginpa9GoFn2cy6Aq2EAuxLt2u2bjFDBPvNY6nw"

	t.Run("valid private key", func(t *testing.T) {
		wallet, err := jupitergo.NewWalletFromPrivateKeyBase58(testPk)
		require.NoError(t, err)
		require.Equal(t, testPk, wallet.PrivateKey.String())
	})

	t.Run("invalid private key", func(t *testing.T) {
		_, err := jupitergo.NewWalletFromPrivateKeyBase58("invalid")
		require.Error(t, err)
	})
}

func TestWallet_SignTransaction(t *testing.T) {
	tx := generateTestNotSignedTx(t)
	require.Len(t, tx.Signatures, 0)

	testPk := "5473ZnvEhn35BdcCcPLKnzsyP6TsgqQrNFpn4i2gFegFiiJLyWginpa9GoFn2cy6Aq2EAuxLt2u2bjFDBPvNY6nw"

	wallet, err := jupitergo.NewWalletFromPrivateKeyBase58(testPk)
	require.NoError(t, err)

	signedTx, err := wallet.SignTransaction(tx)
	require.NoError(t, err)

	require.Len(t, signedTx.Signatures, 1)
}