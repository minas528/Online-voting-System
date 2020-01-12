package handler

import (
	"context"
	"github.com/betsegawlemma/restaurant/entity"
	"github.com/minas528/Online-voting-System/form"
	"github.com/minas528/Online-voting-System/permission"
	"github.com/minas528/Online-voting-System/rtoken"
	"github.com/minas528/Online-voting-System/session"
	"github.com/minas528/Online-voting-System/voters"
	"github.com/minas528/Online-voting-System/entities"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type VoterHandler struct {
	tmpl     *template.Template
	voterserv voters.VotersService
	sessionService voters.SessionService
	voterSess *entities.Session
	loggedInVoter *entities.Voters
	voterRole voters.RoleService
	csrfSignKey []byte
}

type contextKey string

var ctxVoterSessionKey = contextKey("signed_in_user_session")

func NewVoterHandler(T *template.Template, AS voters.VotersService,
	sessServ voters.SessionService,
	role voters.RoleService,vtrsess *entities.Session,csKye []byte) *VoterHandler {
	return &VoterHandler{tmpl:T,voterserv:AS,sessionService:sessServ,
		voterRole:role,voterSess:vtrsess,csrfSignKey:csKye}
}

func (uh *VoterHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxVoterSessionKey, uh.voterSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}


func (uh *VoterHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInVoter == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles, errs := uh.voterserv.VoterRoles(uh.loggedInVoter)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (uh *VoterHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		usr, errs := uh.voterserv.VoterByGID(r.FormValue("g_id"))
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Your GID address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your GID address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}

		uh.loggedInVoter = usr
		claims := rtoken.Claims(usr.GID, uh.voterSess.Expires)
		session.Create(claims, uh.voterSess.UUID, uh.voterSess.SigningKey, w)
		newSess, errs := uh.sessionService.StoreSession(uh.voterSess)
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "login.layout", loginForm)
			return
		}
		uh.voterSess = newSess
		roles, _ := uh.voterserv.VoterRoles(usr)
		if uh.checkAdmin(roles) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}


func (uh *VoterHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxVoterSessionKey).(*entity.Session)
	session.Remove(userSess.UUID, w)
	uh.sessionService.DeleteSession(userSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}


func (uh *VoterHandler) Signup(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	log.Println("here")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("fullname", "g_id", "password", "confirmpassword")
		singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		singnUpForm.PasswordMatches("password", "confirmpassword")
		singnUpForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		pExists := uh.voterserv.PhoneExists(r.FormValue("phone"))
		if pExists {
			singnUpForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}
		eExists := uh.voterserv.GIDExists(r.FormValue("g_id"))
		if eExists {
			singnUpForm.VErrors.Add("g_id", "GID Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		role, errs := uh.voterRole.RoleByName("USER")

		if len(errs) > 0 {
			singnUpForm.VErrors.Add("role", "could not assign role to the user")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		user := &entities.Voters{
			FullName: r.FormValue("fullname"),
			GID:    r.FormValue("g_id"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   role.ID,
		}
		_, errs = uh.voterserv.StoreVoter(user)
		log.Println("again")
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (uh *VoterHandler) loggedIn(r *http.Request) bool {
	if uh.voterSess == nil {
		return false
	}
	userSess := uh.voterSess
	c, err := r.Cookie(userSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}

func (uh *VoterHandler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	voters, errs := uh.voterserv.Voters()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Users   []entities.Voters
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Users:   voters,
		CSRF:    token,
	}
	uh.tmpl.ExecuteTemplate(w, "admin.users.layout", tmplData)
}

func (uh *VoterHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		roles, errs := uh.voterRole.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		accountForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entities.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Roles:   roles,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate the form contents
		accountForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		accountForm.Required("fullname", "g_id", "password", "confirmpassword")
		accountForm.MatchesPattern("phone", form.PhoneRX)
		accountForm.MinLength("password", 8)
		accountForm.PasswordMatches("password", "confirmpassword")
		accountForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !accountForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		pExists := uh.voterserv.PhoneExists(r.FormValue("phone"))
		if pExists {
			accountForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		eExists := uh.voterserv.GIDExists(r.FormValue("g_id"))
		if eExists {
			accountForm.VErrors.Add("g_id", "GID Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			accountForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			accountForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		user := &entities.Voters{
			FullName: r.FormValue("fullname"),
			GID:    r.FormValue("g_id"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   uint(roleID),
		}
		_, errs := uh.voterserv.StoreVoter(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/voters", http.StatusSeeOther)
	}
}

func (uh *VoterHandler) checkAdmin(rs []entities.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}

func (uh *VoterHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		user, errs := uh.voterserv.Voter(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		roles, errs := uh.voterRole.Roles()
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		role, errs := uh.voterRole.Role(user.RoleID)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("userid", idRaw)
		values.Add("fullname", user.FullName)
		values.Add("g_id", user.GID)
		values.Add("role", string(user.RoleID))
		values.Add("phone", user.Phone)
		values.Add("rolename", role.Name)

		upAccForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   []entities.Role
			User    *entities.Voters
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Roles:   roles,
			User:    user,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		upAccForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		upAccForm.Required("fullname", "g_id", "phone")
		upAccForm.MatchesPattern("phone", form.PhoneRX)
		upAccForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !upAccForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		userID := r.FormValue("userid")
		uid, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		user, errs := uh.voterserv.Voter(uint(uid))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		eExists := uh.voterserv.GIDExists(r.FormValue("g_id"))
		if (user.GID != r.FormValue("g_id")) && eExists {
			upAccForm.VErrors.Add("g_id", "GID Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		pExists := uh.voterserv.PhoneExists(r.FormValue("phone"))

		if (user.Phone != r.FormValue("phone")) && pExists {
			upAccForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		usr := &entities.Voters{
			ID:       user.ID,
			FullName: r.FormValue("fullname"),
			GID:    r.FormValue("g_id"),
			Phone:    r.FormValue("phone"),
			Password: user.Password,
			RoleID:   uint(roleID),
		}
		_, errs = uh.voterserv.UpdateVoter(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

func (uh *VoterHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := uh.voterserv.Deletevoter(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}