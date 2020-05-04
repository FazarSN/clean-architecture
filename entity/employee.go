package entity

type Employee struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Street    string `json:"street"`
	CompanyID int    `json:"company_id"`
}

func (e *Employee) Validate() error {
	return nil
}
