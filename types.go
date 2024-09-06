package stayntouch

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
