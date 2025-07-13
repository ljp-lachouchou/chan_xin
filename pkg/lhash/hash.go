package lhash

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func GenPasswordHash(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true

}
func ValidatePassword(password string) error {
	// 条件1: 长度至少8字符
	if len(password) < 8 {
		return errors.WithStack(errors.New("密码长度需至少8个字符"))
	}
	if len(password) > 17 {
		return errors.WithStack(errors.New("密码长度需至多16个字符"))
	}
	// 条件2: 包含小写字母
	lowerRegex := regexp.MustCompile(`[a-z]`)
	if !lowerRegex.MatchString(password) {
		return errors.WithStack(errors.New("密码需包含至少1个小写字母"))
	}

	// 条件3: 包含大写字母
	upperRegex := regexp.MustCompile(`[A-Z]`)
	if !upperRegex.MatchString(password) {

		return errors.WithStack(errors.New("密码需包含至少1个大写字母"))
	}

	// 条件4: 包含数字
	digitRegex := regexp.MustCompile(`[0-9]`)
	if !digitRegex.MatchString(password) {
		return errors.WithStack(errors.New("密码需包含至少1个数字"))
	}

	// 条件5: 包含特殊字符
	specialRegex := regexp.MustCompile(`[!@#$%^&*]`)
	if !specialRegex.MatchString(password) {
		return errors.WithStack(errors.New("密码需包含至少1个特殊字符(!@#$%^&*)"))
	}
	return nil
}
