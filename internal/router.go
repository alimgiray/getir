package getir

import (
	"net/http"

	"github.com/alimgiray/getir/internal/memory"
	"github.com/alimgiray/getir/internal/remote"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/remote", remote.NewRemoteHandler().Handle)
	mux.HandleFunc("/in-memory", memory.NewInMemoryHandler().Handle)
	return mux
}
