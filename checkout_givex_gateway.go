package adyen

import "github.com/davecgh/go-spew/spew"

const (
	adyenGivexURL = "authorise"
)

// CheckoutGivexGateway - allows you to accept Givex gift card payment using Adyen's
// methods and flows.
type CheckoutGivexGateway struct {
	*Adyen

	version string
}

// GivexCheckout - returns CheckoutGivexGateway
func (a *Adyen) GivexCheckout() *CheckoutGivexGateway {
	return &CheckoutGivexGateway{a, CheckoutAPIVersion}
}

// GivexRedem - process a gift card payment using Adyen API.
func (g *CheckoutGivexGateway) GivexRedem(req *GivexRedem) (*PaymentMethodsResponse, error) {

	url := g.checkoutURL(adyenGivexURL, g.version)

	resp, err := g.executeApiKey(url, req)
	if err != nil {
		return nil, err
	}

	spew.Dump(resp)

	//return resp.paymentMethods()

	return nil, nil
}

// GivexBalance - process a gift card balance request using Adyen API.
func (g *CheckoutGivexGateway) GivexBalance(req *GivexRedem) (*GivexBalanceResponse, error) {

	url := g.checkoutURL(adyenGivexURL, g.version)

	resp, err := g.executeApiKey(url, req)
	if err != nil {
		return nil, err
	}

	spew.Dump(resp)

	return resp.givexBalance()
}
