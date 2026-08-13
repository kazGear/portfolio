package model

import "time"

type Job struct {
	Id                  int64      `db:"id"`
	Url                 string     `db:"url"`
	Title               string     `db:"title"`
	Location            string     `db:"location"`
	MinSalaryAtMonth    *int       `db:"min_salary_at_month"`
	MaxSalaryAtMonth    *int       `db:"max_salary_at_month"`
	Description         string     `db:"description"`
	EmploymentType      string     `db:"employment_type"`
	WorkPlace           string     `db:"work_place"`
	SourceSite          string     `db:"source_site"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           *time.Time `db:"updated_at"`
}

type JobFeature struct {
	JobId           int64  `db:"job_id"`
	FeatureName     string `db:"feature_name"`
	RequirementType string `db:"requirement_type"`
	Category        string `db:"category"`
}

type JobOption struct {
	JobId  int64  `db:"job_id"`
	Option string `db:"option"`
}

type ApiResponseCrowdworksTech struct {
	ID                      int      `json:"id"`
	JobOfferID              int      `json:"jobOfferId"`
	ClientName              string   `json:"clientName"`
	DetailedTitle           string   `json:"detailedTitle"`
	WorkingHours            string   `json:"workingHours"`
	SpecificWorkContent     string   `json:"specificWorkContent"`
	RelatedServicesProducts string   `json:"relatedServicesProducts"`
	RequirementSkills       string   `json:"requirementSkills"`
	CommercialDistribution  *string  `json:"commercialDistribution"`
	MatchingWorkingDays     []string `json:"matchingWorkingDays"`
	MatchingResidentDays    []string `json:"matchingResidentDays"`
	DevelopmentLanguages    []string `json:"developmentLanguages"`
	Frameworks              []string `json:"frameworks"`
	Databases               []string `json:"databases"`
	Tools                   []string `json:"tools"`
	Infrastructures         []string `json:"infrastructures"`
	Designs                 []string `json:"designs"`
	PaymentMethods          []string `json:"paymentMethods"`
}
