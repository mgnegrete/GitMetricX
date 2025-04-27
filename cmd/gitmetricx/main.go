package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	commitCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gitmetricx_commits_total",
			Help: "Total number of commits in the repository",
		},
		[]string{"repo"},
	)
)

func init() {
	prometheus.MustRegister(commitCount)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	repo := os.Getenv("GITHUB_REPO")
	owner := os.Getenv("GITHUB_OWNER")
	token := os.Getenv("GITHUB_TOKEN")

	fmt.Printf("GitHub Owner: %s\n", owner)
	fmt.Printf("GitHub Repo: %s\n", repo)
	fmt.Printf("GitHub Token: %s\n", token)

	client := resty.New()
	var contributors []struct {
		Author struct {
			Login string `json:"login"`
		} `json:"author"`
		Total int `json:"total"`
	}

	resp, err := client.R().
		SetHeader("Authorization", "token "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(&contributors).
		Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/stats/contributors", owner, repo))

	if err != nil {
		log.Printf("Error fetching commit data: %v", err)
		http.Error(w, "Failed to fetch commit data", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Printf("GitHub API returned non-200 status: %d", resp.StatusCode())
		http.Error(w, "GitHub API error", http.StatusInternalServerError)
		return
	}

	for _, contributor := range contributors {
		commitCount.WithLabelValues(contributor.Author.Login).Set(float64(contributor.Total))
	}

	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/metrics", http.HandlerFunc(metricsHandler))
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
