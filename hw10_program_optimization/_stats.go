package hw10programoptimization

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/valyala/fastjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

var p fastjson.Parser

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func parseLine(line string, result *users) error {
	var user User
	v, err := p.Parse(line)
	if err != nil {
		return err
	}
	user = User{
		ID:       v.GetInt("Id"),
		Name:     string(v.GetStringBytes("Name")),
		Username: string(v.GetStringBytes("Username")),
		Email:    string(v.GetStringBytes("Email")),
		Phone:    string(v.GetStringBytes("Phone")),
		Password: string(v.GetStringBytes("Password")),
		Address:  string(v.GetStringBytes("Address")),
	}
	*result = append(*result, user)
	return nil
}

func getUsers(r io.Reader) (result users, err error) {
	var line strings.Builder
	var bytes int
	buf := make([]byte, 1)
	testing.Init()
	for {
		bytes, err = r.Read(buf)
		if bytes > 0 {
			line.WriteByte(buf[0])
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = parseLine(line.String(), &result)
				if err != nil {
					return
				}
				err = nil
				return
			}
			return
		}
		if buf[0] == '\n' {
			if err = parseLine(line.String(), &result); err != nil {
				return
			}
			line.Reset()
			continue
		}
	}
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		matched, err := regexp.Match("\\."+domain, []byte(user.Email))
		if err != nil {
			return nil, err
		}

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
