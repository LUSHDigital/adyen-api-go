package adyen

// GivexRedem contains the fields required by the Adyen Givex payment
// API's //authorise endpoint.  See the following for more
// information:
//
// https://docs.adyen.com/developers/payment-methods/gift-cards-givex-and-svs#makepayment
type GivexRedem struct {
	Amount          *GivexAmount `json:"amount" valid:"required"`
	Card            *GivexCard   `json:"card" valid:"required"`
	Reference       string       `json:"reference" valid:"required"`
	MerchantAccount string       `json:"merchantAccount" valid:"required"`
	SelectedBrand   string       `json:"selectedBrand" valid:"required"`
}

type GivexAmount struct {
	Value    int    `json:"value" valid:"required"`
	Currency string `json:"currency" valid:"currency,required"`
}

type GivexCard struct {
	Cvc         string `json:"cvc"`
	ExpiryMonth string `json:"expiryMonth" valid:"required"`
	ExpiryYear  string `json:"expiryYear" valid:"required"`
	Number      string `json:"number" valid:"required"`
	HolderName  string `json:"holderName" valid:"required"`
}

type GivexBalanceResponse struct {
	AdditionalData *GivexAdditionalData `json:"additionalData"`
	PspReference   string               `json:"pspReference"`
	ResultCode     string               `json:"resultCode"`
}

type GivexAdditionalData struct {
	CurrentBalanceValue string `json:"currentBalanceValue"`
}
