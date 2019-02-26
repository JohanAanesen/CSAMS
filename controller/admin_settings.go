package controller

import (
	"bytes"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/view"
	"io"
	"log"
	"net/http"
)

// AdminSettingsGET handles GET-request at admin/setting
func AdminSettingsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/settings/index"

	v.Render(w)
}

// AdminSettingsPOST handles POST-request at admin/setting
func AdminSettingsPOST(w http.ResponseWriter, r *http.Request) {
	// TODO (Svein): Handle incoming data
	http.Redirect(w, r, "/admin/settings", http.StatusOK)
}

// AdminSettingsImportGET handles GET-request at admin/setting/import
func AdminSettingsImportGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/settings/import"

	v.Render(w)
}

// AdminSettingsImportPOST handles POST-request at admin/setting/import
func AdminSettingsImportPOST(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("db_import")
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()
	defer buffer.Reset()

	_, err = io.Copy(&buffer, file)
	if err != nil {
		log.Println(err)
		return
	}

	content := buffer.String()
	fmt.Println(content)

	// TODO (Svein): Backup of current DB
	// TODO (Svein): Query this file
	// TODO (Svein): Save the world

	http.Redirect(w, r, "/admin/settings", http.StatusOK)
}
