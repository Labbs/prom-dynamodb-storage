package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"prom-dynamodb-storage/pkg/dynamodb"
	"prom-dynamodb-storage/pkg/settings"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
	"github.com/urfave/cli"
)

var version = "0.2"

func main() {
	app := cli.NewApp()
	app.Name = "prom-dynamodb-storage"
	app.Flags = settings.NewContext()
	app.Action = runWeb
	app.Version = version

	app.Run(os.Args)
}

func protoToSamples(req *prompb.WriteRequest) model.Samples {
	var samples model.Samples
	for _, ts := range req.Timeseries {
		metric := make(model.Metric, len(ts.Labels))
		for _, l := range ts.Labels {
			metric[model.LabelName(l.Name)] = model.LabelValue(l.Value)
		}

		for _, s := range ts.Samples {
			samples = append(samples, &model.Sample{
				Metric:    metric,
				Value:     model.SampleValue(s.Value),
				Timestamp: model.Time(s.Timestamp),
			})
		}
	}
	return samples
}

func runWeb(ctx *cli.Context) {
	// init logging
	settings.LoggerContext(ctx)

	// init dynanodb session
	client := dynamodb.NewClient(settings.DynamoTable, settings.DynamoRegion)

	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		compressed, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		reqBuf, err := snappy.Decode(nil, compressed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req prompb.WriteRequest
		if err := proto.Unmarshal(reqBuf, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		samples := protoToSamples(&req)

		err = client.Write(samples)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	})

	settings.Logger.Info().Str("event", "boot").Msg("Starting server on " + strconv.Itoa(settings.ListenPort))
	http.ListenAndServe(":"+strconv.Itoa(settings.ListenPort), nil)

}
