package adyen

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestPaymentMethods - test for https://docs.adyen.com/developers/checkout/api-integration
//
// This test requires CheckoutAPI access.  To obtain, visit https://docs.adyen.com/developers/user-management/how-to-get-the-checkout-api-key.
func TestPaymentMethods(t *testing.T) {
	t.Parallel()

	// Mock the Adyen checkout paymentMethods endpoint
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "paymentMethods":[
    {
      "details": null,
      "name":"Credit Card",
      "type":"scheme"
    }
  ]
}`))
	}))

	instance := getHTTPMockInstance(s.URL)

	request := &PaymentMethods{
		MerchantAccount: os.Getenv("ADYEN_ACCOUNT"),
	}

	resp, err := instance.Checkout().PaymentMethods(request)
	knownError, ok := err.(APIError)
	if ok {
		t.Errorf("Response should be succesfull. Known API Error: Code - %s, Message - %s, Type - %s", knownError.ErrorCode, knownError.Message, knownError.ErrorType)
	}

	if err != nil {
		t.Errorf("Response should be succesfull, error - %s", err.Error())
	}

	if resp.PaymentMethods[0].Name != "Credit Card" {
		t.Errorf("expected %s, received %s", "Credit Card", resp.PaymentMethods[0].Name)
	}
	if resp.PaymentMethods[0].Type != "scheme" {
		t.Errorf("expected %s, received %s", "scheme", resp.PaymentMethods[0].Type)
	}
}
