package handler

import "fmt"

func errParamRequired(name, typ string) error{
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreeateOpening

type CreateOpeningRequest struct {
	Role      string     `json:"role"`
	Company   string     `json:"company"`
	Location  string     `json:"location"`
	Remote    *bool      `json:"remote"`
	Link      string     `json:"link"`
	Salary    int64      `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamRequired("role", "string")
	}
	if r.Company == "" {
		return errParamRequired("company", "string")
	}
	if r.Location == "" {
		return errParamRequired("location", "string")
	}
	if r.Link == "" {
		return errParamRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}
	return nil
}

// Update Opening
type UpdateOpeningRequest struct {
	Role      string     `json:"role"`
	Company   string     `json:"company"`
	Location  string     `json:"location"`
	Remote    *bool      `json:"remote"`
	Link      string     `json:"link"`
	Salary    int64      `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error{
	// If any field is provided, validate is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	// if none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}