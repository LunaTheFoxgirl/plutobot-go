package prest

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty"
)

// RestHandlerFunction is a function type that handles rest requests.
//type RestHandlerFunction func(reqtype string, requests ...string) (success bool)

// RestHandlerFunction is a function type that handles rest requests.
//type RestHandlerFunction func(conn restConn, reqtype string, requests ...string) (success bool)

const (
	R_GET    = 0
	R_PUT    = 3
	R_POST   = 1
	R_DELETE = 2
)

type RestRequestData struct {
	Tag   string
	Value string
}

type RestHeader struct {
	Tag   string
	Value string
}

type RestRequest struct {
	Header      RestHeader
	AuthToken   string
	Directory   string
	RequestType int
}

type rcon struct {
	client *resty.Client
}

// RestConnection handles the connections.
type RestConnection struct {
	rcon rcon
}

func New() RestConnection {
	rc := rcon{client: resty.New()}
	return RestConnection{rcon: rc}
}

func (r *RestConnection) Request(req RestRequest, data []RestRequestData) (string, error) {

	dat := make(map[string]string, 0)
	for _, elm := range data {
		dat[elm.Tag] = elm.Value
	}

	var err error
	var ret *resty.Response

	if req.RequestType == R_GET {
		ret, err = r.rcon.client.R().
			SetQueryParams(dat).
			Get(req.Directory)
	} else if req.RequestType == R_POST {
		ret, err = r.rcon.client.R().
			SetBody(dat).
			SetHeader(req.Header.Tag, req.Header.Value).
			SetAuthToken(req.AuthToken).
			Post(req.Directory)
	} else {
		_ = err
		_ = ret
		return "", errors.New("No valid RequestType was defined!")
	}

	if err != nil {
		fmt.Errorf(err.Error())
		return "", err
	}
	return ret.String(), nil
}
