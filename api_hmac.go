package zestapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/zestlabs-io/zest-go-sdk/api/client"

	httptransport "github.com/go-openapi/runtime/client"
)

func NewHMACAPIClient(formats strfmt.Registry, cfg *client.TransportConfig, cloudKey, cloudSecret string) *client.ZestAPI {
	rt := SigningRoundTripper{
		Proxied:     http.DefaultTransport,
		cloudKey:    cloudKey,
		cloudSecret: cloudSecret,
	}
	httpClient := &http.Client{Transport: rt}
	transport := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, httpClient)
	transport.Producers["*/*"] = runtime.ByteStreamProducer()
	transport.Consumers["*/*"] = runtime.ByteStreamConsumer()
	cl := client.New(transport, formats)

	return cl
}

// This type implements the http.RoundTripper interface
type SigningRoundTripper struct {
	Proxied     http.RoundTripper
	cloudKey    string
	cloudSecret string
}

// Generates the authorization token in a form
// `ZEST cloudKey:HMAC_SIGN(<cloudSecret>, reqPath + '\n' + contentType + '\n' + validity + '\n' + md5Body)`
// This further should be passed as `Authorization` header
func (lrt SigningRoundTripper) RoundTrip(r *http.Request) (res *http.Response, e error) {
	// Do "before sending requests" actions here.
	if r == nil {
		panic(fmt.Errorf("WTF - request is nil"))
	}
	b := []byte{}
	if r.Body != nil {
		b, _ = ioutil.ReadAll(r.Body)
		r.Body.Close()

	}
	// fmt.Printf("Signing request to %v\n", r.URL)

	body := ioutil.NopCloser(bytes.NewReader(b))

	r.Body = body
	size := len(b)
	r.ContentLength = int64(size)
	r.Header.Set("Content-Length", strconv.Itoa(size))

	md5Checksum := createMD5(b)
	contentType := r.Header.Get("Content-Type")
	validity := time.Now().Add(time.Second * 60).Unix()
	authString := createAuthorization(r.URL.String(), contentType, validity, md5Checksum, lrt.cloudKey, lrt.cloudSecret)

	r.Header["X-ZEST-Validity"] = []string{strconv.FormatInt(validity, 10)}
	r.Header["Content-Md5"] = []string{md5Checksum}
	r.Header["Content-Type"] = []string{contentType}
	r.Header["Authorization"] = []string{authString}

	// Send the request, get the response (or the error)
	return lrt.Proxied.RoundTrip(r)
}

func createMD5(in []byte) string {
	md5Sum := md5.Sum(in)
	return fmt.Sprintf("%x", md5Sum)
}

func createAuthorization(url, contentType string, validity int64, md5Checksum, cloudKey, cloudSecret string) string {
	return fmt.Sprintf("ZEST %s:%s", cloudKey, createSignature(url, contentType, md5Checksum, validity, cloudSecret))
}

// POST /api/distr-config/create-pool
// Authorization: ZEST <AUTH_ID_KEY>:HMAC_SIGN(<AUTH_ID_SECRET>, 'POST' + '\n' + '/api/distr-config/create-pool' + '\n' + <CONTENT_TYPE> + '\n' + <VALIDITY_SEC> + '\n' + <CONTENT_MD5>)
// Content-Type: <CONTENT_TYPE>
// Content-Md5: <CONTENT_MD5>
// X-ZEST-Validity: <VALIDITY_SEC>
func createSignature(url, contentType, md5Checksum string, validity int64, cloudSecret string) string {
	toSign := fmt.Sprintf("%s\n%s\n%d\n%s", url, contentType, validity, md5Checksum)
	mac := hmac.New(sha256.New, []byte(cloudSecret))
	_, _ = mac.Write([]byte(toSign))
	macSignature := mac.Sum(nil)
	return hex.EncodeToString(macSignature)
}
