package cmdb

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	ErrArgumentsMissing        = errors.New("arguments missing")
	ErrFailedToCreateInfoCMDB  = errors.New("failed to create infocmdb object")
	ErrNoCredentials           = errors.New("must provide credentials")
	ErrNotImplemented          = errors.New("not implemented")
	ErrNoResult                = errors.New("query returned no result")
	ErrTooManyResults          = errors.New("query returned to many results, expected one")
	ErrWebserviceResponseNotOk = errors.New("webservice response was not ok")
)

type Config struct {
	ApiUrl       string `yaml:"apiUrl"`
	ApiUser      string `yaml:"apiUser"`
	ApiPassword  string `yaml:"apiPassword"`
	ApiKey       string
	CmdbBasePath string `yaml:"CmdbBasePath"`
}

type InfoCMDB struct {
	Config Config
	Cache  *cache.Cache
}

type CI_RELATION_DIRECTION string

const (
	CI_RELATION_DIRECTION_ALL             CI_RELATION_DIRECTION = "all"
	CI_RELATION_DIRECTION_DIRECTED_FROM                         = "directed_from"
	CI_RELATION_DIRECTION_DIRECTED_TO                           = "directed_to"
	CI_RELATION_DIRECTION_BIDIRECTIONAL                         = "bidirectional"
	CI_RELATION_DIRECTION_OMNIDIRECTIONAL                       = "omnidirectional"
)

type ATTRIBUTE_VALUE_TYPE string

const (
	ATTRIBUTE_VALUE_TYPE_TEXT    ATTRIBUTE_VALUE_TYPE = "value_text"
	ATTRIBUTE_VALUE_TYPE_DATE                         = "value_date"
	ATTRIBUTE_VALUE_TYPE_DEFAULT                      = "value_default"
	ATTRIBUTE_VALUE_TYPE_CI                           = "value_ci"
)

func init() {
	log.SetLevel(log.InfoLevel)
	if os.Getenv("WORKFLOW_DEBUGGING") == "true" {
		log.SetLevel(log.DebugLevel)
	}
}

func (i *InfoCMDB) LoadConfigFile(configFile string) (err error) {
	_, err = os.Stat(configFile)
	if err == nil {
		log.Debugf("ConfigFile found with given string: %s", configFile)
	} else {
		WorkflowConfigPath := filepath.Dir(os.Getenv("WORKFLOW_CONFIG_PATH"))
		log.Debugf("WORKFLOW_CONFIG_PATH: %s", WorkflowConfigPath)
		configFile = filepath.Join(WorkflowConfigPath, configFile)
	}

	log.Debugf("ConfigFile: %s", configFile)

	_, err = os.Stat(configFile)
	if err != nil {
		return
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}

	return i.LoadConfig(yamlFile)
}

func (i *InfoCMDB) LoadConfig(config []byte) (err error) {
	return yaml.Unmarshal(config, &i.Config)
}

func NewCMDB(config string) (i *InfoCMDB, err error) {
	i = new(InfoCMDB)
	err = i.LoadConfigFile(config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	i.Cache = cache.New(5*time.Minute, 10*time.Minute)

	return i, nil
}

func (i *InfoCMDB) Login() error {
	if i.Config.ApiKey != "" {
		log.Debug("already logged in")
		return nil
	}

	if i.Config.ApiUser == "" {
		return ErrNoCredentials
	}
	return i.LoginWithUserPass(i.Config.ApiUrl, i.Config.ApiUser, i.Config.ApiPassword)
}
