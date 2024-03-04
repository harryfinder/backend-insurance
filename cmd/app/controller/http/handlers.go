package http

import (
	"encoding/json"
	"github.com/harryfinder/backend-Insurance/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	resp models.Response
)

func (s *server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	resp.Send(w, models.SUCCESS, "Ping Pong")
	return
}

func (s *server) Insurance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var Vehicle models.VehicleInsurance

	if err := json.NewDecoder(r.Body).Decode(&Vehicle); err != nil {
		resp.Send(w, models.BADREQUEST, "Неправильный формат данных")
		return
	}
	defer r.Body.Close()

	if err := s.usecase.UInsurance(r.Context(), Vehicle); err != nil {
		resp.Send(w, models.BADREQUEST, err)
		return
	}

	resp.Send(w, models.SUCCESS, "Все успешно обработано")
	return

}
