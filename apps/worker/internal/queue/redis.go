package queue

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type ReviewStartJob struct {
	DocumentID int64  `json:"document_id"`
	PlaybookID *int64 `json:"playbook_id,omitempty"`
}

type Client struct {
	rdb *redis.Client
	q   string
}

func NewClient() *Client {
	url := getenv("REDIS_URL", "redis://redis:6379")
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("invalid REDIS_URL: %v", err)
	}
	q := getenv("REVIEW_QUEUE_NAME", "review:start")
	return &Client{rdb: redis.NewClient(opt), q: q}
}

func (c *Client) Dequeue(ctx context.Context) (*ReviewStartJob, error) {
	// Blocking pop: waits up to 5s then loop again (to allow shutdowns)
	res, err := c.rdb.BLPop(ctx, 5*time.Second, c.q).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if len(res) != 2 {
		return nil, nil
	}
	var job ReviewStartJob
	if err := json.Unmarshal([]byte(res[1]), &job); err != nil {
		log.Printf("bad job payload: %v", err)
		return nil, nil
	}
	return &job, nil
}

func (c *Client) Enqueue(ctx context.Context, job ReviewStartJob) error {
	b, _ := json.Marshal(job)
	return c.rdb.RPush(ctx, c.q, b).Err()
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
