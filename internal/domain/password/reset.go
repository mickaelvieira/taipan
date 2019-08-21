package password

import (
	"time"

	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/nonce"
)

// ResetToken reset password entry
type ResetToken struct {
	Token     string
	IsUsed    bool
	UserID    string
	CreatedAt time.Time
	ExpiredAt time.Time
	UsedAt    time.Time
}

// IsExpired is the roken expired?
func (r *ResetToken) IsExpired() bool {
	now := time.Now().UTC()
	return now.After(r.ExpiredAt)
}

const tokenNonceLen = 30
const tokenValidityInMinutes = 20

// getExpireAt calculates the date when the token will be expired
func getExpireAt() time.Time {
	start := time.Now().UTC()
	return start.Add(time.Minute * tokenValidityInMinutes)
}

// NewResetPasswordToken create a new password reset entry
func NewResetPasswordToken(u *user.User) *ResetToken {
	return &ResetToken{
		Token:     nonce.Generate(tokenNonceLen),
		UserID:    u.ID,
		CreatedAt: time.Now(),
		ExpiredAt: getExpireAt(),
	}
}
