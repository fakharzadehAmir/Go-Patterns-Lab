package Facade

import "fmt"

// WalletFacade provides a simplified interface to access the complex subsystems:
// Account, Wallet, and Notification. It hides the complexity of interacting with
// these subsystems from the client code.
type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	notification *Notification
}

// NewWalletFacade creates a new WalletFacade instance and initializes its subsystems.
func NewWalletFacade(username string) *WalletFacade {
	return &WalletFacade{
		account:      NewAccount(username),
		wallet:       NewWallet(),
		notification: &Notification{},
	}
}

// WalletCreditBalance allows a user to credit their wallet balance.
func (wf *WalletFacade) WalletCreditBalance(username string, credit int) error {
	// Check if the user account exists.
	if !wf.account.getAccount(username) {
		return fmt.Errorf("username is not valid")
	}
	// Credit the wallet balance.
	wf.wallet.balanceCredit(credit)
	// Send a notification for the wallet credit.
	wf.notification.sendNotificationCredit()
	return nil
}

// WalletDebitBalance allows a user to debit their wallet balance.
func (wf *WalletFacade) WalletDebitBalance(username string, credit int) error {
	// Check if the user account exists.
	if !wf.account.getAccount(username) {
		return fmt.Errorf("username is not valid")
	}
	// Debit the wallet balance.
	if err := wf.wallet.balanceDebit(credit); err != nil {
		return err
	}
	// Send a notification for the wallet debit.
	wf.notification.sendNotificationCredit()
	return nil
}

// Account represents user account information.
type Account struct {
	username string
}

// NewAccount creates a new Account instance.
func NewAccount(username string) *Account {
	return &Account{username: username}
}

// getAccount checks if the provided username matches the account's username.
func (a *Account) getAccount(username string) bool {
	return username == a.username
}

// Wallet represents the user's wallet and balance.
type Wallet struct {
	balance int
}

// NewWallet creates a new Wallet instance.
func NewWallet() *Wallet {
	return &Wallet{}
}

// balanceCredit adds the specified amount to the wallet balance.
func (w *Wallet) balanceCredit(amount int) {
	w.balance += amount
}

// balanceDebit subtracts the specified amount from the wallet balance.
// It returns an error if the wallet balance is insufficient.
func (w *Wallet) balanceDebit(amount int) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	}
	return fmt.Errorf("insufficient balance for the requested debit")
}

// Notification represents the notification system for wallet operations.
type Notification struct{}

// sendNotificationCredit sends a notification for wallet credit.
func (n *Notification) sendNotificationCredit() {
	fmt.Println("Notification: Wallet credited successfully")
}

// sendNotificationDebit sends a notification for wallet debit.
func (n *Notification) sendNotificationDebit() {
	fmt.Println("Notification: Wallet debited successfully")
}
