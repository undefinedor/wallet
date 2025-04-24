package main

import (
	"fmt"
	"log"
	"time"

	"github.com/undefinedor/wallet/wallet"
)

func main() {
	// Create a new wallet manager
	wm := wallet.NewWalletManager()

	// Create wallets for two users
	err := wm.CreateWallet("user1")
	if err != nil {
		log.Fatalf("Failed to create wallet for user1: %v", err)
	}

	err = wm.CreateWallet("user2")
	if err != nil {
		log.Fatalf("Failed to create wallet for user2: %v", err)
	}

	// Get user1's wallet
	wallet1, err := wm.GetWallet("user1")
	if err != nil {
		log.Fatalf("Failed to get wallet for user1: %v", err)
	}

	// Deposit money into user1's wallet
	err = wallet1.Deposit(1000)
	if err != nil {
		log.Fatalf("Failed to deposit money: %v", err)
	}

	// Check user1's balance
	fmt.Printf("User1's balance: $%.2f\n", wallet1.GetBalance())

	// Send money from user1 to user2
	err = wm.SendMoney("user1", "user2", 500)
	if err != nil {
		log.Fatalf("Failed to send money: %v", err)
	}

	// Get user2's wallet
	wallet2, err := wm.GetWallet("user2")
	if err != nil {
		log.Fatalf("Failed to get wallet for user2: %v", err)
	}

	// Check both users' balances
	fmt.Printf("User1's balance after transfer: $%.2f\n", wallet1.GetBalance())
	fmt.Printf("User2's balance after transfer: $%.2f\n", wallet2.GetBalance())

	// Withdraw money from user2's wallet
	err = wallet2.Withdraw(200)
	if err != nil {
		log.Fatalf("Failed to withdraw money: %v", err)
	}

	// Check user2's final balance
	fmt.Printf("User2's final balance: $%.2f\n", wallet2.GetBalance())

	// Print transaction history for both users
	fmt.Println("\nTransaction History for User1:")
	printTransactionHistory(wallet1.GetTransactionHistory())

	fmt.Println("\nTransaction History for User2:")
	printTransactionHistory(wallet2.GetTransactionHistory())
}

func printTransactionHistory(transactions []wallet.Transaction) {
	for _, t := range transactions {
		fmt.Printf("Type: %s, Amount: $%.2f, Time: %s\n", t.Type, t.Amount, t.Timestamp.Format(time.RFC3339))
		if t.Type == wallet.Transfer {
			fmt.Printf("  From: %s, To: %s\n", t.From, t.To)
		}
	}
}
