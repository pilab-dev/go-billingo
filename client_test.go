package billingo_test

import (
	"context"
	"testing"
	"time"

	"github.cm/pilab-dev/go-billingo/v3"
	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/require"
)

const testApiKey = "76664878-cd41-11ed-bb04-06ac9760f844"

func TestCreateInvoice(t *testing.T) {
	c, err := billingo.New(testApiKey)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("TestCreateProduct", func(t *testing.T) {
		res, err := c.CreateProductWithResponse(ctx, billingo.Product{
			Comment:              new(string),
			Currency:             "HUF",
			Entitlement:          billingo.ToPtr(billingo.EntitlementAAM),
			GeneralLedgerNumber:  new(string),
			GeneralLedgerTaxcode: new(string),
			Id:                   billingo.ToPtr(1),
			Name:                 "Test product",
			NetUnitPrice:         billingo.ToPtr(float32(600.1)),
			Unit:                 "db",
			Vat:                  "27%",
		})

		require.NoError(t, err)
		require.Equalf(t, 201, res.StatusCode(), "Expected status code 201, got %d", res.StatusCode())
	})

	t.Run("TestCreatePartner", func(t *testing.T) {
		res, err := c.CreatePartnerWithResponse(ctx, billingo.Partner{
			AccountNumber: new(string),
			Address: &billingo.Address{
				Address:     "Bela utca 11.",
				City:        "Budapest",
				CountryCode: "HU",
				PostCode:    "1116",
			},
			CustomBillingSettings: &billingo.PartnerCustomBillingSettings{},
			Emails:                &[]string{},
			GeneralLedgerNumber:   new(string),
			GroupMemberTaxNumber:  new(string),
			Iban:                  new(string),
			Id:                    billingo.ToPtr(4),
			Name:                  billingo.ToPtr("Test partner"),
			Phone:                 new(string),
			Swift:                 new(string),
			TaxType:               billingo.ToPtr(billingo.PartnerTaxTypeEmpty),
			Taxcode:               new(string),
		})

		if res.JSON422 != nil {
			t.Logf("Error 422: %+v", *res.JSON422.Message)
			for _, e := range *res.JSON422.Errors {
				t.Logf("Field: %s, Msg: %+v", *e.Field, *e.Message)
			}
		}

		require.NoError(t, err)
		require.Equalf(t, 201, res.StatusCode(), "Expected status code 201, got %d", res.StatusCode())
	})

	t.Run("CreateInvoice", func(t *testing.T) {
		var testItem billingo.DocumentInsert_Items_Item

		testItem.FromDocumentProductData(billingo.DocumentProductData{
			Comment:       new(string),
			Entitlement:   billingo.ToPtr(billingo.EntitlementAAM),
			Name:          "Maci",
			Quantity:      1,
			Unit:          "db",
			UnitPrice:     200,
			UnitPriceType: billingo.Gross,
			Vat:           "27%",
		})

		items := []billingo.DocumentInsert_Items_Item{
			testItem,
		}

		inv, err := c.CreateDocumentWithResponse(ctx, billingo.DocumentInsert{
			// AdvanceInvoice: &[]int{},
			// BankAccountId:  billingo.ToPtr(0),
			BlockId:        0,
			Comment:        new(string),
			ConversionRate: billingo.ToPtr(float32(1.0)),
			Currency:       "HUF",
			// Discount: &billingo.Discount{
			// 	Type:  billingo.ToPtr(billingo.DiscountType("percent")),
			// 	Value: billingo.ToPtr(0),
			// },
			DueDate: types.Date{
				Time: time.Now().Add(time.Hour * 24 * 30), // 30 days
			},
			Electronic: billingo.ToPtr(false),
			FulfillmentDate: types.Date{
				Time: time.Now().Add(24 * 30 * time.Hour), // 30 days
			},
			InstantPayment: billingo.ToPtr(false),
			Language:       "hu",
			Paid:           billingo.ToPtr(true),
			PartnerId:      4,
			PaymentMethod:  billingo.PaymentMethodCash,
			// Settings:       &billingo.DocumentSettings{},
			Type:     "invoice",
			VendorId: nil,
			Items:    &items,
		})

		// t.Logf("Error: %+v", *inv.JSON401.Error.Message)

		if inv.JSON422 != nil {
			t.Logf("Error 422: %+v", *inv.JSON422.Message)
			for _, e := range *inv.JSON422.Errors {
				t.Logf("Field: %s, Msg: %+v", *e.Field, *e.Message)
			}
		}

		if inv.JSON403 != nil {
			t.Logf("Error 403: %+v", *inv.JSON403.Error.Message)
		}

		require.NoError(t, err)
		require.Equal(t, 201, inv.StatusCode())
		require.Equal(t, "application/json", inv.HTTPResponse.Header.Get("Content-Type"))

		t.Logf("Invoice created: %d", *inv.JSON201.Id)
	})
}
