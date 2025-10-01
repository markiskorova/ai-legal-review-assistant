package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/auth"
	dbase "github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/db"
)

type signupReq struct{ Email, Password string }
type loginReq struct{ Email, Password string }

func handleSignup(w http.ResponseWriter, r *http.Request) {
	var req signupReq
	json.NewDecoder(r.Body).Decode(&req)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	// naive password hash for MVP
	_, err := pool.Exec(ctx, `insert into "user"(email,password_hash,org_id,role) values($1,$2,1,'admin')`, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(201)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	json.NewDecoder(r.Body).Decode(&req)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()

	var id int64
	var hash string
	err := pool.QueryRow(ctx, `select id,password_hash from "user" where email=$1`, req.Email).Scan(&id, &hash)
	if err == pgx.ErrNoRows || hash != req.Password {
		http.Error(w, "invalid creds", 401)
		return
	}
	tok, _ := auth.Sign(id)
	json.NewEncoder(w).Encode(map[string]string{"token": tok})
}

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authz := r.Header.Get("Authorization")
			if len(authz) < 8 {
				http.Error(w, "no token", 401)
				return
			}
			tokenStr := authz[len("Bearer "):]
			t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !t.Valid {
				http.Error(w, "bad token", 401)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func handleMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ok":true}`))
}
