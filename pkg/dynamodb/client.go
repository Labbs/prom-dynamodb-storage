package dynamodb

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/prometheus/common/model"
)

type Client struct {
	session *dynamo.DB
	table   dynamo.Table
}

type metric struct {
	NameID string
	Time   time.Time

	Tags  map[string]string `dynamo:"Tags"`
	Value string            `dynamo:"Value`
}

func NewClient(table string, region string) *Client {
	s := dynamo.New(session.New(), &aws.Config{Region: aws.String(region)})
	return &Client{
		session: s,
		table:   s.Table(table),
	}
}

func tagsFromMetric(m model.Metric) map[string]string {
	tags := make(map[string]string, len(m)-1)
	for l, v := range m {
		if l != model.MetricNameLabel {
			tags[string(l)] = string(v)
		}
	}
	return tags
}

func (c *Client) Write(samples model.Samples) error {
	for _, s := range samples {
		v := float64(s.Value)
		w := metric{NameID: string(s.Metric[model.MetricNameLabel]), Tags: tagsFromMetric(s.Metric), Value: strconv.FormatFloat(v, 'f', 6, 64), Time: s.Timestamp.Time()}
		err := c.table.Put(w).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
