// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errors

import "fmt"

type TwoFactorNotFound struct {
	UserID int64
}

func IsTwoFactorNotFound(err error) bool {
	_, ok := err.(TwoFactorNotFound)
	return ok
}

func (err TwoFactorNotFound) Error() string {
	return fmt.Sprintf("two-factor authentication does not found [user_id: %d]", err.UserID)
}
