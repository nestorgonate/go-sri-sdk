package gosrisdk_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/nestorgonate/go-sri-sdk"
)

func TestTaxpayerServiceListByRuc(t *testing.T) {
	c := context.Background()
	client := gosrisdk.NewClient()
	consultaRuc, err := client.Taxpayer.ListByRuc(c, os.Getenv("RUC"))
	if err != nil {
		if errors.Is(err, gosrisdk.ErrResponseNil) {
			t.Skip(err)
		}
		if errors.Is(err, gosrisdk.ErrRucNoRegistrado) {
			t.Log(err)
			return
		}
		t.Fatal(err)
	}
	consultaRucBytes, err := json.MarshalIndent(consultaRuc, "", "     ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(consultaRucBytes))
}

func TestTaxpayerServiceIsValidRuc(t *testing.T) {
	c := context.Background()
	client := gosrisdk.NewClient()
	consultaRuc, err := client.Taxpayer.IsValidRuc(c, os.Getenv("RUC"))
	if err != nil {
		if errors.Is(err, gosrisdk.ErrResponseNil) {
			t.Skip(err)
		}
		if errors.Is(err, gosrisdk.ErrRucVacio) || errors.Is(err, gosrisdk.ErrRucNoValido) {
			t.Log(err.Error())
			return
		}
		t.Log(err.Error())
		return
	}
	t.Log(consultaRuc)
}

func TestTaxpayerServiceListFacilitiesByRuc(t *testing.T) {
	c := context.Background()
	client := gosrisdk.NewClient()
	establecimientos, err := client.Taxpayer.ListFacilitiesByRuc(c, os.Getenv("RUC"))
	if err != nil {
		if errors.Is(err, gosrisdk.ErrResponseNil) {
			t.Skip(err.Error())
		}
		if errors.Is(err, gosrisdk.ErrRucNoRegistrado) {
			t.Log(err.Error())
			return
		}
		t.Fatal(err)
	}
	establecimientosBytes, err := json.MarshalIndent(establecimientos, "", "     ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(establecimientosBytes))
}
