package http

import (
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/post/repository"
	"github.com/minas528/Online-voting-System/post/service"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPost(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo :=repository.NewMockPostGormRepo(nil)
	pstserv:= service.NewPostService(pstRepo)

	adminPsthandler := handler.NewAdminPostHandler(templ,pstserv,nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/posts", adminPsthandler.Posts)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/posts")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestPostNew(t *testing.T)  {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo :=repository.NewMockPostGormRepo(nil)
	pstserv:= service.NewPostService(pstRepo)

	adminPsthandler := handler.NewAdminPostHandler(templ,pstserv,nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/posts/new", adminPsthandler.PostNew)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/posts/new")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestAdminPostsUpdate(t *testing.T) {

	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo :=repository.NewMockPostGormRepo(nil)
	pstserv:= service.NewPostService(pstRepo)

	adminPsthandler := handler.NewAdminPostHandler(templ,pstserv,nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/posts/update", adminPsthandler.AdminPostsUpdate)
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

	resp, err := tc.PostForm(sUrl + "/admin/posts/update?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminPostDelete(t *testing.T) {
	templ :=template.Must(template.ParseGlob("../../ui/templates/*"))
	pstRepo :=repository.NewMockPostGormRepo(nil)
	pstserv:= service.NewPostService(pstRepo)

	adminPsthandler := handler.NewAdminPostHandler(templ,pstserv,nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/posts/delete", adminPsthandler.AdminPostDelete)
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

	resp, err := tc.PostForm(sUrl + "/admin/posts/delete?id=1",form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
