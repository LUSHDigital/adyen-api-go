package adyen

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

const (
	paymentMethodsURL = "paymentMethods"
	paymentURL        = "payments"
	paymentDetailsURL = "payments/details"
)

// CheckoutGateway - allows you to accept all of Adyen's payment
// methods and flows.
type CheckoutGateway struct {
	*Adyen

	version string
}

// Checkout - returns CheckoutGateway
func (a *Adyen) Checkout() *CheckoutGateway {
	return &CheckoutGateway{a, CheckoutAPIVersion}
}

// PaymentMethods - Perform paymentMethods request in Adyen.
//
// Used to get a collection of available payment methods for a merchant.
func (g *CheckoutGateway) PaymentMethods(req *PaymentMethods) (*PaymentMethodsResponse, error) {
	url := g.checkoutURL(paymentMethodsURL, g.version)

	resp, err := g.execute(url, req)
	if err != nil {
		return nil, err
	}

	return resp.paymentMethods()
}

// Payment creates a payment in the Adyen Checkout Service.
func (g *CheckoutGateway) Payment(req *Payment) (*PaymentResponse, error) {
	url := g.checkoutURL(paymentURL, g.version)
	log.Println("ADYEN URL:")
	spew.Dump(url)
	log.Println("TO ADYEN REQ:")
	spew.Dump(req)
	resp, err := g.execute(url, req)
	log.Println("ADYEN RESP:")
	spew.Dump(resp)
	if err != nil {
		return nil, err
	}

	return resp.payment()
}

// PaymentDetails submits details for a created payment, such as verifying 3D Secure
func (g *CheckoutGateway) PaymentDetails(req *PaymentDetails) (*PaymentDetailsResponse, error) {
	url := g.checkoutURL(paymentDetailsURL, g.version)

	resp, err := g.execute(url, req)
	if err != nil {
		return nil, err
	}

	return resp.paymentDetails()
}
