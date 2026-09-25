package net

import (
	"context"
	"errors"
	"net"
	"net/http"
	"syscall"
	"time"
)

func IsConnectionReset(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	return false
}

func IsNetworkErr(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}

func ExecuteRetry(c context.Context, client *http.Client, request *http.Request) (*http.Response, error) {
	maxRetries := 2
	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := c.Err()
		if err == nil {
			return nil, err
		}
		response, err := client.Do(request)
		if err != nil {
			if IsConnectionReset(err) || IsNetworkErr(err) {
				if attempt == maxRetries {
					break
				}
				delay := time.Duration(attempt) * 500 * time.Millisecond
				time.Sleep(delay)
				continue
			}
			return nil, err
		}
		return response, nil
	}
	return nil, nil
}
