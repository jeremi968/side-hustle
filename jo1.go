package main

import "fmt"

const NMAX int = 99

type IncomeSource struct {
	Category string
	Amount   float64
	Expenses float64
	Profit   float64
}

type tabSource [NMAX]IncomeSource

func calculateProfit(source *IncomeSource) {
	source.Profit = source.Amount - source.Expenses
}

func addIncome(data *tabSource, total *int) {
	if *total >= NMAX {
		fmt.Println("Data is full!")
		return
	}

	var cat string
	var amt, exp float64

	fmt.Print("Enter category: ")
	fmt.Scan(&cat)

	fmt.Print("Enter income amount: ")
	fmt.Scan(&amt)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&exp)

	data[*total].Category = cat
	data[*total].Amount = amt
	data[*total].Expenses = exp
	data[*total].Profit = amt - exp

	*total = *total + 1
	fmt.Println("Income added successfully!\n")
}

func editIncome(data *tabSource, total int) {
	var index int
	fmt.Print("Enter index to edit: ")
	fmt.Scan(&index)

	if index >= 0 && index < total {
		var cat string
		var amt, exp float64

		fmt.Print("Enter new category: ")
		fmt.Scan(&cat)

		fmt.Print("Enter new income amount: ")
		fmt.Scan(&amt)

		fmt.Print("Enter new expenses: ")
		fmt.Scan(&exp)

		data[index].Category = cat
		data[index].Amount = amt
		data[index].Expenses = exp
		data[index].Profit = amt - exp

		fmt.Println("Income updated successfully!\n")
	} else {
		fmt.Println("Invalid index!\n")
	}
}

func deleteIncome(data *tabSource, total *int) {
	var index int
	fmt.Print("Enter index to delete: ")
	fmt.Scan(&index)

	if index >= 0 && index < *total {
		for i := index; i < *total-1; i++ {
			data[i] = data[i+1]
		}
		*total = *total - 1
		fmt.Println("Income deleted successfully!\n")
	} else {
		fmt.Println("Invalid index!\n")
	}
}

func generateReport(data tabSource, total int) {
	var totalIncome float64 = 0
	var totalProfit float64 = 0

	fmt.Println("\n--- Monthly Income Report ---")
	for i := 0; i < total; i++ {
		fmt.Printf("%d. Category: %s | Income: %.2f | Expenses: %.2f | Profit: %.2f\n",
			i, data[i].Category, data[i].Amount, data[i].Expenses, data[i].Profit)
		totalIncome += data[i].Amount
		totalProfit += data[i].Profit
	}

	fmt.Printf("\nTotal Income: %.2f\n", totalIncome)
	fmt.Printf("Total Profit: %.2f\n", totalProfit)
}

func calculateTotalProfit(data tabSource, total int) {
	var totalProfit float64 = 0
	for i := 0; i < total; i++ {
		totalProfit += data[i].Profit
	}
	fmt.Printf("Total Profit: %.2f\n", totalProfit)
}

func main() {
	var data tabSource
	var total int = 0

	for {
		fmt.Println("\n=== Passive Income Tracker Menu ===")
		fmt.Println("1. Add Income Source")
		fmt.Println("2. Edit Income Source")
		fmt.Println("3. Delete Income Source")
		fmt.Println("4. Show Monthly Report")
		fmt.Println("5. Calculate Total Profit")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option (1-6): ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			addIncome(&data, &total)
		} else if choice == 2 {
			editIncome(&data, total)
		} else if choice == 3 {
			deleteIncome(&data, &total)
		} else if choice == 4 {
			generateReport(data, total)
		} else if choice == 5 {
			calculateTotalProfit(data, total)
		} else if choice == 6 {
			fmt.Println("Exiting program. Goodbye!")
			break
		} else {
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}
