package password

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"github/mickaelvieira/taipan/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
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

// returns a random integer between 2 values
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// returns the byte representation of a random minuscule, majuscule, number or special character
func randomSymbol() byte {
	return byte(randomInt(33, 126))
}

// generate a random of the length provided
func generateNonce(len int) string {
	// https://golang.org/pkg/math/rand/#Rand.Seed
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = randomSymbol()
	}

	// we use sha256 algorithm here to get:
	// - a fix length string
	// - a URL safe string
	buf := sha256.New()
	buf.Write(bytes)
	b := buf.Sum(nil)

	return hex.EncodeToString(b)
}

// getExpireAt calculates the date when the token will be expired
func getExpireAt() time.Time {
	start := time.Now().UTC()
	return start.Add(time.Minute * tokenValidityInMinutes)
}

// NewResetPasswordToken create a new password reset entry
func NewResetPasswordToken(u *user.User) *ResetToken {
	return &ResetToken{
		Token:     generateNonce(tokenNonceLen),
		UserID:    u.ID,
		CreatedAt: time.Now(),
		ExpiredAt: getExpireAt(),
	}
}

// Hash generates a hash for the given password
func Hash(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(h), nil
}
