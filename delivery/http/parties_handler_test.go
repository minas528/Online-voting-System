package http

import (
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	"github.com/minas528/Online-voting-System/entities"
	"net/url"

	//"github.com/minas528/Online-voting-System/delivery/http/handler"
	//"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/parties/repository"
	"github.com/minas528/Online-voting-System/parties/service"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParties(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockPartiesGormRepo(nil)
	pstserv:= service.NewPartiesService(pstRepo)

	adminPsthandler := handler.NewAdminPartiesHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/parties", adminPsthandler.Parties)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/parties")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestPartiesNew(t *testing.T)  {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockPartiesGormRepo(nil)
	pstserv:= service.NewPartiesService(pstRepo)

	adminPsthandler := handler.NewAdminPartiesHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/parties/new", adminPsthandler.PartiesNew)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/parties/new")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestAdminPartiesUpdate(t *testing.T) {

	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockPartiesGormRepo(nil)
	pstserv:= service.NewPartiesService(pstRepo)

	adminPsthandler := handler.NewAdminPartiesHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/parties/update", adminPsthandler.AdminPartiesUpdate)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sUrl := ts.URL

	form := url.Values{}
	form.Add("ID",string(entities.PartiesMock.ID))
	form.Add("Vid",entities.PartiesMock.Logo)
	form.Add("Slogan",entities.PartiesMock.Slogan)
	form.Add("Name",entities.PartiesMock.Name)

	resp, err := tc.PostForm(sUrl + "/admin/parties/update?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminPartiesDelete(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo := repository.NewMockPartiesGormRepo(nil)
	pstserv:= service.NewPartiesService(pstRepo)

	adminPsthandler := handler.NewAdminPartiesHandler(templ,pstserv)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/parties/delete", adminPsthandler.AdminPartiesDelete)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sUrl := ts.URL

	form := url.Values{}
	form.Add("ID",string(entities.PartiesMock.ID))
	form.Add("Logo",entities.PartiesMock.Logo)
	form.Add("Slogan",entities.PartiesMock.Slogan)
	form.Add("Name",entities.PostMock.Name)

	resp, err := tc.PostForm(sUrl + "/admin/parties/delete?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}


