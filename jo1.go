package main

import "fmt"

type IncomeSource struct {
    Category string
    Amount   float64
    Expenses float64
    Profit   float64
}

func (i *IncomeSource) CalculateProfit() {
    i.Profit = i.Amount - i.Expenses
}

func addIncomeSource(sources *[]IncomeSource) {
    var category string
    var amount, expenses float64

    fmt.Print("Enter the Category (e.g., freelancing, passive_income): ")
    fmt.Scanln(&category)

    fmt.Print("Enter the Amount (Income): ")
    fmt.Scanln(&amount)

    fmt.Print("Enter the Expenses: ")
    fmt.Scanln(&expenses)

    income := IncomeSource{Category: category, Amount: amount, Expenses: expenses}
    income.CalculateProfit()

    *sources = append(*sources, income)
    fmt.Println("Income source added successfully!\n")
}

func editIncomeSource(sources *[]IncomeSource) {
    var index int
    var category string
    var amount, expenses float64

    fmt.Print("Enter the index of the income source to edit: ")
    fmt.Scanln(&index)

    if index >= 0 && index < len(*sources) {
        fmt.Print("Enter the new Category: ")
        fmt.Scanln(&category)

        fmt.Print("Enter the new Amount (Income): ")
        fmt.Scanln(&amount)

        fmt.Print("Enter the new Expenses: ")
        fmt.Scanln(&expenses)

        (*sources)[index].Category = category
        (*sources)[index].Amount = amount
        (*sources)[index].Expenses = expenses
        (*sources)[index].CalculateProfit()

        fmt.Println("Income source updated successfully!\n")
    } else {
        fmt.Println("Invalid index!\n")
    }
}

func deleteIncomeSource(sources *[]IncomeSource) {
    var index int
    fmt.Print("Enter the index of the income source to delete: ")
    fmt.Scanln(&index)

    if index >= 0 && index < len(*sources) {
        *sources = append((*sources)[:index], (*sources)[index+1:]...)
        fmt.Println("Income source deleted successfully!\n")
    } else {
        fmt.Println("Invalid index!\n")
    }
}

func selectionSort(sources *[]IncomeSource) {
    for i := 0; i < len(*sources)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(*sources); j++ {
            if (*sources)[j].Amount < (*sources)[minIndex].Amount {
                minIndex = j
            }
        }
        (*sources)[i], (*sources)[minIndex] = (*sources)[minIndex], (*sources)[i]
    }
}

func generateReport(sources []IncomeSource) {
    totalIncome, totalProfit := 0.0, 0.0
    fmt.Println("\nMonthly Report:")
    for i, source := range sources {
        fmt.Printf("%d. Category: %s, Amount: %.2f, Expenses: %.2f, Profit: %.2f\n",
            i, source.Category, source.Amount, source.Expenses, source.Profit)
        totalIncome += source.Amount
        totalProfit += source.Profit
    }
    fmt.Printf("Total Income: %.2f\nTotal Profit: %.2f\n", totalIncome, totalProfit)
}

func calculateTotalProfit(sources []IncomeSource) {
    total := 0.0
    for _, source := range sources {
        total += source.Profit
    }
    fmt.Printf("Total Profit from All Sources: %.2f\n", total)
}

// show only passive income sources
func showPassiveIncome(sources []IncomeSource) {
    fmt.Println("\nPassive Income Sources:")
    found := false
    for i, source := range sources {
        if containsPassive(source.Category) {
            fmt.Printf("%d. Category: %s, Amount: %.2f, Expenses: %.2f, Profit: %.2f\n",
                i, source.Category, source.Amount, source.Expenses, source.Profit)
            found = true
        }
    }
    if !found {
        fmt.Println("No passive income sources found.")
    }
}

// check if "passive" is in category
func containsPassive(category string) bool {
    for i := 0; i <= len(category)-7; i++ {
        if category[i:i+7] == "passive" {
            return true
        }
    }
    return false
}

func main() {
    var sources []IncomeSource

    for {
        fmt.Println("\nIncome Tracker Application")
        fmt.Println("1. Add New Income Source")
        fmt.Println("2. Edit Income Source")
        fmt.Println("3. Delete Income Source")
        fmt.Println("4. Generate Monthly Report")
        fmt.Println("5. Exit")
        fmt.Println("6. Calculate Total Profit Only")
        fmt.Println("7. Show Passive Income Sources")
        fmt.Print("Choose an option: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            addIncomeSource(&sources)
        case 2:
            editIncomeSource(&sources)
        case 3:
            deleteIncomeSource(&sources)
        case 4:
            generateReport(sources)
        case 5:
            fmt.Println("Exiting the program.")
            return
        case 6:
            calculateTotalProfit(sources)
        default:
            fmt.Println("Invalid option. Please choose again.")
        }
    }
}
