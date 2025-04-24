# Wallet Application

A simple centralized wallet application that allows users to manage their money through various operations.

## Features

- User wallet management
- Deposit functionality
- Withdrawal functionality
- Inter-user transfers
- Balance checking
- Transaction history tracking

## Design Decisions

1. **Language Choice**: Go was chosen for its:
   - Strong concurrency support
   - Built-in testing framework
   - Static typing and compile-time checks
   - Excellent performance
   - Simple and clean syntax

2. **Architecture Design**:
   - Clear separation of concerns
   - Interface-based design for better testability
   - Immutable transaction records
   - Thread-safe operations

3. **Data Structures**:
   - Map-based user and wallet storage for O(1) lookups
   - Slice-based transaction history for ordered records
   - float64 type for amount handling

## Requirements

- Go 1.16 or later
- Go environment with module support

## Installation and Usage

### As a Library

To use this wallet implementation in your project:

1. Add the module to your project:
   ```bash
   go get github.com/undefinedor/wallet
   ```

2. Import the wallet package in your code:
   ```go
   import "github.com/undefinedor/wallet/wallet"
   ```

3. Use the wallet functionality:
   ```go
   // Create a new wallet manager
   wm := wallet.NewWalletManager()
   
   // Create a wallet
   err := wm.CreateWallet("user1")
   if err != nil {
       log.Fatal(err)
   }
   
   // Get the wallet
   wallet, err := wm.GetWallet("user1")
   if err != nil {
       log.Fatal(err)
   }
   
   // Deposit money
   err = wallet.Deposit(1000)
   if err != nil {
       log.Fatal(err)
   }
   
   // Check balance
   balance := wallet.GetBalance()
   fmt.Printf("Balance: $%.2f\n", balance)
   
   // Get transaction history
   history := wallet.GetTransactionHistory()
   for _, t := range history {
       fmt.Printf("Type: %s, Amount: $%.2f, Time: %s\n", 
           t.Type, t.Amount, t.Timestamp.Format(time.RFC3339))
   }
   ```

### For Development

1. Clone the repository:
   ```bash
   git clone https://github.com/undefinedor/wallet.git
   cd wallet
   ```

2. Run tests:
   ```bash
   go test ./...
   ```

3. Run the example program:
   ```bash
   go run examples/basic.go
   ```

## Code Structure

- `wallet/` - Core wallet implementation
  - `wallet.go` - Main wallet interface and implementation
  - `wallet_test.go` - Test cases
- `examples/` - Example programs
  - `basic.go` - Basic usage example
- `go.mod` - Module definition and dependencies

## Main Features

1. **Create Wallet**
   ```go
   wm := wallet.NewWalletManager()
   err := wm.CreateWallet("user1")
   ```

2. **Deposit**
   ```go
   wallet, _ := wm.GetWallet("user1")
   err := wallet.Deposit(100)
   ```

3. **Withdraw**
   ```go
   err := wallet.Withdraw(50)
   ```

4. **Transfer**
   ```go
   err := wm.SendMoney("user1", "user2", 50)
   ```

5. **Check Balance**
   ```go
   balance := wallet.GetBalance()
   ```

6. **Get Transaction History**
   ```go
   history := wallet.GetTransactionHistory()
   for _, t := range history {
       fmt.Printf("Type: %s, Amount: $%.2f, Time: %s\n", 
           t.Type, t.Amount, t.Timestamp.Format(time.RFC3339))
   }
   ```

## Areas for Improvement

1. Add data persistence layer (database)
2. Implement user authentication and authorization
3. Add API endpoints
4. Implement rate limiting
5. Add transaction rollback mechanism
6. Enhance error handling
7. Add logging and monitoring
8. Support multiple currencies

## Development Time

- Planning: 20 minutes
- Implementation: 1 hours
- Testing: 20 minutes
- Documentation: 20 minutes
Total: ~2 hours

## Unimplemented Features

1. Database persistence
2. User authentication
3. API endpoints
4. Currency conversion
5. Transaction fees
6. Account types (savings, checking, etc.)

## Engineering Best Practices

1. **Code Organization**:
   - Clear package structure
   - Separation of concerns
   - Interface-based design

2. **Testing**:
   - Unit tests for all core functionality
   - Edge case testing
   - Mock interface testing

3. **Error Handling**:
   - Custom error types
   - Proper error propagation
   - Input validation

4. **Concurrency**:
   - Thread-safe operations
   - Mutex protection for shared resources

5. **Documentation**:
   - Clear code comments
   - Comprehensive README
   - Usage examples

## Notes

1. All amount operations use float64 type, in production environment it's recommended to use a dedicated currency handling library
2. Current implementation is in-memory, data will be lost after restart
3. No user authentication implemented, needs to be added for production use
4. No transaction mechanism implemented, needs to be added based on requirements

## Contributing

1. Fork the project
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License 