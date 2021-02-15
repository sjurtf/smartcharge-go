package smartcharge

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestInvoiceService_GetInvoices(t *testing.T) {
	jsonBytes, _ := ioutil.ReadFile("testdata/invoices/invoices.json")

	r := *NewMockResponseOkString(string(jsonBytes))
	c := NewMockClient(r)

	data, _, err := c.Invoice.GetInvoices(123)

	assert.NoError(t, err)
	assert.Equal(t, 414005, data.Result.Data.Invoices[0].Id)
	assert.Equal(t, 87.407, data.Result.Data.Invoices[0].TotalKWH)
}
