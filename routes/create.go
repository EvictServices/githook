package routes

import (
	"fmt"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shi-gg/githook/config"
	"github.com/shi-gg/githook/utils"
)

func normalizeKey(secret string) []byte {
	if len(secret) < 32 {
		secret = fmt.Sprintf("%-32s", secret)
	} else if len(secret) > 32 {
		secret = secret[:32]
	}
	return []byte(secret)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	conf := config.Get()

	url := r.URL.Query().Get("url")

	key := normalizeKey(conf.Secret)

	id, err := utils.Encrypt(url, key)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}
