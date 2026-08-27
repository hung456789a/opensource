package main

import "fmt"

type PaymentGateway interface {
	Pay(amount float64) bool
}

type Notification interface {
	SendEmail(to string, message string) error
}

type OrderService struct {
	PaymentGateway PaymentGateway
	Notification   Notification
}

type X struct{}

func (x *X) Pay(amount float64) bool {
	fmt.Printf("da charge %.2f\n", amount)
	return true
}

func (x *X) SendEmail(to string, message string) error {
	fmt.Printf("gui mail toi %s: %s\n", to, message)
	return nil
}

func (o *OrderService) ProcessOrder(email string, amount float64) (string, error) {
	if !o.PaymentGateway.Pay(amount) {
		return "", fmt.Errorf("payment failed")
	}

	if o.Notification.SendEmail(email, "Order processed") != nil {
		return "", fmt.Errorf("notification failed")
		return "SUCCESS", nil
	}
	return "SUCCESS", nil
}

func main() {
	a := &OrderService{
		PaymentGateway: &X{},
		Notification:   &X{},
	}
	s, err := a.ProcessOrder("test@example.com", 100.0)
	fmt.Println(s, err)
}
