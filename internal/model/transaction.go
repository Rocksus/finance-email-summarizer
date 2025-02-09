package model

// Transaction is the base model, keeps track about transaction information,
// amount billed, and many other things.
type Transaction struct {
}

// Subscription tracks active and inactive subscriptions
// when it was started, when it was stopped,
// and when it should be inserted as a transaction.
type Subscription struct {
}

// Installment tracks active installment.
// It will deduct the full amount from your credit account balance,
// while splitting it across installment date.
// Installment follows your credit account.
type Installment struct {
}
