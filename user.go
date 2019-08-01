// +build windows

package winproc

import (
	"syscall"

	"github.com/gentlemanautomaton/winproc/winsecid"
)

// User holds account information for the security context of a process.
type User struct {
	SID     string
	Account string
	Domain  string
	Type    uint32
}

// System returns true if u describes a system user with one of the following
// security identifiers:
//
//	Local System
//	NT Authority
//	Network Service
//
func (u User) System() bool {
	switch u.SID {
	case winsecid.LocalSystem, winsecid.NTAuthority, winsecid.NetworkService:
		return true
	default:
		return false
	}
}

// String returns a string representation of the user.
func (u User) String() string {
	if u.Account == "" {
		return u.SID
	}
	if u.Domain == "" {
		return u.Account
	}
	return u.Domain + `\` + u.Account
}

func userFromProcess(process syscall.Handle) (User, error) {
	var token syscall.Token
	if err := syscall.OpenProcessToken(process, syscall.TOKEN_QUERY, &token); err != nil {
		return User{}, err
	}
	defer token.Close()

	tokenUser, err := token.GetTokenUser()
	if err != nil {
		return User{}, err
	}

	sid, err := tokenUser.User.Sid.String()
	if err != nil {
		return User{}, err
	}

	account, domain, accType, err := tokenUser.User.Sid.LookupAccount("")
	if err != nil {
		return User{}, err
	}

	return User{
		SID:     sid,
		Account: account,
		Domain:  domain,
		Type:    accType,
	}, nil
}
