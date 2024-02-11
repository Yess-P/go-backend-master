package api

import (
	"backend-master/token"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// authorizationHeaderKey 의 header 체크
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		// 길이가 0과 같으면 header 없으므로 에러 출력
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// authorizationHeader 의 field 가 2보다 작으면 유효하지 않은 header 로 error 출력
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// field 의 0번째 값을 소문자 전환 후 authorizationType 에 대입
		authorizationType := strings.ToLower(fields[0])
		// authorizationType 가 authorizationTypeBearer 와 일치하지 않으면 지원하지 않는 타입으로 에려 출력
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// field 의 2번째 값은 accessToken 이다.
		accessToken := fields[1]
		// payload 를 검증하는 부분
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// authorizationPayloadKey 인 key 와, payload 라는 value 지정
		ctx.Set(authorizationPayloadKey, payload)

		// 다음 handler 에게 전달
		ctx.Next()
	}
}
