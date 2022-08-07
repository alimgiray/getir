package remote

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alimgiray/getir/internal/utils"
)

var (
	wrongParametersErr = errors.New("wrong parameters")
)

type RemoteHandler struct {
	service *RemoteService
}

func NewRemoteHandler() *RemoteHandler {
	return &RemoteHandler{
		service: NewRemoteService(),
	}
}

func (h *RemoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *RemoteHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	if request.StartDate == "" || request.EndDate == "" || request.MaxCount < request.MinCount {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	code, message, records := h.service.FindRecord(request)
	response := &Response{
		Code:    code,
		Message: message,
		Records: records,
	}
	utils.JSON(w, response, http.StatusOK)
}
