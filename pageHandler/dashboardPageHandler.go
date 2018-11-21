package pageHandler

import (
	"fmt"
	"html/template"
	"net/http"
	"ticketBackend/sessionHandler"
	"ticketBackend/ticket"
)

//anzeigen der Dashboard Seite
func DashboardViewPageHandler(response http.ResponseWriter, request *http.Request) {
	if sessionHandler.IsUserLoggedIn(request) {

		var templateFiles []string
		templateFiles = append(templateFiles, "./assets/html/dashboardViewTemplate.html")
		templateFiles = append(templateFiles, "./assets/html/ticketListTemplate.html")
		templateFiles = append(templateFiles, "./assets/html/dashboardViewFooterTemplate.html")

		templates, err := template.ParseFiles(templateFiles...)
		if err != nil {
			fmt.Println(err)
		}

		templates.ExecuteTemplate(response, "outer", sessionHandler.GetSessionUser().Username)

		pTickets := *ticket.GetTickets(ticket.Open)

		for i := 0; i < len(pTickets); i++ {
			templates.ExecuteTemplate(response, "inner", pTickets[i])
		}

		templates.ExecuteTemplate(response, "footer", nil)

		templates.Execute(response, nil)

	} else {
		http.Redirect(response, request, "/", 302)
	}
}
