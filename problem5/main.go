package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Address string
}

type SmsNotifier struct {
	Phone string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s : %s \n", e.Address, message)
	return nil
}

func (sms SmsNotifier) Send(message string) error {
	fmt.Printf("Sending sms to %s : %s \n", sms.Phone, message)
	return nil
}

func Notify(m Notifier, message string) error {
	return m.Send(message)
}

func main() {
	notifiers := []Notifier{
		EmailNotifier{Address: "brian@gmail.com"},
		SmsNotifier{Phone: "95922342341"},
	}

	for _, v := range notifiers {
		Notify(v, "Your server is down")
	}
}
