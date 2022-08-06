package memory

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alimgiray/getir/adapter/sqlite"
	"github.com/alimgiray/getir/internal/utils"

	"gorm.io/gorm"
)

var (
	keyNotFoundErr     = errors.New("key not found")
	valueNotFoundErr   = errors.New("value not found")
	wrongParametersErr = errors.New("wrong parameters")
	saveErr            = errors.New("couldn't save record")
)

type InMemoryHandler struct {
	db *gorm.DB
}

func NewInMemoryHandler() *InMemoryHandler {
	return &InMemoryHandler{db: sqlite.Connect(&Record{})}
}

func (h *InMemoryHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.handleGet(w, r)
	} else if r.Method == http.MethodPost {
		h.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *InMemoryHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		utils.ErrJSON(w, keyNotFoundErr, http.StatusBadRequest)
		return
	}

	record := Record{Key: key}
	err := h.db.Find(&record).Error

	if err != nil {
		utils.ErrJSON(w, valueNotFoundErr, http.StatusBadRequest)
		return
	}

	utils.JSON(w, record, http.StatusOK)
}

func (h *InMemoryHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var record Record

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	if record.Key == "" || record.Value == "" {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	err = h.db.Save(&record).Error
	if err != nil {
		utils.ErrJSON(w, saveErr, http.StatusBadRequest)
		return
	}

	utils.JSON(w, record, http.StatusOK)
}
