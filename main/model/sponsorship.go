package model

type TierDefine struct {
	Description           string `json:"description"`
	Name                  string `json:"name"`
	MonthlyPriceInCents   int32  `json:"monthly_price_in_cents"`
	MonthlyPriceInDollars int32  `json:"monthly_price_in_dollars"`
}

type SponsorshipDefine struct {
	Sponsorable  Owner      `json:"sponsorable"`
	Sponsor      Owner      `json:"sponsor"`
	PrivacyLevel string     `json:"privacy_level"`
	Tier         TierDefine `json:"tier"`
}

type Sponsorship struct {
	Common
	Sponsorship SponsorshipDefine `json:"sponsorship"`
}
