package main

import "fmt"

// OrderSystem представляет систему создания заказа.
type OrderSystem struct{}

func (o *OrderSystem) CreateOrder(item string) {
	fmt.Printf("Order for %s created.\n", item)
}

// PaymentSystem представляет платёжную систему.
type PaymentSystem struct{}

func (p *PaymentSystem) ProcessPayment(amount float64) {
	fmt.Printf("Payment of $%.2f processed.\n", amount)
}

// DeliverySystem представляет отдел доставки.
type DeliverySystem struct{}

func (d *DeliverySystem) ArrangeDelivery(address string) {
	fmt.Printf("Delivery arranged to %s.\n", address)
}

// CustomerSupportFacade предоставляет унифицированный интерфейс для взаимодействия с магазином.
type CustomerSupportFacade struct {
	orderSystem    *OrderSystem
	paymentSystem  *PaymentSystem
	deliverySystem *DeliverySystem
}

// PlaceOrder принимает заказ, обрабатывает платеж и организует доставку.
func (csf *CustomerSupportFacade) PlaceOrder(item string, amount float64, address string) {
	fmt.Println("Placing order...")
	csf.orderSystem.CreateOrder(item)
	csf.paymentSystem.ProcessPayment(amount)
	csf.deliverySystem.ArrangeDelivery(address)
	fmt.Println("Order placed successfully.")
}

// NewCustomerSupportFacade создает новый экземпляр CustomerSupportFacade.
func NewCustomerSupportFacade() *CustomerSupportFacade {
	return &CustomerSupportFacade{
		orderSystem:    &OrderSystem{},
		paymentSystem:  &PaymentSystem{},
		deliverySystem: &DeliverySystem{},
	}
}

func main() {
	// Создаем фасад службы поддержки клиентов.
	customerSupport := NewCustomerSupportFacade()

	// Используем фасад для размещения заказа.
	customerSupport.PlaceOrder("Laptop", 1299.99, "г.Москва, Лубянка д.4")
}
