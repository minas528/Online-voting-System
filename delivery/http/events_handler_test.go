package http


import (
	"github.com/minas528/Online-voting-System/Event/repository"
	"github.com/minas528/Online-voting-System/Event/service"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
"github.com/minas528/Online-voting-System/entities"
"html/template"
"net/http"
"net/http/httptest"
"net/url"
"testing"
)

func TestEvent(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockEventRepository(nil)
	pstserv:= service.NewEventService(pstRepo)

	adminPsthandler := handler.NewEventHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events", adminPsthandler.Events)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/events")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestEventNew(t *testing.T)  {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockEventRepository(nil)
	pstserv:= service.NewEventService(pstRepo)

	adminPsthandler := handler.NewEventHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/new", adminPsthandler.EventNew)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/events/new")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestAdminEventUpdate(t *testing.T) {

	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockEventRepository(nil)
	pstserv:= service.NewEventService(pstRepo)

	adminPsthandler := handler.NewEventHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/update", adminPsthandler.UpdateEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sUrl := ts.URL

	form := url.Values{}
	form.Add("ID",string(entities.PostMock.ID))
	form.Add("Vid",entities.PostMock.Vid)
	form.Add("Disc",entities.PostMock.Disc)
	form.Add("Name",entities.PostMock.Name)
	form.Add("Writer",entities.PostMock.Writer)

	resp, err := tc.PostForm(sUrl + "/admin/events/update?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminEventDelete(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockEventRepository(nil)
	pstserv:= service.NewEventService(pstRepo)

	adminPsthandler := handler.NewEventHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/delete", adminPsthandler.DeleteEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sUrl := ts.URL

	form := url.Values{}
	form.Add("ID",string(entities.PostMock.ID))
	form.Add("Vid",entities.PostMock.Vid)
	form.Add("Disc",entities.PostMock.Disc)
	form.Add("Name",entities.PostMock.Name)
	form.Add("Writer",entities.PostMock.Writer)

	resp, err := tc.PostForm(sUrl + "/admin/events/delete?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

