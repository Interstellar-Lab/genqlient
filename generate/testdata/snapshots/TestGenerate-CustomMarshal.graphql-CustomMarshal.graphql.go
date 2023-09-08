// Code generated by github.com/Interstellar-Lab/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Interstellar-Lab/genqlient/graphql"
	"github.com/Interstellar-Lab/genqlient/internal/testutil"
)

// CustomMarshalResponse is returned by CustomMarshal on success.
type CustomMarshalResponse struct {
	UsersBornOn []CustomMarshalUsersBornOnUser `json:"usersBornOn"`
}

// GetUsersBornOn returns CustomMarshalResponse.UsersBornOn, and is useful for accessing the field via an interface.
func (v *CustomMarshalResponse) GetUsersBornOn() []CustomMarshalUsersBornOnUser { return v.UsersBornOn }

// CustomMarshalUsersBornOnUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type CustomMarshalUsersBornOnUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id        testutil.ID `json:"id"`
	Birthdate time.Time   `json:"-"`
}

// GetId returns CustomMarshalUsersBornOnUser.Id, and is useful for accessing the field via an interface.
func (v *CustomMarshalUsersBornOnUser) GetId() testutil.ID { return v.Id }

// GetBirthdate returns CustomMarshalUsersBornOnUser.Birthdate, and is useful for accessing the field via an interface.
func (v *CustomMarshalUsersBornOnUser) GetBirthdate() time.Time { return v.Birthdate }

func (v *CustomMarshalUsersBornOnUser) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomMarshalUsersBornOnUser
		Birthdate json.RawMessage `json:"birthdate"`
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomMarshalUsersBornOnUser = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Birthdate
		src := firstPass.Birthdate
		if len(src) != 0 && string(src) != "null" {
			err = testutil.UnmarshalDate(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal CustomMarshalUsersBornOnUser.Birthdate: %w", err)
			}
		}
	}
	return nil
}

type __premarshalCustomMarshalUsersBornOnUser struct {
	Id testutil.ID `json:"id"`

	Birthdate json.RawMessage `json:"birthdate"`
}

func (v *CustomMarshalUsersBornOnUser) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomMarshalUsersBornOnUser) __premarshalJSON() (*__premarshalCustomMarshalUsersBornOnUser, error) {
	var retval __premarshalCustomMarshalUsersBornOnUser

	retval.Id = v.Id
	{

		dst := &retval.Birthdate
		src := v.Birthdate
		var err error
		*dst, err = testutil.MarshalDate(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal CustomMarshalUsersBornOnUser.Birthdate: %w", err)
		}
	}
	return &retval, nil
}

// __CustomMarshalInput is used internally by genqlient
type __CustomMarshalInput struct {
	Date time.Time `json:"-"`
}

// GetDate returns __CustomMarshalInput.Date, and is useful for accessing the field via an interface.
func (v *__CustomMarshalInput) GetDate() time.Time { return v.Date }

func (v *__CustomMarshalInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*__CustomMarshalInput
		Date json.RawMessage `json:"date"`
		graphql.NoUnmarshalJSON
	}
	firstPass.__CustomMarshalInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Date
		src := firstPass.Date
		if len(src) != 0 && string(src) != "null" {
			err = testutil.UnmarshalDate(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal __CustomMarshalInput.Date: %w", err)
			}
		}
	}
	return nil
}

type __premarshal__CustomMarshalInput struct {
	Date json.RawMessage `json:"date"`
}

func (v *__CustomMarshalInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *__CustomMarshalInput) __premarshalJSON() (*__premarshal__CustomMarshalInput, error) {
	var retval __premarshal__CustomMarshalInput

	{

		dst := &retval.Date
		src := v.Date
		var err error
		*dst, err = testutil.MarshalDate(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal __CustomMarshalInput.Date: %w", err)
		}
	}
	return &retval, nil
}

// The query, mutation or subscription executed by CustomMarshal.
const CustomMarshal_Operation = `
query CustomMarshal ($date: Date!) {
	usersBornOn(date: $date) {
		id
		birthdate
	}
}
`

func CustomMarshal(
	client_ graphql.Client,
	date time.Time,
) (data_ *CustomMarshalResponse, err error) {
	req_ := &graphql.Request{
		OpName: "CustomMarshal",
		Query:  CustomMarshal_Operation,
		Variables: &__CustomMarshalInput{
			Date: date,
		},
	}
	var err_ error

	data_ = &CustomMarshalResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

