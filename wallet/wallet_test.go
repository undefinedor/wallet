package wallet

import (
	"testing"
)

func TestWalletManager_CreateWallet(t *testing.T) {
	wm := NewWalletManager()

	// Test creating a new wallet
	err := wm.CreateWallet("user1")
	if err != nil {
		t.Errorf("Failed to create wallet: %v", err)
	}

	// Test creating a duplicate wallet
	err = wm.CreateWallet("user1")
	if err == nil {
		t.Error("Expected error when creating duplicate wallet")
	}
}

func TestWallet_Deposit(t *testing.T) {
	wm := NewWalletManager()
	wm.CreateWallet("user1")
	wallet, _ := wm.GetWallet("user1")

	// Test valid deposit
	err := wallet.Deposit(100)
	if err != nil {
		t.Errorf("Failed to deposit: %v", err)
	}
	if wallet.GetBalance() != 100 {
		t.Errorf("Expected balance 100, got %f", wallet.GetBalance())
	}

	// Test invalid deposit
	err = wallet.Deposit(-50)
	if err == nil {
		t.Error("Expected error when depositing negative amount")
	}
}

func TestWallet_Withdraw(t *testing.T) {
	wm := NewWalletManager()
	wm.CreateWallet("user1")
	wallet, _ := wm.GetWallet("user1")
	wallet.Deposit(100)

	// Test valid withdrawal
	err := wallet.Withdraw(50)
	if err != nil {
		t.Errorf("Failed to withdraw: %v", err)
	}
	if wallet.GetBalance() != 50 {
		t.Errorf("Expected balance 50, got %f", wallet.GetBalance())
	}

	// Test insufficient funds
	err = wallet.Withdraw(100)
	if err == nil {
		t.Error("Expected error when withdrawing more than balance")
	}

	// Test invalid withdrawal
	err = wallet.Withdraw(-50)
	if err == nil {
		t.Error("Expected error when withdrawing negative amount")
	}
}

func TestWalletManager_SendMoney(t *testing.T) {
	wm := NewWalletManager()
	wm.CreateWallet("user1")
	wm.CreateWallet("user2")
	wallet1, _ := wm.GetWallet("user1")
	wallet2, _ := wm.GetWallet("user2")
	wallet1.Deposit(100)

	// Test valid transfer
	err := wm.SendMoney("user1", "user2", 50)
	if err != nil {
		t.Errorf("Failed to send money: %v", err)
	}
	if wallet1.GetBalance() != 50 {
		t.Errorf("Expected sender balance 50, got %f", wallet1.GetBalance())
	}
	if wallet2.GetBalance() != 50 {
		t.Errorf("Expected recipient balance 50, got %f", wallet2.GetBalance())
	}

	// Test insufficient funds
	err = wm.SendMoney("user1", "user2", 100)
	if err == nil {
		t.Error("Expected error when sending more than balance")
	}

	// Test invalid transfer
	err = wm.SendMoney("user1", "user2", -50)
	if err == nil {
		t.Error("Expected error when sending negative amount")
	}

	// Test non-existent wallet
	err = wm.SendMoney("user1", "user3", 10)
	if err == nil {
		t.Error("Expected error when sending to non-existent wallet")
	}
}

func TestWallet_GetBalance(t *testing.T) {
	wm := NewWalletManager()
	wm.CreateWallet("user1")
	wallet, _ := wm.GetWallet("user1")

	// Test initial balance
	if wallet.GetBalance() != 0 {
		t.Errorf("Expected initial balance 0, got %f", wallet.GetBalance())
	}

	// Test balance after deposit
	wallet.Deposit(100)
	if wallet.GetBalance() != 100 {
		t.Errorf("Expected balance 100, got %f", wallet.GetBalance())
	}
}

func TestWallet_GetTransactionHistory(t *testing.T) {
	wm := NewWalletManager()
	wm.CreateWallet("user1")
	wallet, _ := wm.GetWallet("user1")

	// Test empty transaction history
	history := wallet.GetTransactionHistory()
	if len(history) != 0 {
		t.Errorf("Expected empty transaction history, got %d transactions", len(history))
	}

	// Test deposit transaction
	wallet.Deposit(100)
	history = wallet.GetTransactionHistory()
	if len(history) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(history))
	}
	if history[0].Type != Deposit {
		t.Errorf("Expected transaction type DEPOSIT, got %s", history[0].Type)
	}
	if history[0].Amount != 100 {
		t.Errorf("Expected transaction amount 100, got %f", history[0].Amount)
	}

	// Test withdraw transaction
	wallet.Withdraw(50)
	history = wallet.GetTransactionHistory()
	if len(history) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(history))
	}
	if history[1].Type != Withdraw {
		t.Errorf("Expected transaction type WITHDRAW, got %s", history[1].Type)
	}
	if history[1].Amount != 50 {
		t.Errorf("Expected transaction amount 50, got %f", history[1].Amount)
	}

	// Test transfer transaction
	wm.CreateWallet("user2")
	wm.SendMoney("user1", "user2", 25)
	history = wallet.GetTransactionHistory()
	if len(history) != 3 {
		t.Errorf("Expected 3 transactions, got %d", len(history))
	}
	if history[2].Type != Transfer {
		t.Errorf("Expected transaction type TRANSFER, got %s", history[2].Type)
	}
	if history[2].Amount != 25 {
		t.Errorf("Expected transaction amount 25, got %f", history[2].Amount)
	}
}
