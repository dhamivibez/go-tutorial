package main

import "fmt"

type Item struct {
	price float64
	stock int
}

func main() {

	var option int

	inventories := map[string]Item{
		"apple": {
			price: 4.99,
			stock: 30,
		},
		"banana": {

			price: 5.99,
			stock: 40,
		},
		"bread": {

			price: 6.99,
			stock: 50,
		},
	}

	for {

		fmt.Println("Inventory management App")
		fmt.Println("Choose option (1 - 4)")
		fmt.Println("1. List inventory")
		fmt.Println("2. Add inventory")
		fmt.Println("3. Update inventory")
		fmt.Println("4. Remove inventory")
		fmt.Println("5. Exit")

		fmt.Print("Choose option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			for name, inventory := range inventories {
				fmt.Printf("%s - Price: %.2f, Stock: %d\n", name, inventory.price, inventory.stock)
			}

		case 2:
			var name string
			var price float64
			var stock int

			fmt.Print("What inventory are you adding?: ")
			fmt.Scanln(&name)

			fmt.Printf("How much is %s? ", name)
			fmt.Scanln(&price)

			fmt.Printf("How many %s are available? ", name)
			fmt.Scanln(&stock)

			inventories[name] = Item{price, stock}

		case 3:
			var name string
			fmt.Print("Which inventory item do you want to update? ")
			fmt.Scanln(&name)

			item, exists := inventories[name]
			if !exists {
				fmt.Println("Item not found in inventory.")
				break
			}

			var newPrice float64
			var newStock int

			fmt.Printf("Enter new price for %s (current: %.2f): ", name, item.price)
			fmt.Scanln(&newPrice)

			fmt.Printf("Enter new stock for %s (current: %d): ", name, item.stock)
			fmt.Scanln(&newStock)

			item.price = newPrice
			item.stock = newStock

			inventories[name] = item

			fmt.Println("Item updated successfully!")
		case 4:
			var name string
			fmt.Print("What inventory do you want to delete? ")
			fmt.Scanln(&name)

			_, exists := inventories[name]

			if !exists {
				fmt.Println("Inventory not found")
				break
			}

			delete(inventories, name)
			fmt.Println("Inventory deleted successfully")
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
