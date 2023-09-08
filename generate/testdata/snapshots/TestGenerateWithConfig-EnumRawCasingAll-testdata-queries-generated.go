// Code generated by github.com/Interstellar-Lab/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/Interstellar-Lab/genqlient/graphql"
)

// QueryWithEnumsOtherUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithEnumsOtherUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns QueryWithEnumsOtherUser.Roles, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsOtherUser) GetRoles() []Role { return v.Roles }

// QueryWithEnumsResponse is returned by QueryWithEnums on success.
type QueryWithEnumsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithEnumsUser `json:"user"`
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	OtherUser QueryWithEnumsOtherUser `json:"otherUser"`
}

// GetUser returns QueryWithEnumsResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsResponse) GetUser() QueryWithEnumsUser { return v.User }

// GetOtherUser returns QueryWithEnumsResponse.OtherUser, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsResponse) GetOtherUser() QueryWithEnumsOtherUser { return v.OtherUser }

// QueryWithEnumsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithEnumsUser struct {
	Roles []Role `json:"roles"`
}

// GetRoles returns QueryWithEnumsUser.Roles, and is useful for accessing the field via an interface.
func (v *QueryWithEnumsUser) GetRoles() []Role { return v.Roles }

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	Role_STUDENT Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	Role_TEACHER Role = "TEACHER"
)

// The query, mutation or subscription executed by QueryWithEnums.
const QueryWithEnums_Operation = `
query QueryWithEnums {
	user {
		roles
	}
	otherUser: user {
		roles
	}
}
`

func QueryWithEnums(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *QueryWithEnumsResponse, err error) {
	req_ := &graphql.Request{
		OpName: "QueryWithEnums",
		Query:  QueryWithEnums_Operation,
	}
	var err_ error

	data_ = &QueryWithEnumsResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

