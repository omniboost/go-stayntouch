package stayntouch

type RevenueLedger string

const (
	RevenueGuestLedger   RevenueLedger = "guest"
	RevenueDepositLedger RevenueLedger = "deposit"
	RevenueArLedger      RevenueLedger = "ar"
)

type Hotels []Hotel

type Hotel struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	CheckInTime  string `json:"check_in_time"`
	CheckOutTime string `json:"check_out_time"`
	Address      struct {
		Address1   string `json:"address1"`
		Address2   string `json:"address2"`
		Address3   string `json:"address3"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postal_code"`
		Country    string `json:"country"`
	} `json:"address"`
	Phone       string `json:"phone"`
	MainContact struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	} `json:"main_contact"`
	CurrencyCode string `json:"currency_code"`
	Timezone     string `json:"timezone"`
}

type Reservations []Reservation

type Reservation struct {
	ID                 int    `json:"id"`
	HotelID            int    `json:"hotel_id"`
	ConfirmationNumber string `json:"confirmation_number"`
	CurrencyCode       string `json:"currency_code"`
	Status             string `json:"status"`
	ArrivalDate        string `json:"arrival_date"`
	ArrivalTime        string `json:"arrival_time"`
	DepartureDate      string `json:"departure_date"`
	DepartureTime      string `json:"departure_time"`
	Addons             []struct {
		ID                  int      `json:"id"`
		Name                string   `json:"name"`
		Description         string   `json:"description"`
		Amount              float64  `json:"amount"`
		AmountType          string   `json:"amount_type"`
		PostType            string   `json:"post_type"`
		PostTypeDescription string   `json:"post_type_description"`
		PostDays            []string `json:"post_days"`
		ChargeCodeID        int      `json:"charge_code_id"`
		ChargeGroupID       int      `json:"charge_group_id"`
		CurrencyCode        string   `json:"currency_code"`
		Quantity            int      `json:"quantity"`
		IncludedInRate      bool     `json:"included_in_rate"`
	} `json:"addons"`
	Bills []struct {
		BillNumber    int    `json:"bill_number"`
		ID            int    `json:"id"`
		InvoiceNumber string `json:"invoice_number"`
		ReservationID int    `json:"reservation_id"`
	} `json:"bills"`
	CreatedTime       string  `json:"created_time"`
	DepositAmount     float64 `json:"deposit_amount"`
	MarketSegmentCode string  `json:"market_segment_code"`
	Notes             []struct {
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"notes,omitempty"`
	PromotionCode       string `json:"promotion_code"`
	ReservationTypeCode string `json:"reservation_type_code,omitempty"`
	SegmentCode         string `json:"segment_code"`
	SourceCode          string `json:"source_code,omitempty"`
	TaxExempt           bool   `json:"tax_exempt"`
	Room                struct {
		ID            int    `json:"id"`
		Number        string `json:"number"`
		RoomTypeID    int    `json:"room_type_id"`
		Status        string `json:"status"`
		ServiceStatus string `json:"service_status"`
		Occupied      bool   `json:"occupied"`
		MaxOccupancy  int    `json:"max_occupancy"`
		SuiteDoor     bool   `json:"suite_door"`
		Floor         struct {
			Number      string `json:"number"`
			Description string `json:"description"`
		} `json:"floor"`
		Features []string `json:"features"`
	} `json:"room,omitempty"`
	StayDates []struct {
		Date               string `json:"date"`
		ID                 int    `json:"id"`
		Amount             string `json:"amount"`
		NetAmount          string `json:"net_amount"`
		TaxAmount          string `json:"tax_amount"`
		ExclusiveTaxAmount string `json:"exclusive_tax_amount"`
		InclusiveTaxAmount string `json:"inclusive_tax_amount"`
		RateID             int    `json:"rate_id"`
		RateSuppressed     bool   `json:"rate_suppressed"`
		RoomTypeID         int    `json:"room_type_id"`
		Adults             int    `json:"adults"`
		Children           int    `json:"children"`
		Infants            int    `json:"infants"`
		OriginalRoomTypeID int    `json:"original_room_type_id"`
		Room               struct {
			ID            int    `json:"id"`
			Number        string `json:"number"`
			RoomTypeID    int    `json:"room_type_id"`
			Status        string `json:"status"`
			ServiceStatus string `json:"service_status"`
			Occupied      bool   `json:"occupied"`
			MaxOccupancy  int    `json:"max_occupancy"`
			SuiteDoor     bool   `json:"suite_door"`
			Floor         struct {
				Number      string `json:"number"`
				Description string `json:"description"`
			} `json:"floor"`
			Features []string `json:"features"`
		} `json:"room"`
	} `json:"stay_dates"`
	Guests []struct {
		Address struct {
			City        string `json:"city"`
			CountryCode string `json:"country_code"`
			ID          int    `json:"id"`
		} `json:"address"`
		ID                     int    `json:"id"`
		IsPrimary              bool   `json:"is_primary"`
		FirstName              string `json:"first_name"`
		LastName               string `json:"last_name"`
		IsVip                  bool   `json:"is_vip"`
		Language               string `json:"language"`
		Nationality            string `json:"nationality"`
		OptedPromotionalEmails bool   `json:"opted_promotional_emails"`
		Title                  string `json:"title"`
		IsFlagged              bool   `json:"is_flagged"`
	} `json:"guests"`
	UpsellApplied       bool   `json:"upsell_applied"`
	EarlyCheckInApplied bool   `json:"early_check_in_applied"`
	LateCheckOutApplied bool   `json:"late_check_out_applied"`
	UpdatedTime         string `json:"updated_time"`
	Creator             struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"creator"`
	RestrictPost bool `json:"restrict_post"`
	DoNotMove    bool `json:"do_not_move"`
	Company      struct {
		Address struct {
			AddressLine1 string `json:"address_line1"`
			AddressLine2 string `json:"address_line2"`
			AddressLine3 string `json:"address_line3"`
			City         string `json:"city"`
			CountryCode  string `json:"country_code"`
			ID           int    `json:"id"`
			PostalCode   string `json:"postal_code"`
			State        string `json:"state"`
		} `json:"address"`
		PrimaryContact struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			JobTitle  string `json:"job_title"`
			Phone     string `json:"phone"`
		} `json:"primary_contact"`
		Commision struct {
			On bool `json:"on"`
		} `json:"commision"`
		ArNumber               string `json:"ar_number"`
		HotelID                int    `json:"hotel_id"`
		ID                     int    `json:"id"`
		IsGlobal               bool   `json:"is_global"`
		Name                   string `json:"name"`
		Type                   string `json:"type"`
		OrganizationIdentifier string `json:"organization_identifier"`
		Email                  string `json:"email"`
		Phone                  string `json:"phone"`
	} `json:"company,omitempty"`
}

type ChargeCodesBasicTaxes struct {
	ID          int    `json:"id"`
	ChargeCode  string `json:"charge_code"`
	Description string `json:"description"`
	Taxes       []struct {
		IsInclusive           bool   `json:"is_inclusive"`
		ChargeCodeID          int    `json:"charge_code_id"`
		ChargeCode            string `json:"charge_code"`
		ChargeCodeDescription string `json:"charge_code_description"`
		Percentage            string `json:"percentage"`
		AmountType            string `json:"amount_type"`
		PostType              string `json:"post_type"`
		ChargeGroup           string `json:"charge_group"`
		ChargeCodeType        string `json:"charge_code_type"`
	} `json:"taxes"`
}

type ChargeCodeTranslations []ChargeCodeTranslation

type ChargeCodeTranslation struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Translations []struct {
		LanguageCode string `json:"language_code"`
		Name         string `json:"name"`
		Description  string `json:"description"`
	} `json:"translations"`
}

type DirectBills []DirectBill

type DirectBill struct {
	Number           string `json:"number"`
	FolioNumber      string `json:"folio_number"`
	ArInvoiceNumber  string `json:"ar_invoice_number"`
	ModInvoiceNumber string `json:"mod_invoice_number"`
	CurrencyCode     string `json:"currency_code"`
	IssueDate        string `json:"issue_date"`
	PaymentDueDate   string `json:"payment_due_date"`
	Hotel            struct {
		ID   int    `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"hotel"`
	Reservation struct {
		ID                 int    `json:"id"`
		PrimaryGuestName   string `json:"primary_guest_name"`
		RoomNumber         string `json:"room_number"`
		ConfirmationNumber string `json:"confirmation_number"`
	} `json:"reservation"`
	PostingAccount struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Number string `json:"number"`
	} `json:"posting_account"`
	Account struct {
		Name                   string `json:"name"`
		ArNumber               string `json:"ar_number"`
		OrganizationIdentifier string `json:"organization_identifier"`
		CorporateID            string `json:"corporate_id"`
		TaxID                  string `json:"tax_id"`
		EInvoiceAddress        string `json:"e_invoice_address"`
		Address                struct {
			Line1      string `json:"line1"`
			Line2      string `json:"line2"`
			Line3      string `json:"line3"`
			City       string `json:"city"`
			Country    string `json:"country"`
			State      string `json:"state"`
			PostalCode string `json:"postal_code"`
		} `json:"address"`
	} `json:"account"`
	Transactions []struct {
		Items []struct {
			Description     string `json:"description"`
			Code            string `json:"code"`
			ReferenceText   string `json:"reference_text"`
			ReferenceNumber string `json:"reference_number"`
			Amount          string `json:"amount"`
			IsInclusive     bool   `json:"is_inclusive"`
			Date            string `json:"date"`
			Taxes           []struct {
				Description     string      `json:"description"`
				Code            string      `json:"code"`
				ReferenceText   string      `json:"reference_text"`
				ReferenceNumber string      `json:"reference_number"`
				Amount          StringFloat `json:"amount"`
				Date            string      `json:"date"`
				IsInclusive     bool        `json:"is_inclusive"`
				Percent         StringFloat `json:"percent"`
				FlatAmount      StringFloat `json:"flat_amount"`
				Type            string      `json:"type"`
			} `json:"taxes"`
		} `json:"items"`
		Code        string      `json:"code"`
		Description string      `json:"description"`
		TotalAmount StringFloat `json:"total_amount"`
	} `json:"transactions"`
	PayableAmount StringFloat `json:"payable_amount"`
}

type InvoicesPerDay []InvoicePerDay

type InvoicePerDay struct {
	Date          string  `json:"date"`
	HotelID       int     `json:"hotel_id"`
	Amount        float64 `json:"amount"`
	ArAccount     string  `json:"ar_account"`
	AccountNumber string  `json:"account_number"`
	BillID        int     `json:"bill_id"`
	InvoiceNumber string  `json:"invoice_number"`
}

type Guests []Guest

type Guest struct {
	Address struct {
		AddressLine1 string `json:"address_line1"`
		City         string `json:"city"`
		CountryCode  string `json:"country_code"`
		ID           int    `json:"id"`
		PostalCode   string `json:"postal_code"`
		State        string `json:"state"`
	} `json:"address,omitempty"`
	ID                     int    `json:"id"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	Birthday               string `json:"birthday,omitempty"`
	IsVip                  bool   `json:"is_vip"`
	Language               string `json:"language,omitempty"`
	Nationality            string `json:"nationality,omitempty"`
	OptedPromotionalEmails bool   `json:"opted_promotional_emails"`
	Title                  string `json:"title,omitempty"`
	IsFlagged              bool   `json:"is_flagged"`
	Email                  string `json:"email,omitempty"`
	WorksAt                string `json:"works_at,omitempty"`
	Likes                  []struct {
		Category string   `json:"category"`
		Names    []string `json:"names"`
	} `json:"likes,omitempty"`
	HomePhone   string `json:"home_phone,omitempty"`
	MobilePhone string `json:"mobile_phone,omitempty"`
}

type Rooms []Room

type Room struct {
	ID            int    `json:"id"`
	Number        string `json:"number"`
	RoomTypeID    int    `json:"room_type_id"`
	Status        string `json:"status"`
	ServiceStatus string `json:"service_status"`
	Occupied      bool   `json:"occupied"`
	MaxOccupancy  int    `json:"max_occupancy,omitempty"`
	SuiteDoor     bool   `json:"suite_door"`
	Floor         struct {
		Number      string `json:"number"`
		Description string `json:"description"`
	} `json:"floor"`
	Features []string `json:"features"`
}

type RoomType struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MaxOccupancy int    `json:"max_occupancy"`
	ImageURL     string `json:"image_url"`
	Suite        bool   `json:"suite"`
	Pseudo       bool   `json:"pseudo"`
}

type Rate struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrencyCode string `json:"currency_code"`
	Type         struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Classification string `json:"classification"`
		Active         bool   `json:"active"`
	} `json:"type"`
	ChargeCode struct {
		ID             int     `json:"id"`
		ChargeCode     string  `json:"charge_code"`
		Description    string  `json:"description"`
		ChargeGroup    string  `json:"charge_group"`
		ChargeCodeType string  `json:"charge_code_type"`
		Amount         float64 `json:"amount"`
		PostType       string  `json:"post_type"`
		AmountType     string  `json:"amount_type"`
		AmountSymbol   string  `json:"amount_symbol"`
	} `json:"charge_code"`
	BasedOnRate struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		BasedOnType  string `json:"based_on_type"`
		BasedOnValue string `json:"based_on_value"`
	} `json:"based_on_rate"`
	DiscountAllowed         bool   `json:"discount_allowed"`
	Active                  bool   `json:"active"`
	TaxInclusiveOrExclusive string `json:"tax_inclusive_or_exclusive"`
	SuppressRate            bool   `json:"suppress_rate"`
	MarketSegmentCode       string `json:"market_segment_code"`
	SourceCode              string `json:"source_code"`
	Commissionable          bool   `json:"commissionable"`
	DepositPolicy           struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"deposit_policy"`
	CancellationPolicy struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"cancellation_policy"`
	Addons []struct {
		ID                  int      `json:"id"`
		Name                string   `json:"name"`
		Description         string   `json:"description"`
		Amount              float64  `json:"amount"`
		AmountType          string   `json:"amount_type"`
		PostType            string   `json:"post_type"`
		PostTypeDescription string   `json:"post_type_description"`
		PostDays            []string `json:"post_days"`
		ChargeCodeID        int      `json:"charge_code_id"`
		ChargeGroupID       int      `json:"charge_group_id"`
		CurrencyCode        string   `json:"currency_code"`
		IncludedInRate      bool     `json:"included_in_rate"`
	} `json:"addons"`
	RoomTypes []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"room_types"`
	HourlyRate  bool `json:"hourly_rate"`
	PublicRate  bool `json:"public_rate"`
	Member      bool `json:"member"`
	PmsOnly     bool `json:"pms_only"`
	ChannelOnly bool `json:"channel_only"`
}

type Accounts struct {
	Results    []Account `json:"results"`
	TotalCount int       `json:"total_count"`
}

type Address struct {
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	AddressLine3 string `json:"address_line3"`
	City         string `json:"city"`
	CountryCode  string `json:"country_code"`
	ID           int    `json:"id"`
	PostalCode   string `json:"postal_code"`
	State        string `json:"state"`
}
type PrimaryContact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	JobTitle  string `json:"job_title"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
type Commision struct {
	On    bool   `json:"on"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Account struct {
	Address        Address        `json:"address"`
	PrimaryContact PrimaryContact `json:"primary_contact"`
	Commision      Commision      `json:"commision"`
	ArNumber       string         `json:"ar_number"`
	HotelID        int            `json:"hotel_id"`
	ID             int            `json:"id"`
	IsGlobal       bool           `json:"is_global"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	AccountNumber  string         `json:"account_number"`
}

type LedgerItems []LedgerItem

type LedgerItem struct {
	ChargeCode            string         `json:"charge_code"`
	ChargeCodeDescription string         `json:"charge_code_description"`
	NetAmount             float64        `json:"net_amount"`
	GrossAmount           float64        `json:"gross_amount"`
	Account               string         `json:"account"`
	VATCode               string         `json:"vat_code"`
	CostCenter            string         `json:"cost_center"`
	Taxes                 LedgerTaxItems `json:"taxes"`
}

type LedgerTaxItems []LedgerTaxItem

type LedgerTaxItem struct {
	TaxInclusive          bool    `json:"tax_inclusive"`
	ChargeCode            string  `json:"charge_code"`
	ChargeCodeDescription string  `json:"charge_code_description"`
	Amount                float64 `json:"amount"`
}
