package main

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"web/art"
)

func ReadTestString() []string {
	data, err := ioutil.ReadFile("test-output.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)
	contentSplit := strings.Split(contentString, "#")

	return contentSplit
}

func TestAscii_Art_End(t *testing.T) {
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=xyz{|}~", nil)
	req1.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	// titles := r.FindAllString(actual, -1)
	titlesMaybe := r.FindStringSubmatch(actual)

	// fmt.Println(titlesMaybe[1])
	// fmt.Println(ReadTestString()[1])

	if titlesMaybe[1] != ReadTestString()[2] {
		t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s",
			"xyz{|}~", ReadTestString()[2], titlesMaybe[1])
	}
}

func TestAscii_Art_Middle(t *testing.T) {
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=OPEN", nil)
	req1.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	titlesMaybe := r.FindStringSubmatch(actual)

	// fmt.Println(titlesMaybe[1])

	if titlesMaybe[1] != ReadTestString()[1] {
		t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s",
			"OPEN", ReadTestString()[1], titlesMaybe[1])
	}
}

func TestAscii_Art_Beginning(t *testing.T) {
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=+%21%22%23", nil)
	req1.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	titlesMaybe := r.FindStringSubmatch(actual)

	// fmt.Println(titlesMaybe[1])

	if titlesMaybe[1] != ReadTestString()[0] {
		t.Errorf("\nTest fails when given case:\n\t\"%s\","+"\nThe test should show:\n%s\nInstead it shows:\n%s",
			"%21%22%23", ReadTestString()[0], titlesMaybe[1])
	}
}
