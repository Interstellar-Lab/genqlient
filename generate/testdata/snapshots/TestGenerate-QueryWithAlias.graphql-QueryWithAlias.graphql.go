// Code generated by github.com/Interstellar-Lab/genqlient, DO NOT EDIT.

package test

import (
	"fmt"

	"github.com/Interstellar-Lab/genqlient/graphql"
	"github.com/Interstellar-Lab/genqlient/internal/testutil"
)

// QueryWithAliasResponse is returned by QueryWithAlias on success.
type QueryWithAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithAliasUser `json:"User"`
}

// GetUser returns QueryWithAliasResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithAliasResponse) GetUser() QueryWithAliasUser { return v.User }

// QueryWithAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID testutil.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	OtherID testutil.ID `json:"otherID"`
}

// GetID returns QueryWithAliasUser.ID, and is useful for accessing the field via an interface.
func (v *QueryWithAliasUser) GetID() testutil.ID { return v.ID }

// GetOtherID returns QueryWithAliasUser.OtherID, and is useful for accessing the field via an interface.
func (v *QueryWithAliasUser) GetOtherID() testutil.ID { return v.OtherID }

// The query executed by QueryWithAlias.
const QueryWithAlias_Operation = `
query QueryWithAlias {
	User: user {
		ID: id
		otherID: id
	}
}
`

func QueryWithAlias(
	client_ graphql.Client,
) (data_ *QueryWithAliasResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "QueryWithAlias",
		Query:  QueryWithAlias_Operation,
	}

	if client_ == nil {
		return nil, fmt.Errorf("got nil graphql.Client")
	}

	data_ = &QueryWithAliasResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

