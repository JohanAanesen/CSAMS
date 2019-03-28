package config

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/email"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/server"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view/plugin"
	"io/ioutil"
	"log"
	"os"
)

// Configuration struct
type Configuration struct {
	Database *db.MySQLInfo  `json:"database"`
	Server   *server.Server `json:"server"`

	View *view.View `json:"view"`

	Session *session.Session `json:"session"`
	Email   *email.SMTPInfo  `json:"email"`

	Template      *view.Template `json:"template"`
	TemplateAdmin *view.Template `json:"template_admin"`
}

// Load a JSON file making a Configuration pointer
func Load(configFile string) (*Configuration, error) {
	// Open file
	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("could not open config file: %v\n", err)
		return &Configuration{}, err
	}

	// Close file
	defer file.Close()
	// Read file
	bytes, _ := ioutil.ReadAll(file)

	// Declare Configuration
	cfg := Configuration{}
	// Unmarshal JSON to Configuration-object
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Printf("could not unmarshal json: %v\n", err)
		return &Configuration{}, err
	}

	return &cfg, nil
}

// Initialize the configuration
func Initialize(configFile string) *Configuration {
	var cfg = &Configuration{}

	cfg, err := Load(configFile)
	if err != nil {
		log.Printf("could not load config file: %v", err.Error())
		return nil
	}

	// Configure Session
	session.Configure(cfg.Session)

	// Configure Database
	db.ConfigureDB(cfg.Database)
	db.OpenDB()

	// Configure View
	view.Configure(cfg.View)
	view.LoadTemplate(cfg.Template.Root, cfg.Template.Children)
	view.LoadAdminTemplate(cfg.TemplateAdmin.Root, cfg.TemplateAdmin.Children)
	view.LoadPlugins(
		plugin.PrettyTime(),
		plugin.DeadlineDue(),
		plugin.MDConvert(),
		plugin.Increment(),
		plugin.SplitChoices(),
		plugin.Atoi(),
		plugin.GetUsername(),
		plugin.SplitString(),
		plugin.Itoa(),
		plugin.Contains(),
		plugin.HasPrefix(),
	)

	return cfg
}
