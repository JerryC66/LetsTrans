package response

import "github.com/firwoodlin/letstrans/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
