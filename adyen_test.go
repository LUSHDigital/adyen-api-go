package adyen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Set environment variables for subsequent tests.
	if err := godotenv.Load(".default.env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	os.Exit(m.Run())
}

func TestNewWithTimeout(t *testing.T) {
	const timeout = time.Second * 123

	act := New(&APICredentials{
		Env:      Testing,
		Username: "un",
	}, WithTimeout(timeout))
	equals(t, timeout, act.client.Timeout)
}

func TestNewWithCurrency(t *testing.T) {
	const currency = "USD"

	act := New(&APICredentials{
		Env:      Testing,
		Username: "un",
	}, WithCurrency(currency))
	equals(t, currency, act.Currency)
}

func TestNewWithCustomOptions(t *testing.T) {
	const merchant, currency, timeout = "merch", "JPY", time.Second * 21

	f1 := func(a *Adyen) {
		a.Currency = currency
		a.client.Timeout = timeout
	}

	f2 := func(a *Adyen) {
		a.MerchantAccount = merchant
	}

	act := New(&APICredentials{
		Env:      Testing,
		Username: "un",
	}, f1, f2)
	equals(t, merchant, act.MerchantAccount)
	equals(t, currency, act.Currency)
	equals(t, timeout, act.client.Timeout)
}

func equals(tb *testing.T, exp interface{}, act interface{}) {
	_, fullPath, line, _ := runtime.Caller(1)
	file := filepath.Base(fullPath)

	if !reflect.DeepEqual(exp, act) {
		fmt.Printf("%s:%d:\n\texp: %[3]v (%[3]T)\n\tgot: %[4]v (%[4]T)\n", file, line, exp, act)
		tb.FailNow()
	}
}

func assert(tb *testing.T, cond bool, message string) {
	_, fullPath, line, _ := runtime.Caller(1)
	file := filepath.Base(fullPath)

	if !cond {
		fmt.Printf("%s:%d:\n\t%s\n", file, line, message)
		tb.FailNow()
	}
}

func getHTTPMockInstance(host string) *Adyen {
	env := Environment{
		apiURL:      host,
		clientURL:   host,
		hppURL:      host,
		checkoutURL: host,
	}

	instance := New(
		&APICredentials{
			Env:    env,
			APIKey: os.Getenv("ADYEN_API_KEY"),
		},
	)

	return instance
}
func getTestInstance() *Adyen {
	instance := New(
		&APICredentials{
			Env:    Testing,
			APIKey: os.Getenv("ADYEN_API_KEY"),
		}, )

	return instance
}
func getTestHMACInstance() *Adyen {
	instance := New(
		&APICredentials{
			Env:      Testing,
			Username: os.Getenv("ADYEN_USERNAME"),
			Password: os.Getenv("ADYEN_PASSWORD"),
			HMAC:     os.Getenv("ADYEN_HMAC"),
		},
	)

	return instance
}

// randInt - get random integer from a given range
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// randomString - generate randorm string of given length
// note: not for use in live code
func randomString(l int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// createTestResponse - create response object for tests
func createTestResponse(input, status string, code int) (*Response, error) {
	body := strings.NewReader(input)

	resp := &http.Response{
		Status:        status,
		StatusCode:    code,
		ContentLength: int64(body.Len()),
		Body:          ioutil.NopCloser(body),
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)

	if err != nil {
		return nil, err
	}

	response := &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	return response, nil
}
