package payment

import "errors"

// processPayment handles the actual payment processing
func processPayment(amount int) error {
    if amount <= 0 {
        return errors.New("invalid amount")
    }
    return nil
}
