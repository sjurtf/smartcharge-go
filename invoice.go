package smartcharge

import (
	"net/http"
	"strconv"
)

type InvoiceService struct {
	client *Client
}

type InvoiceResult struct {
	Result Data `json:"Result"`
}

type Data struct {
	Data    InvoiceData `json:"datas"`
	Success bool        `json:"sucess"`
}

type InvoiceData struct {
	Invoices []Invoice
}

type Invoice struct {
	Id         int     `json:"PK_InvoiceID"`
	TotalPrice float64 `json:"TotalPrice"`
	TotalVAT   float64 `json:"TotalVAT"`
	DateIssued string  `json:"DateIssued"`
	TotalKWH   float64 `json:"TotalKWH"`
}

func (i *InvoiceService) GetInvoices(customerId int) (*InvoiceResult, *http.Response, error) {

	reqUrl := "v2/Invoices/GetInvoicesByCustomer/" + strconv.Itoa(customerId)
	req, err := i.client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	invoiceResult := &InvoiceResult{}
	resp, err := i.client.Do(req, invoiceResult)
	if err != nil {
		return nil, resp, err
	}

	return invoiceResult, resp, err
}
