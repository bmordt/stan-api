package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bmordt/stan-api/src/models"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	testLogger = newTestLogger()

	testRequestLocation  = "../testdata/example-request.json"
	testResponseLocation = "../testdata/example-response.json"

	testRequestBytes, testResponseBytes []byte

	testStanRequest = models.StanRequest{
		Payload: []models.StanPayload{
			models.StanPayload{
				Drm:          true,
				EpisodeCount: 0,
				Image: models.StanImage{
					ShowImage: "show-image-0",
				},
				Slug:  "some-slug-0",
				Title: "some-title-0",
			},
			models.StanPayload{
				Drm:          true,
				EpisodeCount: 1,
				Image: models.StanImage{
					ShowImage: "show-image-1",
				},
				Slug:  "some-slug-1",
				Title: "some-title-1",
			},
			models.StanPayload{
				Drm:          true,
				EpisodeCount: 2,
				Image: models.StanImage{
					ShowImage: "show-image-2",
				},
				Slug:  "some-slug-2",
				Title: "some-title-2",
			},
			models.StanPayload{
				Drm:          false,
				EpisodeCount: 3,
				Image: models.StanImage{
					ShowImage: "show-image-3",
				},
				Slug:  "some-slug-3",
				Title: "some-title-3",
			},
			models.StanPayload{
				Drm:          false,
				EpisodeCount: 4,
				Image: models.StanImage{
					ShowImage: "show-image-4",
				},
				Slug:  "some-slug-4",
				Title: "some-title-4",
			},
			models.StanPayload{
				Drm:          true,
				EpisodeCount: 5,
				Image: models.StanImage{
					ShowImage: "show-image-5",
				},
				Slug:  "some-slug-5",
				Title: "some-title-5",
			},
			models.StanPayload{
				Drm:          true,
				EpisodeCount: 6,
				Image: models.StanImage{
					ShowImage: "show-image-6",
				},
				Slug:  "some-slug-6",
				Title: "some-title-6",
			},
		},
	}

	testStanResponse = models.StanResponse{
		Response: []models.Response{
			models.Response{
				Image: "show-image-1",
				Slug:  "some-slug-1",
				Title: "some-title-1",
			},
			models.Response{
				Image: "show-image-2",
				Slug:  "some-slug-2",
				Title: "some-title-2",
			},
			models.Response{
				Image: "show-image-5",
				Slug:  "some-slug-5",
				Title: "some-title-5",
			},
			models.Response{
				Image: "show-image-6",
				Slug:  "some-slug-6",
				Title: "some-title-6",
			},
		},
	}
)

func init() {
	var err error
	testRequestBytes, err = os.ReadFile(testRequestLocation)
	if err != nil {
		testLogger.Fatalf("Error reading request test file")
	}
	testResponseBytes, err = os.ReadFile(testResponseLocation)
	if err != nil {
		testLogger.Fatalf("Error reading response test file")
	}
}

func TestFilterEpisodes(t *testing.T) {
	t.Run("Given a valid request, FilterEpisodes returns the expected response containing Slug Image and Title in an array", func(t *testing.T) {
		actualResp := FilterEpisodes(&testStanRequest)
		assert.Len(t, actualResp.Response, 4)

		t.Run("Each obj contains correct Slug Image and Title", func(t *testing.T) {
			for i, resp := range actualResp.Response {
				assert.Equal(t, testStanResponse.Response[i], resp)
			}
		})
	})
}

func TestFilterStanJson(t *testing.T) {
	t.Run("Given a valid request, the http response contains the expected response", func(t *testing.T) {
		expectedResponse := &models.StanResponse{}
		err := json.Unmarshal(testResponseBytes, expectedResponse)
		assert.NoError(t, err)

		s := NewStanService(testLogger)

		testIncomingReq := &http.Request{
			Body: getBody(testRequestBytes),
		}
		w := httptest.NewRecorder()

		s.FilterStanJson(w, testIncomingReq)
		resp := w.Result()

		t.Run("Response code is 200", func(t *testing.T) {
			assert.Equal(t, 200, resp.StatusCode)
		})

		t.Run("Response json should have response key", func(t *testing.T) {
			actualResp := &models.StanResponse{}
			err := json.Unmarshal(w.Body.Bytes(), actualResp)
			assert.NoError(t, err)

			assert.Equal(t, len(expectedResponse.Response), len(actualResp.Response))

			t.Run("Response contains the expected data", func(t *testing.T) {
				for i, resp := range actualResp.Response {
					assert.Equal(t, expectedResponse.Response[i], resp)
				}
			})
		})
	})
	t.Run("Given an invalid request, the http response contains 400 and string of \"Could not decode request\"", func(t *testing.T) {
		s := NewStanService(testLogger)

		testIncomingReq := &http.Request{
			Body: getBody([]byte("invalid req")),
		}
		w := httptest.NewRecorder()

		s.FilterStanJson(w, testIncomingReq)
		resp := w.Result()

		t.Run("Response code is 400", func(t *testing.T) {
			assert.Equal(t, 400, resp.StatusCode)
		})
		t.Run("Response contains a message", func(t *testing.T) {
			actualResp := make(map[string]string)
			err := json.Unmarshal(w.Body.Bytes(), &actualResp)
			assert.NoError(t, err)

			assert.Contains(t, actualResp["error"], "Could not decode request")
		})
	})
}

func newTestLogger() *logrus.Entry {
	testLogger := logrus.New()
	return testLogger.WithFields(logrus.Fields{})
}

func getBody(testJsonBytes []byte) io.ReadCloser {
	testBody := strings.NewReader(string(testJsonBytes))
	return io.NopCloser(testBody)
}
