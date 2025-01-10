package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{
		name:          "go_page",
		renderer:      "go",
		template:      "home",
		errorExpected: false,
		errorMessage:  "error rendering go template:",
	}, {
		name:          "go_page_no_template",
		renderer:      "go",
		template:      "no-file",
		errorExpected: true,
		errorMessage:  "no error rendering non-existent go template when one is expected",
	}, {
		name:          "jet_page",
		renderer:      "jet",
		template:      "home",
		errorExpected: false,
		errorMessage:  "error rendering jet template:",
	}, {
		name:          "jet_page_no_template",
		renderer:      "jet",
		template:      "no-file",
		errorExpected: true,
		errorMessage:  "no error rendering non-existent jet template when one is expected",
	}, {
		name:          "invalid_renderer_engine",
		renderer:      "foo",
		template:      "home",
		errorExpected: true,
		errorMessage:  "no error rendering with non-existent template engine",
	},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s:", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}
