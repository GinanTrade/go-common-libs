package authmodel

type Auth struct {
	TokenType string `json:"token_type"`
	Jti       string `json:"jti"`
	Email     string `json:"email"`
	//Sub       string  `json:"sub"`
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	Iat       float64 `json:"iat"`
	Exp       float64 `json:"exp"`
	AccountId string  `json:"account_id"`
	CompanyId float64 `json:"company_id"`
}
