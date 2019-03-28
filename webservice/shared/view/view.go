package view

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// View struct
type View struct {
	Extension string   `json:"extension"`
	Folder    string   `json:"folder"`
	Name      string   `json:"name"`
	Caching   bool     `json:"caching"`
	Template  Template `json:"template"`

	Vars map[string]interface{}

	request *http.Request
}

// Template struct
type Template struct {
	Root     string   `json:"root"`
	Children []string `json:"children"`
}

var (
	rootTemplate   string
	childTemplates []string

	adminRootTemplate   string
	adminChildTemplates []string

	templateCollection = make(map[string]*template.Template)
	pluginCollection   = make(template.FuncMap)

	cfgView *View

	mutex       sync.RWMutex
	pluginMutex sync.RWMutex
)

// Configure the view
func Configure(v *View) {
	cfgView = v
}

// LoadTemplate loads the templates fro the view
func LoadTemplate(root string, children []string) {
	rootTemplate = root
	childTemplates = children
}

// LoadAdminTemplate loads the admin templates for the view
func LoadAdminTemplate(root string, children []string) {
	adminRootTemplate = root
	adminChildTemplates = children
}

// New creates a new View, and returns it's pointer
func New(r *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})

	v.Extension = cfgView.Extension
	v.Folder = cfgView.Folder
	v.Name = cfgView.Name

	v.request = r

	v.Vars["Auth"] = session.IsLoggedIn(r)
	v.Vars["IsTeacher"] = session.IsTeacher(r)
	v.Vars["RequestURI"] = r.RequestURI

	return v
}

// Render renders a template to the writer
func (v *View) Render(w http.ResponseWriter) {
	// Get the template collection from cache
	mutex.RLock()
	tc, ok := templateCollection[v.Name]
	mutex.RUnlock()

	pluginMutex.RLock()
	pc := pluginCollection
	pluginMutex.RUnlock()

	// If the template collection is not cached or caching is disabled
	if !ok || !cfgView.Caching {
		// Check if there was a request to render an admin-page
		adminRequest := strings.HasPrefix(v.request.RequestURI, "/admin")

		var root string
		var children []string

		// Check request
		if adminRequest {
			root = adminRootTemplate
			children = adminChildTemplates
		} else {
			root = rootTemplate
			children = childTemplates
		}

		// List of template names
		var templateList []string
		templateList = append(templateList, root)
		templateList = append(templateList, v.Name)
		templateList = append(templateList, children...)

		// Loop through each template and test the full path
		for i, name := range templateList {
			// Get the absolute path of the root template
			filePath := v.Folder + string(os.PathSeparator) + name + "." + v.Extension
			path, err := filepath.Abs(filePath)
			if err != nil {
				log.Println("filePath: ", filePath)
				http.Error(w, "template path error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			templateList[i] = path
		}

		// Determine if there is an error in the template syntax
		templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

		if err != nil {
			log.Printf("view.go error: %v", err)
			http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Cache the template collection
		mutex.Lock()
		templateCollection[v.Name] = templates
		mutex.Unlock()

		// Save the template collection
		tc = templates
	}

	// Display the content to the screen
	err := tc.Funcs(pc).ExecuteTemplate(w, rootTemplate+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "template file error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// LoadPlugins for templating
func LoadPlugins(funcMaps ...template.FuncMap) {
	funcMap := make(template.FuncMap)

	for _, m := range funcMaps {
		for key, value := range m {
			funcMap[key] = value
		}
	}

	// mutex?
	pluginCollection = funcMap
}
