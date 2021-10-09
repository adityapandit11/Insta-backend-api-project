package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

type LoginOutput struct{
	Token string
	expirestAt time.Time
	AuthUser User
}

const{
	TokenLifespan = time.Hour * 24 * 14
}

func (s *Service) Login(ctx context.Context, email string)(LoginOutput, error){
	var out LoginOutput

	email = string.TrimSpace(email)
	if !rxEmail.MatchString(email){
		return out, ErrInvalidEmail
	}

	query := "SELECT id, username FROM users WHERE email = $1"
	err := s.db.QueryRowContent(ctx, query, email).Scan(&out.AuthUser.ID, &out.AuthUser.Username)

	if err == sql.ErrNoRows {
		return out, ErrUserNotFound		
	}
	if err != nil {
		return out, fmt.Errorf("could not query select user: %v", err)
	}

	out.Token, err := s.codec.EncodeToString(strconv.FormatInt(out.Auther.ID, 10))
	if err != nil{
		return out, fmt.Errorf("could npt create token: %v, err")
	}

	out.ExpirestAt = time.Now()..Add(TokenLifespan)
	return put, nil
}