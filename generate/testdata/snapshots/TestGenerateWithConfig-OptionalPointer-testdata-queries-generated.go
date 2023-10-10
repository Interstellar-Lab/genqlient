// Code generated by github.com/Interstellar-Lab/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/Interstellar-Lab/genqlient/graphql"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *ListInputQueryUser `json:"user"`
}

// GetUser returns ListInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *ListInputQueryResponse) GetUser() *ListInputQueryUser { return v.User }

// ListInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `json:"id"`
}

// GetId returns ListInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *ListInputQueryUser) GetId() string { return v.Id }

// QueryWithSlicesResponse is returned by QueryWithSlices on success.
type QueryWithSlicesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *QueryWithSlicesUser `json:"user"`
}

// GetUser returns QueryWithSlicesResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesResponse) GetUser() *QueryWithSlicesUser { return v.User }

// QueryWithSlicesUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithSlicesUser struct {
	Emails                []string  `json:"emails"`
	EmailsOrNull          []string  `json:"emailsOrNull"`
	EmailsWithNulls       []*string `json:"emailsWithNulls"`
	EmailsWithNullsOrNull []*string `json:"emailsWithNullsOrNull"`
}

// GetEmails returns QueryWithSlicesUser.Emails, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmails() []string { return v.Emails }

// GetEmailsOrNull returns QueryWithSlicesUser.EmailsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsOrNull() []string { return v.EmailsOrNull }

// GetEmailsWithNulls returns QueryWithSlicesUser.EmailsWithNulls, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNulls() []*string { return v.EmailsWithNulls }

// GetEmailsWithNullsOrNull returns QueryWithSlicesUser.EmailsWithNullsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNullsOrNull() []*string { return v.EmailsWithNullsOrNull }

// SimpleQueryNoOverrideResponse is returned by SimpleQueryNoOverride on success.
type SimpleQueryNoOverrideResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *SimpleQueryNoOverrideUser `json:"user"`
}

// GetUser returns SimpleQueryNoOverrideResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideResponse) GetUser() *SimpleQueryNoOverrideUser { return v.User }

// SimpleQueryNoOverrideUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleQueryNoOverrideUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   string  `json:"id"`
	Name *string `json:"name"`
}

// GetId returns SimpleQueryNoOverrideUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetId() string { return v.Id }

// GetName returns SimpleQueryNoOverrideUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleQueryNoOverrideUser) GetName() *string { return v.Name }

// SimpleQueryWithPointerFalseOverrideResponse is returned by SimpleQueryWithPointerFalseOverride on success.
type SimpleQueryWithPointerFalseOverrideResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *SimpleQueryWithPointerFalseOverrideUser `json:"user"`
}

// GetUser returns SimpleQueryWithPointerFalseOverrideResponse.User, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideResponse) GetUser() *SimpleQueryWithPointerFalseOverrideUser {
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
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetId returns SimpleQueryWithPointerFalseOverrideUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideUser) GetId() string { return v.Id }

// GetName returns SimpleQueryWithPointerFalseOverrideUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleQueryWithPointerFalseOverrideUser) GetName() string { return v.Name }

// __ListInputQueryInput is used internally by genqlient
type __ListInputQueryInput struct {
	Names []*string `json:"names"`
}

// GetNames returns __ListInputQueryInput.Names, and is useful for accessing the field via an interface.
func (v *__ListInputQueryInput) GetNames() []*string { return v.Names }

// The query executed by ListInputQuery.
const ListInputQuery_Operation = `
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`

func ListInputQuery(
	ctx_ context.Context,
	client_ graphql.Client,
	names []*string,
) (data_ *ListInputQueryResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "ListInputQuery",
		Query:  ListInputQuery_Operation,
		Variables: &__ListInputQueryInput{
			Names: names,
		},
	}

	data_ = &ListInputQueryResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by QueryWithSlices.
const QueryWithSlices_Operation = `
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`

func QueryWithSlices(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *QueryWithSlicesResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "QueryWithSlices",
		Query:  QueryWithSlices_Operation,
	}

	data_ = &QueryWithSlicesResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by SimpleQueryNoOverride.
const SimpleQueryNoOverride_Operation = `
query SimpleQueryNoOverride {
	user {
		id
		name
	}
}
`

func SimpleQueryNoOverride(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *SimpleQueryNoOverrideResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "SimpleQueryNoOverride",
		Query:  SimpleQueryNoOverride_Operation,
	}

	data_ = &SimpleQueryNoOverrideResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

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
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *SimpleQueryWithPointerFalseOverrideResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "SimpleQueryWithPointerFalseOverride",
		Query:  SimpleQueryWithPointerFalseOverride_Operation,
	}

	data_ = &SimpleQueryWithPointerFalseOverrideResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

