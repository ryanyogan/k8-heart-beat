package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	buildTime := time.Now().Format("20060102_03:04:05")
	commit := "232jdk3"
	release := "0.0.0"
	p := ping(buildTime, commit, release)
	p(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong.  Have: %d, want: %d", have, want)
	}

	pong, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}

	err = json.Unmarshal(pong, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.Release != release {
		t.Errorf("Release version is wrong. Got: %s, expected: %s", info.Release, release)
	}
	if info.BuildTime != buildTime {
		t.Errorf("Build time is wrong. Got: %s, expected: %s", info.BuildTime, buildTime)
	}
	if info.Commit != commit {
		t.Errorf("Commit sha is wrong. Got: %s, expected: %s", info.Commit, commit)
	}
}
