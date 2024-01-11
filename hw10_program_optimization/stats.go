package hw10programoptimization

import (
	"fmt"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	ID int
	// Name     string
	// Username string
	Email string
	// Phone    string
	// Password string
	// Address  string
}

type DomainStat map[string]int
type users []User

var p fastjson.Parser

func parseLine(line string, domain string, result *DomainStat) error {
	var user User
	v, err := p.Parse(line)
	if err != nil {
		return err
	}
	user = User{
		ID: v.GetInt("Id"),
		// Name:     string(v.GetStringBytes("Name")),
		// Username: string(v.GetStringBytes("Username")),
		Email: string(v.GetStringBytes("Email")),
		// Phone:    string(v.GetStringBytes("Phone")),
		// Password: string(v.GetStringBytes("Password")),
		// Address:  string(v.GetStringBytes("Address")),
	}
	if strings.HasSuffix(user.Email, "."+domain) {
		num := (*result)[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
		num++
		(*result)[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
	}

	return nil
}

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	userStat, err := getUsersStat(r, domain)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return userStat, nil
}

func getUsersStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	content, err := io.ReadAll(r)
	if err != nil {
		return result, err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if err = parseLine(line, domain, &result); err != nil {
			return result, err
		}
	}
	return result, err
}
