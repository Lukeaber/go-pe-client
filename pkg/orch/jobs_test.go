package orch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobs(t *testing.T) {

	// Test success
	setupGetResponder(t, jobs, "", "jobs-response.json")
	actual, err := orchClient.Jobs()
	require.Nil(t, err)
	require.Equal(t, expectedJobs, actual)

	// Test error
	setupErrorResponder(t, jobs)
	actual, err = orchClient.Jobs()
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

func TestJob(t *testing.T) {

	testURL := strings.ReplaceAll(job, "{job-id}", "123")

	// Test success
	setupGetResponder(t, testURL, "", "job-response.json")
	actual, err := orchClient.Job("123")
	require.Nil(t, err)
	require.Equal(t, expectedJob, actual)

	// Test error
	setupErrorResponder(t, testURL)
	actual, err = orchClient.Job("123")
	require.Nil(t, actual)
	require.Equal(t, expectedError, err)

}

var expectedJobs = &Jobs{Items: []Job{Job{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234", Name: "1234", Command: "deploy", Options: Options{Concurrency: interface{}(nil), Noop: false, Trace: false, Debug: false, Scope: Scope{Application: "", Nodes: []string(nil), Query: []interface{}(nil), NodeGroup: ""}, EnforceEnvironment: true, Environment: "production", Evaltrace: false, Target: interface{}(nil), Description: "deploy the web app"}, NodeCount: 5, Owner: Owner{ID: "751a8f7e-b53a-4ccd-9f4f-e93db6aa38ec", Login: "brian"}, Description: "deploy the web app", Timestamp: "2016-05-20T16:45:31Z", Environment: Environment{Name: "production"}, Status: []Status(nil), Nodes: Nodes{ID: "https://localhost:8143/orchestrator/v1/jobs/375/nodes"}, Report: Report{ID: "https://localhost:8143/orchestrator/v1/jobs/375/report"}}}, Pagination: Pagination{Limit: 20, Offset: 0, Total: 42}}

var expectedJob = &Job{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234", Name: "1234", Command: "deploy", Options: Options{Concurrency: interface{}(nil), Noop: false, Trace: false, Debug: false, Scope: Scope{Application: "Wordpress_app", Nodes: []string(nil), Query: []interface{}(nil), NodeGroup: ""}, EnforceEnvironment: true, Environment: "production", Evaltrace: false, Target: interface{}(nil), Description: ""}, NodeCount: 5, Owner: Owner{ID: "751a8f7e-b53a-4ccd-9f4f-e93db6aa38ec", Login: "admin"}, Description: "deploy the web app", Timestamp: "2016-05-20T16:45:31Z", Environment: Environment{Name: "production"}, Status: []Status{Status{State: "new", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:44:31Z"}, Status{State: "ready", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:44:31Z"}, Status{State: "running", EnterTime: "2016-04-11T18:44:31Z", ExitTime: "2016-04-11T18:45:31Z"}, Status{State: "finished", EnterTime: "2016-04-11T18:45:31Z", ExitTime: ""}}, Nodes: Nodes{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234/nodes"}, Report: Report{ID: "https://orchestrator.example.com:8143/orchestrator/v1/jobs/1234/report"}}
