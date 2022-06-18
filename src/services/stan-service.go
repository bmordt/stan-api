package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bmordt/stan-api/src/middleware"
	"github.com/bmordt/stan-api/src/models"
)

var (
	apiError middleware.ApiError
)

type StanService struct {
	Logger *logrus.Entry
}

func NewStanService(logger *logrus.Entry) *StanService {
	return &StanService{
		Logger: logger,
	}
}

func (s *StanService) FilterStanJson(w http.ResponseWriter, r *http.Request) {
	s.Logger.Infof("Inside FilterStanJson")

	stanReq := &models.StanRequest{}
	err := json.NewDecoder(r.Body).Decode(&stanReq)
	if err != nil {
		errMsg := fmt.Sprintf("Could not decode request: %s", err.Error())
		s.Logger.Errorf("FilterStanJson :: %v", errMsg)
		apiError.ApiError(w, http.StatusBadRequest, errMsg)
		return
	}
	s.Logger.Infof("FilterStanJson :: Incoming stan filter request: %+v", stanReq)

	//Filter out
	filteredEps := FilterEpisodes(stanReq)
	s.Logger.Infof("FilterStanJson :: Filtered to: %+v", filteredEps)
	middleware.ModelResponse(w, 200, filteredEps)
}

//FilterEpisodes
//return the ones with DRM enabled and at least one episode
func FilterEpisodes(req *models.StanRequest) models.StanResponse {
	if req == nil {
		return models.StanResponse{}
	}

	respArray := []models.Response{}

	for _, payload := range req.Payload {
		if payload.Drm && payload.EpisodeCount > 0 {
			respArray = append(respArray, models.Response{
				Image: payload.Image.ShowImage,
				Slug:  payload.Slug,
				Title: payload.Title,
			})
		}
	}
	return models.StanResponse{
		Response: respArray,
	}
}
