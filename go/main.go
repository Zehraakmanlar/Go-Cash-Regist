package main

import "fmt"

type Item struct {
	Name     string
	Price    int
	Discount int
}
type Describable interface {
	Description() string
}

func (item Item) Description() string {
	var description string
	if item.Discount > 0 {
		description = fmt.Sprintf("%s - %d (İndirimli fiyat:%d)", item.Name, item.Price, item.Price-item.Discount)
	} else {
		description = fmt.Sprintf("%s - %d", item.Name, item.Price)
	}
	return description
}

func calculatePrice(item Item) float64 {
	if item.Discount > 0 {
		lastPrice := float64(item.Price - item.Discount)
		return lastPrice
	}
	return float64(item.Price) // indirim yok ise
}

func totalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += float64(item.Price - item.Discount)
	}
	return total
}

func main() {

	apple := Item{Name: "Elma", Price: 30, Discount: 10}
	orange := Item{Name: "Portakal", Price: 1.00, Discount: 0}

	items := []Describable{apple, orange}

	for _, item := range items {
		fmt.Println(item.Description())
	}

	total := totalPrice([]Item{apple, orange})
	fmt.Printf("Toplam Fiyat: %.2f TL\n", total)
	return
}
