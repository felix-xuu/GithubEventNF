package model

type PlanDefine struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	MonthlyPriceInCents int32    `json:"monthly_price_in_cents"`
	YearlyPriceInCents  int32    `json:"yearly_price_in_cents"`
	PriceModel          string   `json:"price_model"`
	HasFreeTrial        bool     `json:"has_free_trial"`
	UnitName            string   `json:"unit_name"`
	Bullets             []string `json:"bullets"`
}

type MarketplacePurchaseDefine struct {
	Account         Owner      `json:"account"`
	BillingCycle    string     `json:"billing_cycle"`
	UnitCount       int32      `json:"unit_count"`
	OnFreeTrial     bool       `json:"on_free_trial"`
	FreeTrialEndsOn string     `json:"free_trial_ends_on"`
	NextBillingDate string     `json:"next_billing_date"`
	Plan            PlanDefine `json:"plan"`
}

type MarketplacePurchase struct {
	Common
	MarketplacePurchase MarketplacePurchaseDefine `json:"marketplace_purchase"`
}
