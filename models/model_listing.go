package models

type Listing struct {

	Id NftId `json:"id,omitempty"`

	Owner CosmosAddress `json:"owner,omitempty"`

	ListingType *interface{} `json:"listing_type,omitempty"`

	State *interface{} `json:"state,omitempty"`

	BidToken *interface{} `json:"bid_token,omitempty"`

	MinimumDepositRate *interface{} `json:"minimum_deposit_rate,omitempty"`

	AutomaticRefinancing *interface{} `json:"automatic_refinancing,omitempty"`

	StartedAt *interface{} `json:"started_at,omitempty"`

	EndAt *interface{} `json:"end_at,omitempty"`

	FullPaymentEndAt *interface{} `json:"full_payment_end_at,omitempty"`

	CollectedAmount Denom `json:"collected_amount,omitempty"`

	MinimumBiddingPeriod *interface{} `json:"minimum_bidding_period,omitempty"`
}
