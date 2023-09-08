// Code generated by github.com/Interstellar-Lab/genqlient, DO NOT EDIT.

package queries

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Interstellar-Lab/genqlient/graphql"
	"github.com/Interstellar-Lab/genqlient/internal/testutil"
)

// InputObjectQueryResponse is returned by InputObjectQuery on success.
type InputObjectQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *InputObjectQueryUser `json:"user"`
}

// GetUser returns InputObjectQueryResponse.User, and is useful for accessing the field via an interface.
func (v *InputObjectQueryResponse) GetUser() *InputObjectQueryUser { return v.User }

// InputObjectQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type InputObjectQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `json:"id"`
}

// GetId returns InputObjectQueryUser.Id, and is useful for accessing the field via an interface.
func (v *InputObjectQueryUser) GetId() string { return v.Id }

type PokemonInput struct {
	Species string `json:"species"`
	Level   int    `json:"level"`
}

// GetSpecies returns PokemonInput.Species, and is useful for accessing the field via an interface.
func (v *PokemonInput) GetSpecies() string { return v.Species }

// GetLevel returns PokemonInput.Level, and is useful for accessing the field via an interface.
func (v *PokemonInput) GetLevel() int { return v.Level }

// QueryWithStructsResponse is returned by QueryWithStructs on success.
type QueryWithStructsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *QueryWithStructsUser `json:"user"`
}

// GetUser returns QueryWithStructsResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithStructsResponse) GetUser() *QueryWithStructsUser { return v.User }

// QueryWithStructsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithStructsUser struct {
	AuthMethods []*QueryWithStructsUserAuthMethodsAuthMethod `json:"authMethods"`
}

// GetAuthMethods returns QueryWithStructsUser.AuthMethods, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUser) GetAuthMethods() []*QueryWithStructsUserAuthMethodsAuthMethod {
	return v.AuthMethods
}

// QueryWithStructsUserAuthMethodsAuthMethod includes the requested fields of the GraphQL type AuthMethod.
type QueryWithStructsUserAuthMethodsAuthMethod struct {
	Provider *string `json:"provider"`
	Email    *string `json:"email"`
}

// GetProvider returns QueryWithStructsUserAuthMethodsAuthMethod.Provider, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUserAuthMethodsAuthMethod) GetProvider() *string { return v.Provider }

// GetEmail returns QueryWithStructsUserAuthMethodsAuthMethod.Email, and is useful for accessing the field via an interface.
func (v *QueryWithStructsUserAuthMethodsAuthMethod) GetEmail() *string { return v.Email }

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type UserQueryInput struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id         *string       `json:"id"`
	Role       *Role         `json:"role"`
	Names      []*string     `json:"names"`
	HasPokemon *PokemonInput `json:"hasPokemon,omitempty"`
	Birthdate  *time.Time    `json:"-"`
}

// GetEmail returns UserQueryInput.Email, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetEmail() *string { return v.Email }

// GetName returns UserQueryInput.Name, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetName() *string { return v.Name }

// GetId returns UserQueryInput.Id, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetId() *string { return v.Id }

// GetRole returns UserQueryInput.Role, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetRole() *Role { return v.Role }

// GetNames returns UserQueryInput.Names, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetNames() []*string { return v.Names }

// GetHasPokemon returns UserQueryInput.HasPokemon, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetHasPokemon() *PokemonInput { return v.HasPokemon }

// GetBirthdate returns UserQueryInput.Birthdate, and is useful for accessing the field via an interface.
func (v *UserQueryInput) GetBirthdate() *time.Time { return v.Birthdate }

func (v *UserQueryInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*UserQueryInput
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.UserQueryInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			*dst = new(time.Time)
			err = testutil.UnmarshalDate(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalUserQueryInput struct {
	Email *string `json:"email"`

	Name *string `json:"name"`

	Id *string `json:"id"`

	Role *Role `json:"role"`

	Names []*string `json:"names"`

	HasPokemon *PokemonInput `json:"hasPokemon,omitempty"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *UserQueryInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *UserQueryInput) __premarshalJSON() (*__premarshalUserQueryInput, error) {
	var retval __premarshalUserQueryInput

	retval.Email = v.Email
	retval.Name = v.Name
	retval.Id = v.Id
	retval.Role = v.Role
	retval.Names = v.Names
	retval.HasPokemon = v.HasPokemon
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		if src != nil {
			var err error
			*dst, err = testutil.MarshalDate(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to marshal UserQueryInput.Birthdate: %w", err)
			}
		}
	}
	return &retval, nil
}

// __InputObjectQueryInput is used internally by genqlient
type __InputObjectQueryInput struct {
	Query *UserQueryInput `json:"query,omitempty"`
}

// GetQuery returns __InputObjectQueryInput.Query, and is useful for accessing the field via an interface.
func (v *__InputObjectQueryInput) GetQuery() *UserQueryInput { return v.Query }

// The query, mutation or subscription executed by InputObjectQuery.
const InputObjectQuery_Operation = `
query InputObjectQuery ($query: UserQueryInput) {
	user(query: $query) {
		id
	}
}
`

func InputObjectQuery(
	ctx_ context.Context,
	client_ graphql.Client,
	query *UserQueryInput,
) (data_ *InputObjectQueryResponse, err error) {
	req_ := &graphql.Request{
		OpName: "InputObjectQuery",
		Query:  InputObjectQuery_Operation,
		Variables: &__InputObjectQueryInput{
			Query: query,
		},
	}
	var err_ error

	data_ = &InputObjectQueryResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query, mutation or subscription executed by QueryWithStructs.
const QueryWithStructs_Operation = `
query QueryWithStructs {
	user {
		authMethods {
			provider
			email
		}
	}
}
`

func QueryWithStructs(
	ctx_ context.Context,
	client_ graphql.Client,
) (data_ *QueryWithStructsResponse, err error) {
	req_ := &graphql.Request{
		OpName: "QueryWithStructs",
		Query:  QueryWithStructs_Operation,
	}
	var err_ error

	data_ = &QueryWithStructsResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

