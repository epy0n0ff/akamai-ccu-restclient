package restclient

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDecodePurgeResponse(t *testing.T) {
	var resp string = `
{
   "detail" : "Request accepted.",
   "estimatedSeconds" : 420,
   "httpStatus" : 201,
   "pingAfterSeconds" : 420,
   "progressUri" : "/ccu/v2/purges/95b5a092-043f-4af0-843f-aaf0043faaf0",
   "purgeId" : "95b5a092-043f-4af0-843f-aaf0043faaf0",
   "supportId" : "17PY1321286429616716-211907680"
}`

	var r PurgeResponse
	err := json.NewDecoder(bytes.NewBufferString(resp)).Decode(&r)
	if err != nil {
		t.Fatalf(`unexpected error: %v`, err)
		return
	}

	if !reflect.DeepEqual(`Request accepted.`, r.Detail) {
		t.Fatalf(`unexpected error: Detail "Request accepted." != %s`, r.Detail)
	}
	if !reflect.DeepEqual(420, r.EstimatedSeconds) {
		t.Fatalf(`unexpected error: EstimatedSeconds 420 != %d`, r.EstimatedSeconds)
	}
	if !reflect.DeepEqual(201, r.HttpStatus) {
		t.Fatalf(`unexpected error: HttpStatus 201 != %d`, r.HttpStatus)
	}
	if !reflect.DeepEqual(`95b5a092-043f-4af0-843f-aaf0043faaf0`, r.PurgeId) {
		t.Fatalf(`unexpected error: PurgeId "95b5a092-043f-4af0-843f-aaf0043faaf0" != %s`, r.PurgeId)
	}
	if !reflect.DeepEqual(`/ccu/v2/purges/95b5a092-043f-4af0-843f-aaf0043faaf0`, r.ProgressUri) {
		t.Fatalf(`unexpected error: progressUri "/ccu/v2/purges/95b5a092-043f-4af0-843f-aaf0043faaf0" != %s`, r.ProgressUri)
	}
	if !reflect.DeepEqual(420, r.PingAfterSeconds) {
		t.Fatalf(`unexpected error: pingAfterSeconds 420 != %d`, r.PingAfterSeconds)
	}
	if !reflect.DeepEqual(`17PY1321286429616716-211907680`, r.SupportId) {
		t.Fatalf(`unexpected error: supportId "17PY1321286429616716-211907680" != %s`, r.SupportId)
	}
}

func TestDecodeQueueResponse(t *testing.T) {
	var resp string = `
{
   "detail" : "The queue may take a minute to reflect new or removed requests.",
   "httpStatus" : 200,
   "queueLength" : 17,
   "supportId" : "17QY1321286863376510-220300384"
}`

	var r QueueResponse
	err := json.NewDecoder(bytes.NewBufferString(resp)).Decode(&r)
	if err != nil {
		t.Fatalf(`unexpected error: %v`, err)
		return
	}

	if !reflect.DeepEqual(`The queue may take a minute to reflect new or removed requests.`, r.Detail) {
		t.Fatalf(`unexpected error: Detail "The queue may take a minute to reflect new or removed requests." != %s`, r.Detail)
	}
	if !reflect.DeepEqual(200, r.HttpStatus) {
		t.Fatalf(`unexpected error: HttpStatus 200 != %d`, r.HttpStatus)
	}
	if !reflect.DeepEqual(17, r.QueueLength) {
		t.Fatalf(`unexpected error: QueueLength 17 != %d`, r.QueueLength)
	}
	if !reflect.DeepEqual(`17QY1321286863376510-220300384`, r.SupportId) {
		t.Fatalf(`unexpected error: SupportId "17QY1321286863376510 != %s`, r.SupportId)
	}
}

func TestDecodePurgeStatusResponse(t *testing.T) {
	var resp string = `
{
    "completionTime": null,
    "httpStatus": 200,
    "originalEstimatedSeconds": 480,
    "originalQueueLength": 6,
    "pingAfterSeconds": 60,
    "progressUri": "/ccu/v2/purges/142eac1d-99ab-11e3-945a-7784545a7784",
    "purgeId": "142eac1d-99ab-11e3-945a-7784545a7784",
    "purgeStatus": "In-Progress",
    "submissionTime": "2014-02-19T21:16:20Z",
    "submittedBy": "test1",
    "supportId": "17SY1392844709041263-238396512"
}`

	var r PurgeStatusResponse
	err := json.NewDecoder(bytes.NewBufferString(resp)).Decode(&r)
	if err != nil {
		t.Fatalf(`unexpected error: %v`, err)
		return
	}

	if r.CompletionTime != nil {
		t.Fatalf(`unexpected error: CompletionTime != nil`)
	}
	if !reflect.DeepEqual(200, r.HttpStatus) {
		t.Fatalf(`unexpected error: HttpStatus 200 != %d`, r.HttpStatus)
	}
	if !reflect.DeepEqual(480, r.OriginalEstimatedSeconds) {
		t.Fatalf(`unexpected error: OriginalEstimatedSeconds 480 != %d`, r.OriginalEstimatedSeconds)
	}
	if !reflect.DeepEqual(6, r.OriginalQueueLength) {
		t.Fatalf(`unexpected error: OriginalQueueLength 6 != %d`, r.OriginalQueueLength)
	}
	if !reflect.DeepEqual(60, r.PingAfterSeconds) {
		t.Fatalf(`unexpected error: PingAfterSeconds 60 != %d`, r.PingAfterSeconds)
	}
	if !reflect.DeepEqual(`/ccu/v2/purges/142eac1d-99ab-11e3-945a-7784545a7784`, r.ProgressUri) {
		t.Fatalf(`unexpected error: ProgressUri "/ccu/v2/purges/142eac1d-99ab-11e3-945a-7784545a7784" != %s`, r.ProgressUri)
	}
	if !reflect.DeepEqual(`In-Progress`, r.PurgeStatus) {
		t.Fatalf(`unexpected error: PurgeStatus "In-Progress" != %s`, r.PurgeStatus)
	}
	if !reflect.DeepEqual(`2014-02-19T21:16:20Z`, r.SubmissionTime.Format(time.RFC3339)) {
		t.Fatalf(`unexpected error: SubmissionTime "2014-02-19T21:16:20Z" != %s`, r.SubmissionTime.Format(time.RFC3339))
	}
	if !reflect.DeepEqual(`test1`, r.SubmittedBy) {
		t.Fatalf(`unexpected error: SubmittedBy "test1" != %s`, r.SubmittedBy)
	}
	if !reflect.DeepEqual(`17SY1392844709041263-238396512`, r.SupportId) {
		t.Fatalf(`unexpected error: SupportId "17SY1392844709041263-238396512" != %s`, r.SupportId)
	}
}
