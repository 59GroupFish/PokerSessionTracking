package internal

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"sort"

// 	"golang.org/x/oauth2/google"
// 	"google.golang.org/api/option"
// 	"google.golang.org/api/sheets/v4"
// )

// // simplifyDebts calculates the most simplified way to settle debts within a group.
// // It aims to minimize the number of individual payments.
// //
// // The function takes a slice of 'Person' structs as input.
// // Each Person struct has:
// //   - Name (string): The name of the person.
// //   - OwedOwing (float64): The amount of money.
// //   - If negative, the person owes this amount (e.g., -100 means they owe 100).
// //   - If positive, the person is owed this amount (e.g., 120 means they are owed 120).
// //
// // The function returns a slice of 'Payment' structs.
// // Each Payment struct has:
// //   - Payer (string): The name of the person making the payment.
// //   - Payee (string): The name of the person receiving the payment.
// //   - Amount (float64): The amount of money transferred in this payment.
// func simplifyDebts(session Session) []Payment {
// 	// Create mutable copies of people to update their balances during the settlement process.
// 	// This avoids modifying the original input slice.
// 	var owes []Person // Slice to store people who owe money (negative OwedOwing)
// 	var owed []Person // Slice to store people who are owed money (positive OwedOwing)

// 	// Separate people into two groups: those who owe and those who are owed.
// 	// People with an OwedOwing of 0 are already settled and are ignored.
// 	for _, p := range people {
// 		if p.OwedOwing < 0 {
// 			owes = append(owes, p)
// 		} else if p.OwedOwing > 0 {
// 			owed = append(owed, p)
// 		}
// 	}

// 	// Sort the 'owes' slice:
// 	// We sort by OwedOwing in ascending order. This means the person who owes the most
// 	// (i.e., has the largest absolute negative value, e.g., -100 before -50) comes first.
// 	sort.Slice(owes, func(i, j int) bool {
// 		return owes[i].OwedOwing < owes[j].OwedOwing
// 	})

// 	// Sort the 'owed' slice:
// 	// We sort by OwedOwing in descending order. This means the person who is owed the most
// 	// (i.e., has the largest positive value, e.g., 120 before 30) comes first.
// 	sort.Slice(owed, func(i, j int) bool {
// 		return owed[i].OwedOwing > owed[j].OwedOwing
// 	})

// 	var payments []Payment // Slice to store the simplified payment transactions.
// 	debtorIdx := 0         // Index for the current debtor in the 'owes' slice.
// 	creditorIdx := 0       // Index for the current creditor in the 'owed' slice.

// 	// Define a small epsilon value for floating-point comparisons.
// 	// Due to the nature of float arithmetic, direct comparison with 0 can be unreliable.
// 	const epsilon = 1e-9

// 	// Iterate as long as there are still debtors and creditors to settle.
// 	for debtorIdx < len(owes) && creditorIdx < len(owed) {
// 		// Get pointers to the current debtor and creditor.
// 		// Using pointers allows us to directly modify their OwedOwing balances in the slices.
// 		debtor := &owes[debtorIdx]
// 		creditor := &owed[creditorIdx]

// 		// Determine the amount to be paid in this specific transaction.
// 		// It's the minimum of the absolute amount the debtor owes and the amount the creditor is owed.
// 		// We use math.Abs() because debtor.OwedOwing is negative.
// 		paymentAmount := math.Min(math.Abs(debtor.OwedOwing), creditor.OwedOwing)

// 		// Record this payment transaction.
// 		payments = append(payments, Payment{
// 			Payer:  debtor.Name,
// 			Payee:  creditor.Name,
// 			Amount: paymentAmount,
// 		})

// 		// Update the balances of the debtor and creditor.
// 		// The debtor's negative balance increases (moves towards zero).
// 		debtor.OwedOwing += paymentAmount
// 		// The creditor's positive balance decreases (moves towards zero).
// 		creditor.OwedOwing -= paymentAmount

// 		// Check if the current debtor's balance is settled (close to zero).
// 		// If so, move to the next debtor.
// 		if math.Abs(debtor.OwedOwing) < epsilon {
// 			debtorIdx++
// 		}

// 		// Check if the current creditor's balance is settled (close to zero).
// 		// If so, move to the next creditor.
// 		if math.Abs(creditor.OwedOwing) < epsilon {
// 			creditorIdx++
// 		}
// 	}

// 	return payments
// }
