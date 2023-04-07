package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewServerMain(t *testing.T) {
	s := NewServer()

	ts := httptest.NewServer(http.HandlerFunc(s.HandleRequest))

	reqCount := 1000
	for i := 0; i < reqCount; i++ {
		id := i%100 + 1
		res, err := http.Get(fmt.Sprintf("%s/?id=%d", ts.URL, id))
		if err != nil {
			t.Errorf("failed to get response %v", err)
		}
		var user *User
		if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
			t.Errorf("failed to decode response %v", err)
		}
		//fmt.Printf("id:%d user:%+v", i, user)
	}
	fmt.Printf("\ndb hit %d\n", s.db.dbhit)

}
