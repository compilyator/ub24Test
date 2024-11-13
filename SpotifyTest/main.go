package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const (
	tokenURL           = "https://accounts.spotify.com/api/token"
	recommendationsURL = "https://api.spotify.com/v1/recommendations"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SpotifyRecommendationsResponse struct {
	Tracks []struct {
		Name    string `json:"name"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Album struct {
			ReleaseDate string `json:"release_date"`
		} `json:"album"`
	} `json:"tracks"`
}

func getSpotifyAccessToken(clientID, clientSecret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s\n", string(reqDump))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RESPONSE:\n%s\n", string(respDump))

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token: %s", resp.Status)
	}

	var tokenResponse SpotifyTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func getSpotifyRecommendations(accessToken, seedGenres string, wg *sync.WaitGroup, ch chan<- []byte, errCh chan<- error, i int) {
	defer wg.Done()

	time.Sleep(500)

	params := url.Values{}
	params.Add("seed_genres", seedGenres)
	params.Add("limit", "10")

	req, err := http.NewRequest("GET", recommendationsURL+"?"+params.Encode(), nil)
	if err != nil {
		errCh <- err
		return
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	fmt.Printf("Making request #%d\n", i)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errCh <- err

		fmt.Printf("Request failed to be done #%d\n", i)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errCh <- fmt.Errorf("failed to get recommendations: %s", resp.Status)

		fmt.Printf("Request done #%d with error code%d\n", i, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errCh <- err
		return
	}

	ch <- body

	fmt.Printf("Request done #%d\n", i)

}

func recommendHandler(w http.ResponseWriter, r *http.Request) {
	seedGenres := r.URL.Query().Get("seed_genres")
	if seedGenres == "" {
		http.Error(w, "Missing seed_genres parameter", http.StatusBadRequest)
		return
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	token, err := getSpotifyAccessToken(clientID, clientSecret)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting Spotify access token: %v", err), http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	ch := make(chan []byte, 1)
	errCh := make(chan error, 1)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		fmt.Printf("Starting thread #%d\n", i)
		go getSpotifyRecommendations(token, seedGenres, &wg, ch, errCh, i)
	}

	wg.Wait()
	close(ch)
	close(errCh)

	select {
	case recommendations := <-ch:
		var recResponse SpotifyRecommendationsResponse
		err := json.Unmarshal(recommendations, &recResponse)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing Spotify recommendations: %v", err), http.StatusInternalServerError)
			return
		}

		recommendationDetails := []string{}
		for _, track := range recResponse.Tracks {
			artistNames := []string{}
			for _, artist := range track.Artists {
				artistNames = append(artistNames, artist.Name)
			}
			releaseYear := "Unknown"
			if len(track.Album.ReleaseDate) >= 4 {
				releaseYear = track.Album.ReleaseDate[:4]
			}
			recommendationDetails = append(recommendationDetails, fmt.Sprintf("%s - %s (%s)", strings.Join(artistNames, ", "), track.Name, releaseYear))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recommendationDetails)
	case err := <-errCh:
		http.Error(w, fmt.Sprintf("Error getting Spotify recommendations: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/recommend", recommendHandler)

	port := ":8080"
	log.Printf("Server started at http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
