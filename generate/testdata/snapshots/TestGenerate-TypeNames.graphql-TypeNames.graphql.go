// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Interstellar-Lab/genqlient/graphql"
	"github.com/Interstellar-Lab/genqlient/internal/testutil"
)

// Item includes the requested fields of the GraphQL interface Content.
//
// Item is implemented by the following types:
// ItemArticle
// ItemTopic
// ItemVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type Item interface {
	implementsGraphQLInterfaceItem()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() NameType
}

func (v *ItemArticle) implementsGraphQLInterfaceItem() {}
func (v *ItemTopic) implementsGraphQLInterfaceItem()   {}
func (v *ItemVideo) implementsGraphQLInterfaceItem()   {}

func __unmarshalItem(b []byte, v *Item) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(ItemArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(ItemTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(ItemVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for Item: "%v"`, tn.TypeName)
	}
}

func __marshalItem(v *Item) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *ItemArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*ItemArticle
		}{typename, v}
		return json.Marshal(result)
	case *ItemTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*ItemTopic
		}{typename, v}
		return json.Marshal(result)
	case *ItemVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*ItemVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for Item: "%T"`, v)
	}
}

// ItemArticle includes the requested fields of the GraphQL type Article.
type ItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name NameType    `json:"name"`
}

// GetTypename returns ItemArticle.Typename, and is useful for accessing the field via an interface.
func (v *ItemArticle) GetTypename() string { return v.Typename }

// GetId returns ItemArticle.Id, and is useful for accessing the field via an interface.
func (v *ItemArticle) GetId() testutil.ID { return v.Id }

// GetName returns ItemArticle.Name, and is useful for accessing the field via an interface.
func (v *ItemArticle) GetName() NameType { return v.Name }

// ItemTopic includes the requested fields of the GraphQL type Topic.
type ItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name NameType    `json:"name"`
}

// GetTypename returns ItemTopic.Typename, and is useful for accessing the field via an interface.
func (v *ItemTopic) GetTypename() string { return v.Typename }

// GetId returns ItemTopic.Id, and is useful for accessing the field via an interface.
func (v *ItemTopic) GetId() testutil.ID { return v.Id }

// GetName returns ItemTopic.Name, and is useful for accessing the field via an interface.
func (v *ItemTopic) GetName() NameType { return v.Name }

// ItemVideo includes the requested fields of the GraphQL type Video.
type ItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name NameType    `json:"name"`
}

// GetTypename returns ItemVideo.Typename, and is useful for accessing the field via an interface.
func (v *ItemVideo) GetTypename() string { return v.Typename }

// GetId returns ItemVideo.Id, and is useful for accessing the field via an interface.
func (v *ItemVideo) GetId() testutil.ID { return v.Id }

// GetName returns ItemVideo.Name, and is useful for accessing the field via an interface.
func (v *ItemVideo) GetName() NameType { return v.Name }

type NameType string

// Resp is returned by TypeNames on success.
type Resp struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User       User   `json:"user"`
	RandomItem Item   `json:"-"`
	Users      []User `json:"users"`
}

// GetUser returns Resp.User, and is useful for accessing the field via an interface.
func (v *Resp) GetUser() User { return v.User }

// GetRandomItem returns Resp.RandomItem, and is useful for accessing the field via an interface.
func (v *Resp) GetRandomItem() Item { return v.RandomItem }

// GetUsers returns Resp.Users, and is useful for accessing the field via an interface.
func (v *Resp) GetUsers() []User { return v.Users }

func (v *Resp) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*Resp
		RandomItem json.RawMessage `json:"randomItem"`
		graphql.NoUnmarshalJSON
	}
	firstPass.Resp = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalItem(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal Resp.RandomItem: %w", err)
			}
		}
	}
	return nil
}

type __premarshalResp struct {
	User User `json:"user"`

	RandomItem json.RawMessage `json:"randomItem"`

	Users []User `json:"users"`
}

func (v *Resp) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *Resp) __premarshalJSON() (*__premarshalResp, error) {
	var retval __premarshalResp

	retval.User = v.User
	{

		dst := &retval.RandomItem
		src := v.RandomItem
		var err error
		*dst, err = __marshalItem(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal Resp.RandomItem: %w", err)
		}
	}
	retval.Users = v.Users
	return &retval, nil
}

// User includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type User struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns User.Id, and is useful for accessing the field via an interface.
func (v *User) GetId() testutil.ID { return v.Id }

// GetName returns User.Name, and is useful for accessing the field via an interface.
func (v *User) GetName() string { return v.Name }

// The query, mutation or subscription executed by TypeNames.
const TypeNames_Operation = `
query TypeNames {
	user {
		id
		name
	}
	randomItem {
		__typename
		id
		name
	}
	users {
		id
		name
	}
}
`

func TypeNames(
	client_ graphql.Client,
) (data_ *Resp, err error) {
	req_ := &graphql.Request{
		OpName: "TypeNames",
		Query:  TypeNames_Operation,
	}
	var err_ error

	data_ = &Resp{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

