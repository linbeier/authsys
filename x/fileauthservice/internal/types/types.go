package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinTransPrice is Initial Starting Price for a file authority to transfer
var MinTransPrice = sdk.Coins{sdk.NewInt64Coin("token", 1)}

//authority set is implement by letters, 'r' for read, 'w' for write, 't' for transfer, they are combined in a certain sequnce. For example, 'rwt' represents the account have read, write and transfer authorities.
type Fileauth struct {
	Name   string         `json:"name"`
	Hash   string         `json:"hash"`
	Owner  sdk.AccAddress `json:"owner"`
	Origin sdk.AccAddress `json:"origin"`
	Auth   string         `json:"auth"`
}

// NewFileauth returns a new Fileauth with the minprice as the price
func NewFileauth() Fileauth {
	return Fileauth{}
}

// implement fmt.Stringer
func (f Fileauth) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s
Owner: %s
Origin: %s
Hash: %s
Auth: %s`, f.Name, f.Owner, f.Origin, f.Hash, f.Auth))
}

type Filerecord struct {
	Name   string         `json:"name"`
	Hash   string         `json:"hash"`
	From   sdk.AccAddress `json:"from"`
	Origin sdk.AccAddress `json:"origin"`
	Time   time.Time      `json:"time"`
}

func NewFilerecord() Filerecord {
	return Filerecord{}
}

func (f Filerecord) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s
	From: %s
	Time: %s
	Hash: %s`, f.Name, f.From, f.Time, f.Hash))
}
