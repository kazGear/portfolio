package model

import "time"

// TODO: ポインタ(nilAble)は後々減らしていく
type Job struct {
	Id                  int64      `db:"id"`
	Url                 string     `db:"url"`
	Title               string     `db:"title"`
	CompanyName         string     `db:"company_name"`
	Location            string     `db:"location"`
	MinSalaryAtHour     *int       `db:"min_salary_at_hour"`
	MaxSalaryAtHour     *int       `db:"max_salary_at_hour"`
	MinSalaryAtMonth    *int       `db:"min_salary_at_month"`
	MaxSalaryAtMonth    *int       `db:"max_salary_at_month"`
	SkillsText 	        string     `db:"skills_text"`
	RequiredSkillsText  string     `db:"required_skills_text"`
	PreferredSkillsText string     `db:"preferred_skills_text"`
	Description         string     `db:"description"`
	EmploymentType      string     `db:"employment_type"`
	WorkPlace          string      `db:"work_place"`
	IsActive            bool       `db:"is_active"`
	SimilarityScore     *float64   `db:"similarity_score"`
	SourceSite          string     `db:"source_site"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           *time.Time `db:"updated_at"`
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
