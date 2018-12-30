package ctrl

import (
	"html/template"
	"net/http"

	"github.com/gy-kim/building-better-go-web-apps-with-the-gorilla-toolkit/bw/vm"
)

type loginController struct {
	loginTemplate *template.Template
}

func (lc *loginController) GetLogin(w http.ResponseWriter, r *http.Request) {
	vmodel := vm.Base{}
	lc.loginTemplate.Execute(w, vmodel)
}
