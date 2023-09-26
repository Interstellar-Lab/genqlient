// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package queries

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// SimpleQueryResponse is returned by SimpleQuery on success.
type SimpleQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryUser `json:"user"`
}

// GetUser returns SimpleQueryResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryResponse) GetUser() SimpleQueryUser { return v.User }

// SimpleQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `json:"id"`
}

// GetId returns SimpleQueryUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryUser) GetId() string { return v.Id }

// The query executed by SimpleQuery.
const SimpleQuery_Operation = `
query SimpleQuery {
	user {
		id
	}
}
`

func SimpleQuery(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *SimpleQueryResponse, ext_ map[string]interface{}, err_ error) {
	req_ := &graphql.Request{
		OpName: "SimpleQuery",
		Query:  SimpleQuery_Operation,
	}

	if client_ == nil {
		return nil, nil, fmt.Errorf("got nil graphql.Client")
	}

	data_ = &SimpleQueryResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, resp_.Extensions, err_
}

