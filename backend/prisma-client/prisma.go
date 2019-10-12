// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4466"
var Secret = ""

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) UsersConnection(params *UsersConnectionParams) *UserConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"usersConnection",
		[]string{"edges", "pageInfo"})

	return &UserConnectionExec{ret}
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateManyMutationInput `json:"data"`
	Where *UserWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateManyMutationInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type UserOrderByInput string

const (
	UserOrderByInputIDAsc                   UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc                  UserOrderByInput = "id_DESC"
	UserOrderByInputSpotifyIdAsc            UserOrderByInput = "spotifyId_ASC"
	UserOrderByInputSpotifyIdDesc           UserOrderByInput = "spotifyId_DESC"
	UserOrderByInputSpotifyTokenAccessAsc   UserOrderByInput = "spotifyTokenAccess_ASC"
	UserOrderByInputSpotifyTokenAccessDesc  UserOrderByInput = "spotifyTokenAccess_DESC"
	UserOrderByInputSpotifyTokenRefreshAsc  UserOrderByInput = "spotifyTokenRefresh_ASC"
	UserOrderByInputSpotifyTokenRefreshDesc UserOrderByInput = "spotifyTokenRefresh_DESC"
	UserOrderByInputSpotifyTokenExpiryAsc   UserOrderByInput = "spotifyTokenExpiry_ASC"
	UserOrderByInputSpotifyTokenExpiryDesc  UserOrderByInput = "spotifyTokenExpiry_DESC"
	UserOrderByInputSpotifyTokenTypeAsc     UserOrderByInput = "spotifyTokenType_ASC"
	UserOrderByInputSpotifyTokenTypeDesc    UserOrderByInput = "spotifyTokenType_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserWhereUniqueInput struct {
	ID        *string `json:"id,omitempty"`
	SpotifyId *string `json:"spotifyId,omitempty"`
}

type UserWhereInput struct {
	ID                               *string          `json:"id,omitempty"`
	IDNot                            *string          `json:"id_not,omitempty"`
	IDIn                             []string         `json:"id_in,omitempty"`
	IDNotIn                          []string         `json:"id_not_in,omitempty"`
	IDLt                             *string          `json:"id_lt,omitempty"`
	IDLte                            *string          `json:"id_lte,omitempty"`
	IDGt                             *string          `json:"id_gt,omitempty"`
	IDGte                            *string          `json:"id_gte,omitempty"`
	IDContains                       *string          `json:"id_contains,omitempty"`
	IDNotContains                    *string          `json:"id_not_contains,omitempty"`
	IDStartsWith                     *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith                  *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith                       *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith                    *string          `json:"id_not_ends_with,omitempty"`
	SpotifyId                        *string          `json:"spotifyId,omitempty"`
	SpotifyIdNot                     *string          `json:"spotifyId_not,omitempty"`
	SpotifyIdIn                      []string         `json:"spotifyId_in,omitempty"`
	SpotifyIdNotIn                   []string         `json:"spotifyId_not_in,omitempty"`
	SpotifyIdLt                      *string          `json:"spotifyId_lt,omitempty"`
	SpotifyIdLte                     *string          `json:"spotifyId_lte,omitempty"`
	SpotifyIdGt                      *string          `json:"spotifyId_gt,omitempty"`
	SpotifyIdGte                     *string          `json:"spotifyId_gte,omitempty"`
	SpotifyIdContains                *string          `json:"spotifyId_contains,omitempty"`
	SpotifyIdNotContains             *string          `json:"spotifyId_not_contains,omitempty"`
	SpotifyIdStartsWith              *string          `json:"spotifyId_starts_with,omitempty"`
	SpotifyIdNotStartsWith           *string          `json:"spotifyId_not_starts_with,omitempty"`
	SpotifyIdEndsWith                *string          `json:"spotifyId_ends_with,omitempty"`
	SpotifyIdNotEndsWith             *string          `json:"spotifyId_not_ends_with,omitempty"`
	SpotifyTokenAccess               *string          `json:"spotifyTokenAccess,omitempty"`
	SpotifyTokenAccessNot            *string          `json:"spotifyTokenAccess_not,omitempty"`
	SpotifyTokenAccessIn             []string         `json:"spotifyTokenAccess_in,omitempty"`
	SpotifyTokenAccessNotIn          []string         `json:"spotifyTokenAccess_not_in,omitempty"`
	SpotifyTokenAccessLt             *string          `json:"spotifyTokenAccess_lt,omitempty"`
	SpotifyTokenAccessLte            *string          `json:"spotifyTokenAccess_lte,omitempty"`
	SpotifyTokenAccessGt             *string          `json:"spotifyTokenAccess_gt,omitempty"`
	SpotifyTokenAccessGte            *string          `json:"spotifyTokenAccess_gte,omitempty"`
	SpotifyTokenAccessContains       *string          `json:"spotifyTokenAccess_contains,omitempty"`
	SpotifyTokenAccessNotContains    *string          `json:"spotifyTokenAccess_not_contains,omitempty"`
	SpotifyTokenAccessStartsWith     *string          `json:"spotifyTokenAccess_starts_with,omitempty"`
	SpotifyTokenAccessNotStartsWith  *string          `json:"spotifyTokenAccess_not_starts_with,omitempty"`
	SpotifyTokenAccessEndsWith       *string          `json:"spotifyTokenAccess_ends_with,omitempty"`
	SpotifyTokenAccessNotEndsWith    *string          `json:"spotifyTokenAccess_not_ends_with,omitempty"`
	SpotifyTokenRefresh              *string          `json:"spotifyTokenRefresh,omitempty"`
	SpotifyTokenRefreshNot           *string          `json:"spotifyTokenRefresh_not,omitempty"`
	SpotifyTokenRefreshIn            []string         `json:"spotifyTokenRefresh_in,omitempty"`
	SpotifyTokenRefreshNotIn         []string         `json:"spotifyTokenRefresh_not_in,omitempty"`
	SpotifyTokenRefreshLt            *string          `json:"spotifyTokenRefresh_lt,omitempty"`
	SpotifyTokenRefreshLte           *string          `json:"spotifyTokenRefresh_lte,omitempty"`
	SpotifyTokenRefreshGt            *string          `json:"spotifyTokenRefresh_gt,omitempty"`
	SpotifyTokenRefreshGte           *string          `json:"spotifyTokenRefresh_gte,omitempty"`
	SpotifyTokenRefreshContains      *string          `json:"spotifyTokenRefresh_contains,omitempty"`
	SpotifyTokenRefreshNotContains   *string          `json:"spotifyTokenRefresh_not_contains,omitempty"`
	SpotifyTokenRefreshStartsWith    *string          `json:"spotifyTokenRefresh_starts_with,omitempty"`
	SpotifyTokenRefreshNotStartsWith *string          `json:"spotifyTokenRefresh_not_starts_with,omitempty"`
	SpotifyTokenRefreshEndsWith      *string          `json:"spotifyTokenRefresh_ends_with,omitempty"`
	SpotifyTokenRefreshNotEndsWith   *string          `json:"spotifyTokenRefresh_not_ends_with,omitempty"`
	SpotifyTokenExpiry               *string          `json:"spotifyTokenExpiry,omitempty"`
	SpotifyTokenExpiryNot            *string          `json:"spotifyTokenExpiry_not,omitempty"`
	SpotifyTokenExpiryIn             []string         `json:"spotifyTokenExpiry_in,omitempty"`
	SpotifyTokenExpiryNotIn          []string         `json:"spotifyTokenExpiry_not_in,omitempty"`
	SpotifyTokenExpiryLt             *string          `json:"spotifyTokenExpiry_lt,omitempty"`
	SpotifyTokenExpiryLte            *string          `json:"spotifyTokenExpiry_lte,omitempty"`
	SpotifyTokenExpiryGt             *string          `json:"spotifyTokenExpiry_gt,omitempty"`
	SpotifyTokenExpiryGte            *string          `json:"spotifyTokenExpiry_gte,omitempty"`
	SpotifyTokenExpiryContains       *string          `json:"spotifyTokenExpiry_contains,omitempty"`
	SpotifyTokenExpiryNotContains    *string          `json:"spotifyTokenExpiry_not_contains,omitempty"`
	SpotifyTokenExpiryStartsWith     *string          `json:"spotifyTokenExpiry_starts_with,omitempty"`
	SpotifyTokenExpiryNotStartsWith  *string          `json:"spotifyTokenExpiry_not_starts_with,omitempty"`
	SpotifyTokenExpiryEndsWith       *string          `json:"spotifyTokenExpiry_ends_with,omitempty"`
	SpotifyTokenExpiryNotEndsWith    *string          `json:"spotifyTokenExpiry_not_ends_with,omitempty"`
	SpotifyTokenType                 *string          `json:"spotifyTokenType,omitempty"`
	SpotifyTokenTypeNot              *string          `json:"spotifyTokenType_not,omitempty"`
	SpotifyTokenTypeIn               []string         `json:"spotifyTokenType_in,omitempty"`
	SpotifyTokenTypeNotIn            []string         `json:"spotifyTokenType_not_in,omitempty"`
	SpotifyTokenTypeLt               *string          `json:"spotifyTokenType_lt,omitempty"`
	SpotifyTokenTypeLte              *string          `json:"spotifyTokenType_lte,omitempty"`
	SpotifyTokenTypeGt               *string          `json:"spotifyTokenType_gt,omitempty"`
	SpotifyTokenTypeGte              *string          `json:"spotifyTokenType_gte,omitempty"`
	SpotifyTokenTypeContains         *string          `json:"spotifyTokenType_contains,omitempty"`
	SpotifyTokenTypeNotContains      *string          `json:"spotifyTokenType_not_contains,omitempty"`
	SpotifyTokenTypeStartsWith       *string          `json:"spotifyTokenType_starts_with,omitempty"`
	SpotifyTokenTypeNotStartsWith    *string          `json:"spotifyTokenType_not_starts_with,omitempty"`
	SpotifyTokenTypeEndsWith         *string          `json:"spotifyTokenType_ends_with,omitempty"`
	SpotifyTokenTypeNotEndsWith      *string          `json:"spotifyTokenType_not_ends_with,omitempty"`
	And                              []UserWhereInput `json:"AND,omitempty"`
	Or                               []UserWhereInput `json:"OR,omitempty"`
	Not                              []UserWhereInput `json:"NOT,omitempty"`
}

type UserCreateInput struct {
	ID                  *string `json:"id,omitempty"`
	SpotifyId           string  `json:"spotifyId"`
	SpotifyTokenAccess  string  `json:"spotifyTokenAccess"`
	SpotifyTokenRefresh string  `json:"spotifyTokenRefresh"`
	SpotifyTokenExpiry  string  `json:"spotifyTokenExpiry"`
	SpotifyTokenType    string  `json:"spotifyTokenType"`
}

type UserUpdateInput struct {
	SpotifyId           *string `json:"spotifyId,omitempty"`
	SpotifyTokenAccess  *string `json:"spotifyTokenAccess,omitempty"`
	SpotifyTokenRefresh *string `json:"spotifyTokenRefresh,omitempty"`
	SpotifyTokenExpiry  *string `json:"spotifyTokenExpiry,omitempty"`
	SpotifyTokenType    *string `json:"spotifyTokenType,omitempty"`
}

type UserUpdateManyMutationInput struct {
	SpotifyId           *string `json:"spotifyId,omitempty"`
	SpotifyTokenAccess  *string `json:"spotifyTokenAccess,omitempty"`
	SpotifyTokenRefresh *string `json:"spotifyTokenRefresh,omitempty"`
	SpotifyTokenExpiry  *string `json:"spotifyTokenExpiry,omitempty"`
	SpotifyTokenType    *string `json:"spotifyTokenType,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []UserSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []UserSubscriptionWhereInput `json:"NOT,omitempty"`
}

type UserExec struct {
	exec *prisma.Exec
}

func (instance UserExec) Exec(ctx context.Context) (*User, error) {
	var v User
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type User struct {
	ID                  string `json:"id"`
	SpotifyId           string `json:"spotifyId"`
	SpotifyTokenAccess  string `json:"spotifyTokenAccess"`
	SpotifyTokenRefresh string `json:"spotifyTokenRefresh"`
	SpotifyTokenExpiry  string `json:"spotifyTokenExpiry"`
	SpotifyTokenType    string `json:"spotifyTokenType"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"UserWhereInput", "UserOrderByInput", "UserEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetMany(
		edges,
		nil,
		[3]string{"", "", "User"},
		"node",
		[]string{"id", "createdAt", "updatedAt", "name", "desc"})

	return &UserEdgeExecArray{nodes}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (*UserConnection, error) {
	var v UserConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []UserEdge `json:"edges"`
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (*UserEdge, error) {
	var v UserEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserEdge struct {
	Node   User   `json:"node"`
	Cursor string `json:"cursor"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "spotifyId", "spotifyTokenAccess", "spotifyTokenRefresh", "spotifyTokenExpiry", "spotifyTokenType"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (*UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *User        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (*UserPreviousValues, error) {
	var v UserPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserPreviousValues struct {
	ID                  string `json:"id"`
	SpotifyId           string `json:"spotifyId"`
	SpotifyTokenAccess  string `json:"spotifyTokenAccess"`
	SpotifyTokenRefresh string `json:"spotifyTokenRefresh"`
	SpotifyTokenExpiry  string `json:"spotifyTokenExpiry"`
	SpotifyTokenType    string `json:"spotifyTokenType"`
}