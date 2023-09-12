// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContent is implemented by the following types:
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic
// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldListOfListsOfListsOfContent() {
}

func __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(b []byte, v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContent) error {
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
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceListOfListOfListsFieldListOfListsOfListsOfContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceListOfListOfListsFieldListOfListsOfListsOfContent: "%T"`, v)
	}
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetId() testutil.ID {
	return v.Id
}

// GetName returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentArticle) GetName() string {
	return v.Name
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetId() testutil.ID {
	return v.Id
}

// GetName returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentTopic) GetName() string {
	return v.Name
}

// InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetTypename() string {
	return v.Typename
}

// GetId returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetId() testutil.ID {
	return v.Id
}

// GetName returns InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldListOfListsOfListsOfContentVideo) GetName() string {
	return v.Name
}

// InterfaceListOfListOfListsFieldResponse is returned by InterfaceListOfListOfListsField on success.
type InterfaceListOfListOfListsFieldResponse struct {
	ListOfListsOfListsOfContent [][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent `json:"-"`
	WithPointer                 [][][]*InterfaceListOfListOfListsFieldWithPointerContent         `json:"-"`
}

// GetListOfListsOfListsOfContent returns InterfaceListOfListOfListsFieldResponse.ListOfListsOfListsOfContent, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldResponse) GetListOfListsOfListsOfContent() [][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent {
	return v.ListOfListsOfListsOfContent
}

// GetWithPointer returns InterfaceListOfListOfListsFieldResponse.WithPointer, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldResponse) GetWithPointer() [][][]*InterfaceListOfListOfListsFieldWithPointerContent {
	return v.WithPointer
}

func (v *InterfaceListOfListOfListsFieldResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*InterfaceListOfListOfListsFieldResponse
		ListOfListsOfListsOfContent [][][]json.RawMessage `json:"listOfListsOfListsOfContent"`
		WithPointer                 [][][]json.RawMessage `json:"withPointer"`
		graphql.NoUnmarshalJSON
	}
	firstPass.InterfaceListOfListOfListsFieldResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.ListOfListsOfListsOfContent
		src := firstPass.ListOfListsOfListsOfContent
		*dst = make(
			[][][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]InterfaceListOfListOfListsFieldListOfListsOfListsOfContent,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if len(src) != 0 && string(src) != "null" {
						err = __unmarshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(
							src, dst)
						if err != nil {
							return fmt.Errorf(
								"unable to unmarshal InterfaceListOfListOfListsFieldResponse.ListOfListsOfListsOfContent: %w", err)
						}
					}
				}
			}
		}
	}

	{
		dst := &v.WithPointer
		src := firstPass.WithPointer
		*dst = make(
			[][][]*InterfaceListOfListOfListsFieldWithPointerContent,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]*InterfaceListOfListOfListsFieldWithPointerContent,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]*InterfaceListOfListOfListsFieldWithPointerContent,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if len(src) != 0 && string(src) != "null" {
						*dst = new(InterfaceListOfListOfListsFieldWithPointerContent)
						err = __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(
							src, *dst)
						if err != nil {
							return fmt.Errorf(
								"unable to unmarshal InterfaceListOfListOfListsFieldResponse.WithPointer: %w", err)
						}
					}
				}
			}
		}
	}
	return nil
}

type __premarshalInterfaceListOfListOfListsFieldResponse struct {
	ListOfListsOfListsOfContent [][][]json.RawMessage `json:"listOfListsOfListsOfContent"`

	WithPointer [][][]json.RawMessage `json:"withPointer"`
}

func (v *InterfaceListOfListOfListsFieldResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *InterfaceListOfListOfListsFieldResponse) __premarshalJSON() (*__premarshalInterfaceListOfListOfListsFieldResponse, error) {
	var retval __premarshalInterfaceListOfListOfListsFieldResponse

	{

		dst := &retval.ListOfListsOfListsOfContent
		src := v.ListOfListsOfListsOfContent
		*dst = make(
			[][][]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]json.RawMessage,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]json.RawMessage,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					var err error
					*dst, err = __marshalInterfaceListOfListOfListsFieldListOfListsOfListsOfContent(
						&src)
					if err != nil {
						return nil, fmt.Errorf(
							"unable to marshal InterfaceListOfListOfListsFieldResponse.ListOfListsOfListsOfContent: %w", err)
					}
				}
			}
		}
	}
	{

		dst := &retval.WithPointer
		src := v.WithPointer
		*dst = make(
			[][][]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			*dst = make(
				[][]json.RawMessage,
				len(src))
			for i, src := range src {
				dst := &(*dst)[i]
				*dst = make(
					[]json.RawMessage,
					len(src))
				for i, src := range src {
					dst := &(*dst)[i]
					if src != nil {
						var err error
						*dst, err = __marshalInterfaceListOfListOfListsFieldWithPointerContent(
							src)
						if err != nil {
							return nil, fmt.Errorf(
								"unable to marshal InterfaceListOfListOfListsFieldResponse.WithPointer: %w", err)
						}
					}
				}
			}
		}
	}
	return &retval, nil
}

// InterfaceListOfListOfListsFieldWithPointerArticle includes the requested fields of the GraphQL type Article.
type InterfaceListOfListOfListsFieldWithPointerArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldWithPointerArticle.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetTypename() string { return v.Typename }

// GetId returns InterfaceListOfListOfListsFieldWithPointerArticle.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceListOfListOfListsFieldWithPointerArticle.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerArticle) GetName() *string { return v.Name }

// InterfaceListOfListOfListsFieldWithPointerContent includes the requested fields of the GraphQL interface Content.
//
// InterfaceListOfListOfListsFieldWithPointerContent is implemented by the following types:
// InterfaceListOfListOfListsFieldWithPointerArticle
// InterfaceListOfListOfListsFieldWithPointerTopic
// InterfaceListOfListOfListsFieldWithPointerVideo
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceListOfListOfListsFieldWithPointerContent interface {
	implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() *testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() *string
}

func (v *InterfaceListOfListOfListsFieldWithPointerArticle) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) implementsGraphQLInterfaceInterfaceListOfListOfListsFieldWithPointerContent() {
}

func __unmarshalInterfaceListOfListOfListsFieldWithPointerContent(b []byte, v *InterfaceListOfListOfListsFieldWithPointerContent) error {
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
		*v = new(InterfaceListOfListOfListsFieldWithPointerArticle)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(InterfaceListOfListOfListsFieldWithPointerTopic)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(InterfaceListOfListOfListsFieldWithPointerVideo)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for InterfaceListOfListOfListsFieldWithPointerContent: "%v"`, tn.TypeName)
	}
}

func __marshalInterfaceListOfListOfListsFieldWithPointerContent(v *InterfaceListOfListOfListsFieldWithPointerContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *InterfaceListOfListOfListsFieldWithPointerArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldWithPointerArticle
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListOfListOfListsFieldWithPointerTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldWithPointerTopic
		}{typename, v}
		return json.Marshal(result)
	case *InterfaceListOfListOfListsFieldWithPointerVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*InterfaceListOfListOfListsFieldWithPointerVideo
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for InterfaceListOfListOfListsFieldWithPointerContent: "%T"`, v)
	}
}

// InterfaceListOfListOfListsFieldWithPointerTopic includes the requested fields of the GraphQL type Topic.
type InterfaceListOfListOfListsFieldWithPointerTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldWithPointerTopic.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetTypename() string { return v.Typename }

// GetId returns InterfaceListOfListOfListsFieldWithPointerTopic.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceListOfListOfListsFieldWithPointerTopic.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerTopic) GetName() *string { return v.Name }

// InterfaceListOfListOfListsFieldWithPointerVideo includes the requested fields of the GraphQL type Video.
type InterfaceListOfListOfListsFieldWithPointerVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   *testutil.ID `json:"id"`
	Name *string      `json:"name"`
}

// GetTypename returns InterfaceListOfListOfListsFieldWithPointerVideo.Typename, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetTypename() string { return v.Typename }

// GetId returns InterfaceListOfListOfListsFieldWithPointerVideo.Id, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetId() *testutil.ID { return v.Id }

// GetName returns InterfaceListOfListOfListsFieldWithPointerVideo.Name, and is useful for accessing the field via an interface.
func (v *InterfaceListOfListOfListsFieldWithPointerVideo) GetName() *string { return v.Name }

// The query executed by InterfaceListOfListOfListsField.
const InterfaceListOfListOfListsField_Operation = `
query InterfaceListOfListOfListsField {
	listOfListsOfListsOfContent {
		__typename
		id
		name
	}
	withPointer: listOfListsOfListsOfContent {
		__typename
		id
		name
	}
}
`

func InterfaceListOfListOfListsField(
	client_ graphql.Client,
) (data_ *InterfaceListOfListOfListsFieldResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "InterfaceListOfListOfListsField",
		Query:  InterfaceListOfListOfListsField_Operation,
	}

	data_ = &InterfaceListOfListOfListsFieldResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

