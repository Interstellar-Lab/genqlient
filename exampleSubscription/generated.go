// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"

	"github.com/Interstellar-Lab/genqlient/graphql"
)

// countResponse is returned by count on success.
type countResponse struct {
	Count int `json:"count"`
}

// GetCount returns countResponse.Count, and is useful for accessing the field via an interface.
func (v *countResponse) GetCount() int { return v.Count }

// The query, mutation or subscription executed by count.
const count_Operation = `
subscription count {
	count
}
`

// count
//
// To close the connection, use [graphql.Client.CloseWebSocket()]
func count(
	ctx_ context.Context,
	client_ graphql.Client,
) (dataChan_ chan countWsResponse, errChan_ chan error, err error) {
	req_ := &graphql.Request{
		OpName: "count",
		Query:  count_Operation,
	}
	var err_ error

	dataChan_ = make(chan countWsResponse, 1)
	respChan_ := make(chan json.RawMessage, 1)

	errChan_, err_ = client_.DialWebSocket(ctx_, req_, respChan_)
	if err_ != nil {
		return nil, nil, err_
	}
	go countForwardData(dataChan_, respChan_, errChan_)

	return dataChan_, errChan_, err_
}

type countWsResponse struct {
	Data       *countResponse         `json:"data"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	Errors     error                  `json:"errors"`
}

func countForwardData(dataChan_ chan countWsResponse, respChan_ chan json.RawMessage, errChan_ chan error) {
	defer close(dataChan_)
	var gqlResp graphql.Response
	var wsResp countWsResponse
	for {
		jsonRaw, more_ := <-respChan_
		if !more_ {
			return
		}
		err := json.Unmarshal(jsonRaw, &gqlResp)
		if err != nil {
			errChan_ <- err
			return
		}
		if len(gqlResp.Errors) == 0 {
			err = json.Unmarshal(jsonRaw, &wsResp)
			if err != nil {
				errChan_ <- err
				return
			}
		} else {
			wsResp.Errors = gqlResp.Errors
		}
		dataChan_ <- wsResp
	}
}
