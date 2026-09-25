package gosrisdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gosrisdk/internal/net"
	"net/http"
)

type TaxpayerService struct {
	client *Client
}

const (
	taxpayerEndpoint = "/sri-catastro-sujeto-servicio-internet/rest"
	userAgent        = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/152.0.0.0 Safari/537.36"
)

func (this *TaxpayerService) IsValidRuc(c context.Context, ruc string, opts ...RequestOption) (bool, error) {
	var validRuc bool
	var response *http.Response
	if ruc == "" {
		return validRuc, ErrRucVacio
	}
	if len(ruc) < 13 {
		return validRuc, ErrRucNoValido
	}
	endpoint := this.client.baseUrl + taxpayerEndpoint + fmt.Sprintf("/ConsolidadoContribuyente/existePorNumeroRuc?numeroRuc=%v",
		ruc)
	request, err := http.NewRequestWithContext(c, http.MethodGet, endpoint, nil)
	if err != nil {
		return validRuc, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)
	for _, opt := range opts {
		opt(request)
	}
	response, err = this.client.httpClient.Do(request)
	if err != nil {
		response, err = net.ExecuteRetry(c, this.client.httpClient, request)
		if err != nil {
			return validRuc, err
		}
	}
	if response == nil {
		return validRuc, ErrResponseNil
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&validRuc)
	if err != nil {
		return validRuc, err
	}
	return validRuc, err
}

func (this *TaxpayerService) ListByRuc(c context.Context, ruc string, opts ...RequestOption) (ConsultaRuc, error) {
	var consultaRuc ConsultaRuc
	var response *http.Response
	endpoint := this.client.baseUrl + taxpayerEndpoint + fmt.Sprintf("/ConsultaRuc/obtenerPorNumerosRuc?&ruc=%v", ruc)
	request, err := http.NewRequestWithContext(c, http.MethodGet, endpoint, nil)
	if err != nil {
		return consultaRuc, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)
	for _, opt := range opts {
		opt(request)
	}
	response, err = this.client.httpClient.Do(request)
	if err != nil {
		response, err = net.ExecuteRetry(c, this.client.httpClient, request)
		if err != nil {
			return consultaRuc, err
		}
	}
	if response == nil {
		return consultaRuc, ErrResponseNil
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNoContent {
		return consultaRuc, ErrRucNoRegistrado
	}
	if response.StatusCode != http.StatusOK {
		return consultaRuc, errors.New(response.Status)
	}
	err = json.NewDecoder(response.Body).Decode(&consultaRuc)
	if err != nil {
		return consultaRuc, err
	}
	return consultaRuc, nil
}

func (this *TaxpayerService) ListFacilitiesByRuc(c context.Context, ruc string, opts ...RequestOption) ([]Establecimiento, error) {
	var establecimientos []Establecimiento
	var response *http.Response
	endpoint := this.client.baseUrl + taxpayerEndpoint + fmt.Sprintf("/Establecimiento/consultarPorNumeroRuc?numeroRuc=%v",
		ruc)
	request, err := http.NewRequestWithContext(c, http.MethodGet, endpoint, nil)
	if err != nil {
		return establecimientos, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)
	for _, opt := range opts {
		opt(request)
	}
	response, err = this.client.httpClient.Do(request)
	if err != nil {
		response, err = net.ExecuteRetry(c, this.client.httpClient, request)
		if err != nil {
			return establecimientos, err
		}
	}
	if response == nil || response.Body == nil {
		return establecimientos, ErrResponseNil
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNoContent {
		return establecimientos, ErrRucNoRegistrado
	}
	if response.StatusCode != http.StatusOK {
		return establecimientos, errors.New(response.Status)
	}
	err = json.NewDecoder(response.Body).Decode(&establecimientos)
	if err != nil {
		return establecimientos, err
	}
	return establecimientos, nil
}
