package middleware

import (
	"GopherTok/common/consts"
	"GopherTok/common/errorx"
	"GopherTok/common/utils"
	"GopherTok/server/relation/api/internal/config"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type JwtMiddleware struct {
	Config config.Config
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config: c,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		token := r.FormValue("token")
		if token == "" {
			token = r.PostFormValue("token")
			if token == "" {
				httpx.WriteJson(w, http.StatusUnauthorized, errorx.NewCodeError(30004, errorx.ErrTokenEmpty))
				return
			}
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 {
			httpx.WriteJson(w, http.StatusUnauthorized, errorx.NewCodeError(30002, errorx.ErrHeadFormat))
			return
		}
		parseToken, isExpire, err := utils.ParseToken(parts[0], parts[1], m.Config.Token.AccessToken, m.Config.Token.RefreshToken)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, errorx.NewCodeError(30003, errorx.ErrTokenProve))
			return
		}
		if isExpire {
			parts[0], parts[1] = utils.GetToken(parseToken.ID, parseToken.State, m.Config.Token.AccessToken, m.Config.Token.RefreshToken)
			//w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[0], parts[1]))

		}
		token = parts[0] + " " + parts[1]
		r = r.WithContext(context.WithValue(r.Context(), consts.UserId, parseToken.ID))
		r = r.WithContext(context.WithValue(r.Context(), consts.Token, token))
		next(w, r)
	}
}
