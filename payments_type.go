package paypal

import (
	"fmt"
	"net/url"
)

type PayerInfo struct {
	Email           string           `json:"email"`
	FirstName       string           `json:"first_name"`
	LastName        string           `json:"last_name"`
	PayerId         string           `json:"payer_id"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type PaymentMethod string

const (
	K_PAYMENT_METHOD_PAYPAL      PaymentMethod = "paypal"
	K_PAYMENT_METHOD_CREDIT_CARD PaymentMethod = "credit_card"
)

type Payer struct {
	PaymentMethod PaymentMethod `json:"payment_method"`
	Status        string        `json:"status,omitempty"`
	PayerInfo     *PayerInfo    `json:"payer_info,omitempty"`
}

type AmountDetails struct {
	Subtotal         string `json:"subtotal,omitempty"`
	Shipping         string `json:"shipping,omitempty"`
	Tax              string `json:"tax,omitempty"`
	HandlingFee      string `json:"handling_fee,omitempty"`
	ShippingDiscount string `json:"shipping_discount,omitempty"`
	Insurance        string `json:"insurance,omitempty"`
	GiftWrap         string `json:"gift_wrap,omitempty"`
}

type Amount struct {
	Total    string         `json:"total,omitempty"`
	Currency string         `json:"currency,omitempty"`
	Details  *AmountDetails `json:"details,omitempty"`
}

type PaymentOptions struct {
	AllowedPaymentMethod string `json:"allowed_payment_method"`
}

type Item struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Quantity    interface{} `json:"quantity"` // string or int
	Price       string      `json:"price"`
	Tax         string      `json:"tax"`
	SKU         string      `json:"sku"`
	Currency    string      `json:"currency"`
}

type ShippingAddress struct {
	RecipientName string `json:"recipient_name"`
	Line1         string `json:"line1"`
	Line2         string `json:"line2"`
	City          string `json:"city"`
	CountryCode   string `json:"country_code"`
	PostalCode    string `json:"postal_code"`
	Phone         string `json:"phone"`
	State         string `json:"state"`
}

type ItemList struct {
	Items           []*Item          `json:"items,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type SaleState string

const (
	K_SALE_STATE_COMPLETED          SaleState = "completed"
	K_SALE_STATE_PARTIALLY_REFUNDED SaleState = "partially_refunded"
	K_SALE_STATE_PENDING            SaleState = "pending"
	K_SALE_STATE_REFUNDED           SaleState = "refunded"
	K_SALE_STATE_DENIED             SaleState = "denied"
)

type Sale struct {
	Id                        string              `json:"id,omitempty"`
	PurchaseUnitReferenceId   string              `json:"purchase_unit_reference_id,omitempty"`
	Amount                    *Amount             `json:"amount,omitempty"`
	PaymentMode               string              `json:"payment_mode,omitempty"`
	State                     SaleState           `json:"state,omitempty"`
	ReasonCode                string              `json:"reason_code,omitempty"`
	ProtectionEligibility     string              `json:"protection_eligibility,omitempty"`
	ProtectionEligibilityType string              `json:"protection_eligibility_type,omitempty"`
	ClearingTime              string              `json:"clearing_time,omitempty"`
	PaymentHoldStatus         string              `json:"payment_hold_status,omitempty"`
	PaymentHoldReasons        []PaymentHoldReason `json:"payment_hold_reasons,omitempty"`
	TransactionFee            *Currency           `json:"transaction_fee,omitempty"`
	ReceivableAmount          *Currency           `json:"receivable_amount,omitempty"`
	ExchangeRate              string              `json:"exchange_rate,omitempty"`
	FMFDetails                *FMFDetails         `json:"fmf_details,omitempty"`
	ReceiptId                 string              `json:"receipt_id,omitempty"`
	ParentPayment             string              `json:"parent_payment,omitempty"`
	ProcessorResponse         *ProcessorResponse  `json:"processor_response,omitempty"`
	BillingAgreementId        string              `json:"billing_agreement_id,omitempty"`
	CreateTime                string              `json:"create_time,omitempty"`
	UpdateTime                string              `json:"update_time,omitempty"`
	Links                     []*Link             `json:"links,omitempty,omitempty"`
	InvoiceNumber             string              `json:"invoice_number,omitempty"`
	Custom                    string              `json:"custom,omitempty"`
	SoftDescriptor            string              `json:"soft_descriptor,omitempty"`
}

type PaymentHoldReason struct {
	PaymentHoldReason string `json:"payment_hold_reason,omitempty"`
}

type FMFDetails struct {
	FilterType  string `json:"filter_type,omitempty"`
	FilterId    string `json:"filter_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProcessorResponse struct {
	ResponseCode string `json:"response_code,omitempty"`
	AVSCode      string `json:"avs_code,omitempty"`
	CVVCode      string `json:"cvv_code,omitempty"`
	AdviceCode   string `json:"advice_code,omitempty"`
	ECISubmitted string `json:"eci_submitted,omitempty"`
	Vpas         string `json:"vpas,omitempty"`
}

type Refund struct {
	Id            string  `json:"id"`
	CreateTime    string  `json:"create_time"`
	UpdateTime    string  `json:"update_time"`
	State         string  `json:"state"`
	Amount        *Amount `json:"amount"`
	SaleId        string  `json:"sale_id"`
	ParentPayment string  `json:"parent_payment"`
	InvoiceNumber string  `json:"invoice_number"`
	Links         []*Link `json:"links,omitempty"`
}

type RelatedResources struct {
	Sale *Sale `json:"sale,omitempty"`
}

type Transaction struct {
	Amount         *Amount         `json:"amount"`
	Description    string          `json:"description,omitempty"`
	Custom         string          `json:"custom,omitempty"`
	InvoiceNumber  string          `json:"invoice_number,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	SoftDescriptor string          `json:"soft_descriptor,omitempty"`
	ItemList       *ItemList       `json:"item_list,omitempty"`

	// 返回结果添加的字段
	RelatedResources []*RelatedResources `json:"related_resources,omitempty"`
}

type Payee struct {
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
}

type RedirectURLs struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type PaymentIntent string

const (
	K_PAYMENT_INTENT_SALE      PaymentIntent = "sale"
	K_PAYMENT_INTENT_AUTHORIZE PaymentIntent = "authorize"
	K_PAYMENT_INTENT_ORDER     PaymentIntent = "order"
)

type PaymentState string

const (
	K_PAYMENT_STATE_CREATED  PaymentState = "created"
	K_PAYMENT_STATE_APPROVED PaymentState = "approved"
	K_PAYMENT_STATE_FAILED   PaymentState = "failed"
)

type Payment struct {
	Intent              PaymentIntent  `json:"intent"`
	ExperienceProfileId string         `json:"experience_profile_id,omitempty"`
	Payer               *Payer         `json:"payer"`
	Transactions        []*Transaction `json:"transactions"`
	NoteToPayer         string         `json:"note_to_payer,omitempty"`
	RedirectURLs        *RedirectURLs  `json:"redirect_urls"`

	// 返回结果添加的字段
	Id            string       `json:"id,omitempty"`
	CreateTime    string       `json:"create_time,omitempty"`
	State         PaymentState `json:"state,omitempty"`
	FailureReason string       `json:"failure_reason,omitempty"`
	UpdateTime    string       `json:"update_time,omitempty"`
	Links         []*Link      `json:"links,omitempty"`
}

type PaymentListParam struct {
	Count      int
	StartId    string
	StartIndex int
	StartTime  string
	EndTime    string
	SortBy     string
	SortOrder  string
}

func (this *PaymentListParam) QueryString() string {
	var p = url.Values{}
	if len(this.StartId) > 0 {
		p.Set("start_id", this.StartId)
	}
	if len(this.StartTime) > 0 {
		p.Set("start_time", this.StartTime)
	}
	if len(this.EndTime) > 0 {
		p.Set("end_time", this.EndTime)
	}
	if this.StartIndex > 0 {
		p.Set("start_index", fmt.Sprintf("%d", this.StartIndex))
	}
	if this.Count > 0 {
		p.Set("count", fmt.Sprintf("%f", this.Count))
	}
	if len(this.SortBy) > 0 {
		p.Set("sort_by", this.SortBy)
	}
	if len(this.SortOrder) > 0 {
		p.Set("sort_order", this.SortOrder)
	}
	return "?" + p.Encode()
}

type PaymentList struct {
	Payments []*Payment `json:"payments"`
	Count    int        `json:"count"`
	NextId   string     `json:"next_id"`
}

type refundSaleParam struct {
	Amount struct {
		Total    string `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	InvoiceNumber string `json:"invoice_number"`
}
