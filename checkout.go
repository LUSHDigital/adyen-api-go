package adyen

// PaymentMethods contains the fields required by the checkout
// API's /paymentMethods endpoint.  See the following for more
// information:
//
// https://docs.adyen.com/api-explorer/#/PaymentSetupAndVerificationService/v32/paymentMethods
type PaymentMethods struct {
	Amount           *Amount `json:"amount"`
	Channel          string  `json:"channel"`
	CountryCode      string  `json:"countryCode"`
	MerchantAccount  string  `json:"merchantAccount"`
	ShopperLocale    string  `json:"shopperLocale"`
	ShopperReference string  `json:"shopperReference"`
}

// PaymentMethodsResponse is returned by Adyen in response to
// a PaymentMethods request.
type PaymentMethodsResponse struct {
	PaymentMethods         []PaymentMethodDetails         `json:"paymentMethods"`
	OneClickPaymentMethods []OneClickPaymentMethodDetails `json:"oneClickPaymentMethods,omitempty"`
}

// PaymentMethodDetails describes the PaymentMethods part of
// a PaymentMethodsResponse.
type PaymentMethodDetails struct {
	Details []PaymentMethodDetailsInfo `json:"details,omitempty"`
	Name    string                     `json:"name"`
	Type    string                     `json:"type"`
}

// PaymentMethodDetailsInfo describes the collection of all
// payment methods.
type PaymentMethodDetailsInfo struct {
	Items []PaymentMethodItems `json:"items"`
	Key   string               `json:"key"`
	Type  string               `json:"type"`
}

// PaymentMethodItems describes a single payment method.
type PaymentMethodItems struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OneClickPaymentMethodDetails describes the OneClickPayment part of
// a PaymentMethods response.
type OneClickPaymentMethodDetails struct {
	Details       []PaymentMethodTypes       `json:"details"`
	Name          string                     `json:"name"`
	Type          string                     `json:"type"`
	StoredDetails PaymentMethodStoredDetails `json:"storedDetails"`
}

// PaymentMethodTypes describes any additional information associated
// with a OneClick payment.
type PaymentMethodTypes struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// PaymentMethodStoredDetails describes the information stored for a
// OneClick payment.
type PaymentMethodStoredDetails struct {
	Card PaymentMethodCard `json:"card"`
}

// PaymentMethodCard describes the card information associated with a
// OneClick payment.
type PaymentMethodCard struct {
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	HolderName  string `json:"holderName"`
	Number      string `json:"number"`
}

// Payment will create a payment within the checkout service
type Payment struct {
	AdditionalData     *AdditionalData    `json:"additionalData,omitempty"`
	Amount             *Amount            `json:"amount" valid:"required"`
	BillingAddress     *Address           `json:"billingAddress,omitempty"`
	DeliveryAddress    *Address           `json:"deliveryAddress,omitempty"`
	Reference          string             `json:"reference"`
	MerchantAccount    string             `json:"merchantAccount" valid:"required"`
	ReturnURL          string             `json:"returnUrl" valid:"required"`
	ShopperReference   string             `json:"shopperReference,omitempty"`
	Recurring          *Recurring         `json:"recurring,omitempty"`
	ShopperInteraction string             `json:"shopperInteraction,omitempty"`
	BrowserInfo        *BrowserInfo       `json:"browserInfo,omitempty"`
	PaymentMethod      *SecuredFieldsCard `json:"paymentMethod" valid:"required"`
	ShopperEmail       string             `json:"shopperEmail,omitempty"`
	ShopperIP          string             `json:"shopperIP,omitempty"`
	ShopperLocale      string             `json:"shopperLocale,omitempty"`
	Channel            string             `json:"channel"`
}

// PaymentResponse
type PaymentResponse struct {
	ResultCode        string                            `json:"resultCode"`
	PaymentData       string                            `json:"paymentData"`
	Redirect          *PaymentRedirectResponse          `json:"redirect"`
	Details           *[]PaymentRedirectDetailsResponse `json:"details"`
	RefusalReason     string                            `json:"refusalReason"`
	RefusalReasonCode string                            `json:"refusalReasonCode"`
	PSPReference      string                            `json:"pspReference"`
	// TODO: Implement the rest of the payment response parameters
}

// PaymentRedirectResponse
type PaymentRedirectResponse struct {
	Data   *PaymentRedirectDataResponse `json:"data"`
	Method string                       `json:"method"`
	URL    string                       `json:"url"`
}

// PaymentRedirectDataResponse
type PaymentRedirectDataResponse struct {
	PaReq   string `json:"PaReq"`
	TermUrl string `json:"TermUrl"`
	MD      string `json:"MD"`
}

// PaymentRedirectDetailsResponse
type PaymentRedirectDetailsResponse struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// PaymentFraudResultResponse
type PaymentFraudResultResponse struct {
	AccountScore string                             `json:"accountScore"`
	Results      *PaymentFraudResultResultsResponse `json:"results"`
}

// PaymentFraudResultResultsResponse
type PaymentFraudResultResultsResponse struct {
	AccountScore int    `json:"accountScore"`
	CheckId      int    `json:"checkId"`
	Name         string `json:"name"`
}

// PaymentDetails
type PaymentDetails struct {
	PaymentData string      `json:"paymentData"`
	Details     interface{} `json:"details"` // TODO: Discover the type of the input details. This will be included in PaymentResponse
}

// PaymentDetailsResponse will give you the result of submitting details on a payment
// TODO: Extend this fore all result data
type PaymentDetailsResponse struct {
	Status    int    `json:"status"`
	ErrorCode string `json:"errorCode"`
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`

	ResultCode string `json:"resultCode"`
}
