// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SimpleQueryWithPointerFalseOverrideResponse is returned by SimpleQueryWithPointerFalseOverride on success.
type SimpleQueryWithPointerFalseOverrideResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User SimpleQueryWithPointerFalseOverrideUser `json:"user"`
}

// GetUser returns SimpleQueryWithPointerFalseOverrideResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideResponse) GetUser() SimpleQueryWithPointerFalseOverrideUser {
	return v.User
}

// SimpleQueryWithPointerFalseOverrideUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryWithPointerFalseOverrideUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns SimpleQueryWithPointerFalseOverrideUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideUser) GetId() testutil.ID { return v.Id }

// GetName returns SimpleQueryWithPointerFalseOverrideUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideUser) GetName() string { return v.Name }

// The query executed by SimpleQueryWithPointerFalseOverride.
const SimpleQueryWithPointerFalseOverride_Operation = `
query SimpleQueryWithPointerFalseOverride {
	user {
		id
		name
	}
}
`

func SimpleQueryWithPointerFalseOverride(
	client_ graphql.Client,
) (data_ *SimpleQueryWithPointerFalseOverrideResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "SimpleQueryWithPointerFalseOverride",
		Query:  SimpleQueryWithPointerFalseOverride_Operation,
	}

	if client_ == nil {
		return nil, fmt.Errorf("got nil graphql.Client")
	}

	data_ = &SimpleQueryWithPointerFalseOverrideResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

