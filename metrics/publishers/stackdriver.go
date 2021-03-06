package publishers

import (
	"fmt"

	"github.com/bazelbuild/continuous-integration/metrics/clients"
	"github.com/bazelbuild/continuous-integration/metrics/data"
	"github.com/bazelbuild/continuous-integration/metrics/metrics"
)

type Stackdriver struct {
	client    *clients.StackdriverClient
	projectID string
}

func (sd *Stackdriver) Name() string {
	return "Stackdriver"
}

func (sd *Stackdriver) RegisterMetric(metric metrics.Metric) error {
	// Nothing to do.
	return nil
}

func (sd *Stackdriver) Publish(metric metrics.Metric, newData data.DataSet) error {
	metricName := metric.Name()
	set, ok := newData.(data.StackDriverTimeSeriesDataSet)
	if !ok {
		return fmt.Errorf("Metric '%s' does not produce a valid StackDriverTimeSeriesDataSet instance", metricName)
	}

	err := sd.client.WriteTimeSeries(set.CreateTimeSeriesRequest(sd.projectID))
	if err != nil {
		return fmt.Errorf("Could not write time series for metric '%s' in project '%s': %v", metricName, sd.projectID, err)
	}
	return nil
}

func CreateStackdriverPublisher(client *clients.StackdriverClient, projectID string) *Stackdriver {
	return &Stackdriver{client: client, projectID: projectID}
}
