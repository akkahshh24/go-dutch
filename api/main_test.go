package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// func newTestServer(t *testing.T, store *db.Store, tokenMaker token.Maker) *Server {
// 	config := util.Config{
// 		TokenSymmetricKey:   util.RandomString(32),
// 		AccessTokenDuration: time.Minute,
// 	}

// 	return NewServer(config, store, tokenMaker)
// }

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
