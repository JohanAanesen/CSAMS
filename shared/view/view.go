package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

//View todo
type View struct {
	BaseURI   string   `json:"baseURI"`
	Extension string   `json:"extension"`
	Folder    string   `json:"folder"`
	Name      string   `json:"name"`
	Caching   bool     `json:"caching"`
	Template  Template `json:"template"`
	Vars      map[string]interface{}
	request   *http.Request
}

//Template todo
type Template struct {
	Root     string   `json:"root"`
	Children []string `json:"children"`
}

//Flash todo
type Flash struct {
	Message string
	Class   string
}

var (
	childTemplates     []string
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
	pluginCollection   = make(template.FuncMap)
	cfgView            *View
	mutexPlugins       sync.RWMutex
	mutex              sync.RWMutex
)

//LoadPlugins todo
func LoadPlugins(funcMaps ...template.FuncMap) {
	funcMap := make(template.FuncMap)

	// Loop through the maps
	for _, m := range funcMaps {
		// Loop through each key and value
		for key, value := range m {
			funcMap[key] = value
		}
	}

	mutexPlugins.Lock()
	pluginCollection = funcMap
	mutexPlugins.Unlock()
}

//Configure todo
func Configure(v *View) {
	cfgView = v
}

//ReadConfig todo
func ReadConfig() *View {
	return cfgView
}

//LoadTemplate todo
func LoadTemplate(root string, children []string) {
	rootTemplate = root
	childTemplates = children
}

//New todo
func New(r *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.Vars["AuthLevel"] = "anon"

	v.BaseURI = cfgView.BaseURI
	v.Extension = cfgView.Extension
	v.Folder = cfgView.Folder
	v.Name = cfgView.Name

	v.Vars["BaseURI"] = v.BaseURI

	v.request = r

	return v
}

// Render renders a template to the writer
func (v *View) Render(w http.ResponseWriter) {
	// Get the template collection from cache
	mutex.RLock()
	tc, ok := templateCollection[v.Name]
	mutex.RUnlock()

	// Get the plugin collection
	mutexPlugins.RLock()
	pc := pluginCollection
	mutexPlugins.RUnlock()

	// If the template collection is not cached or caching is disabled
	if !ok || !cfgView.Caching {

		// List of template names
		var templateList []string
		templateList = append(templateList, rootTemplate)
		templateList = append(templateList, v.Name)
		templateList = append(templateList, childTemplates...)

		// Loop through each template and test the full path
		for i, name := range templateList {
			// Get the absolute path of the root template
			path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
			if err != nil {
				http.Error(w, "template path error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			templateList[i] = path
		}

		// Determine if there is an error in the template syntax
		templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

		if err != nil {
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
	}
}
