package mltq

import "testing"

func TestLogToQuery(t *testing.T) {
	tests := []struct {
		input         string
		expectedQuery string
	}{
		{
			`{"t":{"$date":"2020-11-20T06:33:15.797+00:00"},"s":"I","c":"COMMAND","id":51803,"ctx":"conn63203","msg":"Slow query","attr":{"type":"command","ns":"default_chat.users","command":{"find":"users","filter":{"_updatedAt":{"$gte":{"$date":"2020-11-20T06:33:08.015Z"},"$lt":{"$date":"2020-11-20T06:33:12.622Z"}},"visitorResponded":true,"test":{"$ne":true},"dummy":{"$ne":true},"isDeleted":{"$ne":true}},"sort":{"_updatedAt":1},"limit":500,"maxTimeMS":2000,"returnKey":false,"showRecordId":false,"$clusterTime":{"clusterTime":{"$timestamp":{"t":1605853994,"i":11}},"signature":{"hash":{"$binary":{"base64":"reHNIHbNvR7Hi6LUIj4mVDwAX8k=","subType":"0"}},"keyId":6881916503147413508}},"lsid":{"id":{"$uuid":"b4b2219f-f0f7-4102-b695-b864274d7fe0"}},"$db":"default_users"},"planSummary":"COLLSCAN","keysExamined":0,"docsExamined":27431,"hasSortStage":true,"cursorExhausted":true,"numYields":30,"nreturned":1,"queryHash":"DA4A9692","planCacheKey":"EDBD4246","reslen":1489,"locks":{"ReplicationStateTransition":{"acquireCount":{"w":31}},"Global":{"acquireCount":{"r":31}},"Database":{"acquireCount":{"r":31}},"Collection":{"acquireCount":{"r":31}},"Mutex":{"acquireCount":{"r":1}}},"storage":{},"protocol":"op_query","durationMillis":175}}`,
			"x",
		},
	}

	for _, tt := range tests {
		if q, _ := LogToQuery(tt.input); q != tt.expectedQuery {
			t.Fatalf("query did not match\nexpected: %s\ngot:%s", tt.expectedQuery, q)
			return
		}
	}
}
