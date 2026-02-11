package payment

import "errors"

// processPayment handles the actual payment processing
func processPayment(amount int) error {
    if amount <= 0 {
        return errors.New("invalid amount")
    }
    return nil
}

// ProcessCheckout handles checkout payments with unified auth
func ProcessCheckout(token string, amount int) error {
    // TODO: call auth.ValidateCheckoutToken(token)
    return processPayment(amount)
}
