package pageHandler

import (
	"de/vorlesung/projekt/crew/sessionHandler"
	"net/http"
)

// Matrikelnummern:
//
// 3333958
// 3880065
// 8701350

// A-8.1:
// Die Bearbeitung der Tickets soll ausschließlich ¨uber eine WEB-Seite erfolgen.
//
// Loginseite
// Anmeldung mit Benutzer und Passwort
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	if sessionHandler.IsUserLoggedIn(request) {

		// Seite für den Angemeldeten User aufrufen
		http.Redirect(response, request, "/dashboard", 302)
	} else {

		// Loginseite falls kein User angemeldet ist
		http.ServeFile(response, request, sessionHandler.GetAssetsDir()+"html/loginView.html")
	}
}
