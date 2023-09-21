package Facade

import (
	"testing"
)

func TestWalletFacade(t *testing.T) {
	walletFacade := NewWalletFacade("testuser")

	creditAmount := 100
	err := walletFacade.WalletCreditBalance("testuser", creditAmount)
	if err != nil {
		t.Errorf("Expected no error when crediting the wallet balance, but got: %v", err)
	}

	if walletFacade.wallet.balance != creditAmount {
		t.Errorf("Expected wallet balance to be %d, but got %d", creditAmount, walletFacade.wallet.balance)
	}

	debitAmount := 50
	err = walletFacade.WalletDebitBalance("testuser", debitAmount)
	if err != nil {
		t.Errorf("Expected no error when debiting the wallet balance, but got: %v", err)
	}

	expectedBalance := creditAmount - debitAmount
	if walletFacade.wallet.balance != expectedBalance {
		t.Errorf("Expected wallet balance to be %d, but got %d", expectedBalance, walletFacade.wallet.balance)
	}

	invalidDebitAmount := 200
	err = walletFacade.WalletDebitBalance("testuser", invalidDebitAmount)
	if err == nil {
		t.Errorf("Expected an error when debiting with insufficient balance, but got nil")
	}
}

func TestAccount(t *testing.T) {
	account := NewAccount("testuser")

	exists := account.getAccount("testuser")
	if !exists {
		t.Errorf("Expected account to exist, but it doesn't.")
	}

	exists = account.getAccount("nonexistentuser")
	if exists {
		t.Errorf("Expected account not to exist, but it does.")
	}
}

func TestWallet(t *testing.T) {
	wallet := NewWallet()

	creditAmount := 100
	wallet.balanceCredit(creditAmount)
	if wallet.balance != creditAmount {
		t.Errorf("Expected wallet balance to be %d, but got %d", creditAmount, wallet.balance)
	}

	debitAmount := 50
	err := wallet.balanceDebit(debitAmount)
	if err != nil {
		t.Errorf("Expected no error when debiting the wallet balance, but got: %v", err)
	}

	expectedBalance := creditAmount - debitAmount
	if wallet.balance != expectedBalance {
		t.Errorf("Expected wallet balance to be %d, but got %d", expectedBalance, wallet.balance)
	}

	invalidDebitAmount := 200
	err = wallet.balanceDebit(invalidDebitAmount)
	if err == nil {
		t.Errorf("Expected an error when debiting with insufficient balance, but got nil")
	}
}
