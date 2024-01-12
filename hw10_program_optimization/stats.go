package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	Email string
}

type DomainStat map[string]int

var p fastjson.Parser

func parseLine(line string, domain string, result *DomainStat) error {
	var user User
	v, err := p.Parse(line)
	if err != nil {
		return err
	}
	user = User{
		Email: string(v.GetStringBytes("Email")),
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
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if err := parseLine(scanner.Text(), domain, &result); err != nil {
			return result, err
		}
	}

	return result, nil
}
