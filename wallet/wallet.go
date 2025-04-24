package wallet

import (
	"errors"
	"sync"
	"time"
)

// TransactionType represents the type of transaction
type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Withdraw TransactionType = "WITHDRAW"
	Transfer TransactionType = "TRANSFER"
)

// Transaction represents a single transaction record
type Transaction struct {
	Type      TransactionType
	Amount    float64
	From      string
	To        string
	Timestamp time.Time
}

// Wallet represents a user's wallet with balance and transaction history
type Wallet struct {
	UserID       string
	Balance      float64
	mu           sync.RWMutex
	transactions []Transaction
}

// WalletManager manages all wallets in the system
type WalletManager struct {
	wallets map[string]*Wallet
	mu      sync.RWMutex
}

// NewWalletManager creates a new wallet manager
func NewWalletManager() *WalletManager {
	return &WalletManager{
		wallets: make(map[string]*Wallet),
	}
}

// CreateWallet creates a new wallet for a user
func (wm *WalletManager) CreateWallet(userID string) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	if _, exists := wm.wallets[userID]; exists {
		return errors.New("wallet already exists for user")
	}

	wm.wallets[userID] = &Wallet{
		UserID:       userID,
		Balance:      0,
		transactions: make([]Transaction, 0),
	}
	return nil
}

// GetWallet returns a wallet for a user
func (wm *WalletManager) GetWallet(userID string) (*Wallet, error) {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	wallet, exists := wm.wallets[userID]
	if !exists {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}

// Deposit adds money to a wallet
func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	w.Balance += amount
	w.transactions = append(w.transactions, Transaction{
		Type:      Deposit,
		Amount:    amount,
		To:        w.UserID,
		Timestamp: time.Now(),
	})
	return nil
}

// Withdraw removes money from a wallet
func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.Balance < amount {
		return errors.New("insufficient funds")
	}

	w.Balance -= amount
	w.transactions = append(w.transactions, Transaction{
		Type:      Withdraw,
		Amount:    amount,
		From:      w.UserID,
		Timestamp: time.Now(),
	})
	return nil
}

// GetBalance returns the current balance of a wallet
func (w *Wallet) GetBalance() float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.Balance
}

// GetTransactionHistory returns the transaction history of a wallet
func (w *Wallet) GetTransactionHistory() []Transaction {
	w.mu.RLock()
	defer w.mu.RUnlock()

	// Return a copy of the transactions to prevent external modification
	transactions := make([]Transaction, len(w.transactions))
	copy(transactions, w.transactions)
	return transactions
}

// SendMoney transfers money from one wallet to another
func (wm *WalletManager) SendMoney(fromUserID, toUserID string, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	wm.mu.Lock()
	defer wm.mu.Unlock()

	fromWallet, exists := wm.wallets[fromUserID]
	if !exists {
		return errors.New("sender wallet not found")
	}

	toWallet, exists := wm.wallets[toUserID]
	if !exists {
		return errors.New("recipient wallet not found")
	}

	fromWallet.mu.Lock()
	if fromWallet.Balance < amount {
		fromWallet.mu.Unlock()
		return errors.New("insufficient funds")
	}
	fromWallet.Balance -= amount
	fromWallet.transactions = append(fromWallet.transactions, Transaction{
		Type:      Transfer,
		Amount:    amount,
		From:      fromUserID,
		To:        toUserID,
		Timestamp: time.Now(),
	})
	fromWallet.mu.Unlock()

	toWallet.mu.Lock()
	toWallet.Balance += amount
	toWallet.transactions = append(toWallet.transactions, Transaction{
		Type:      Transfer,
		Amount:    amount,
		From:      fromUserID,
		To:        toUserID,
		Timestamp: time.Now(),
	})
	toWallet.mu.Unlock()

	return nil
}
