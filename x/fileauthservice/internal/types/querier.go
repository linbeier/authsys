package types

import "strings"

// Query endpoints supported by the fileauthservice querier
//const (
// TODO: Describe query parameters, update <action> with your query
// Query<Action>    = "<action>"
//)

/*
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}
*/

// QueryResAuth Queries Result Payload for a resolve query
type QueryResAuth struct {
	Auth string `json:"auth"`
}

// implement fmt.Stringer
func (r QueryResAuth) String() string {
	return r.Auth
}

// QueryResNames Queries Result Payload for a names query
type QueryResFileNames []string

// implement fmt.Stringer
func (n QueryResFileNames) String() string {
	return strings.Join(n[:], "\n")
}

type QueryResAccounts []string

func (m QueryResAccounts) String() string {
	return strings.Join(m[:], "\n")
}
