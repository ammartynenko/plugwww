package packages

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	//----------------------------------------------------------------------
	// конфигурация
	//-----------------------------------------------------------------------
	ConfigStruct struct {
		gorm.Model
		//global -> статичный сегмент
		TemplatePath          string        `yaml:"templatepath"`
		TemplateDebug         bool          `yaml:"templatedebug"`
		TemplateDebugFatal    bool          `yaml:"teamplatedebugfatal"`
		AdressHTTP            string        `yaml:"adresshttp"`
		ReadTimeout           time.Duration `yaml:"readtimeout"`
		WriteTimeout          time.Duration `yaml:"writetimeout"`
		CertFile              string        `yaml:"certfile"`
		KeyFile               string        `yaml:"keyfile"`
		RedirectTrailingSlash bool          `yaml:"redirecttrailingslash"`
		RedirectFixedPath     bool          `yaml:"redirectfixedpath"`
		//Logging -> статичный сегмент
		Logfile        string        `yaml:"logfile"`
		LogTagsyslog   string        `yaml:"logtagsyslog"`
		LogPrefix      string        `yaml:"logprefix"`
		StaticPath     string        `yaml:"staticfile"`
		StaticPrefix   string        `yaml:"staticprefix"`
		FileTimerSleep time.Duration `yaml:"filetimersleep"`
		//roles
		Roles      []string `yaml:"roles"`
		RolesAdmin []string `yaml:"rolesadmin"`

		//database sqlite (hdd + ramfs)
		DBSQLiteHDD         string        `yaml:"dbsqlitehdd"`
		DBSQLiteRAM         string        `yaml:"dbsqliteram"`
		DBSqliteTimerBackup time.Duration `yaml:"dbsqlitetimer"`

		//sitemap + default category + post links
		Hostname        string `yaml:"hostname"`
		DefaultCategory string `yaml:"defaultcategory"`

		//database mysql
		DBTypeDB             string `yaml:"dbtypedb"`
		DBHost               string `yaml:"dbhost"`
		DBPort               string `yaml:"dbport"`
		DBUser               string `yaml:"dbuser"`
		DBPassword           string `yaml:"dbpassword"`
		DBDatabase           string `yaml:"dbdatabase"`
		DBSSLMode            bool   `yaml:"dbsslmode"`
		DBSetMaxIdleConns    int    `yaml:"dbsetmaxidleconns"`
		DBSetMaxOpenConns    int    `yaml:"dbsetmaxopenconns"`
		DBSetConnMaxLifetime int    `yaml:"dbsetconnmaxlifetime"`
		//project
		PaginateCountOnPage int           `yaml:"paginatecountonpage"`
		PaginateCountLinks  int           `yaml:"paginatecountlinks"`
		PaginateSortType    []string      `yaml:"paginatesorttype"`
		PaginateDebug       bool          `yaml:"paginatedebug"`
		UploadPath          string        `yaml:"uploadpath"`
		SitemapPath         string        `yaml:"sitemappath"`
		SitemapHost         string        `yaml:"sitemaphost"`
		HostFullPathHTTP    string        `yaml:"hostfullpathhttp"`
		HostFullPathHTTPS   string        `yaml:"hostfullpathhttps"`
		SeoTitle            string        `yaml:"seotitle"`
		SeoDesc             string        `yaml:"seodesc"`
		SeoKeys             string        `yaml:"seokeys"`
		SeoRobot            string        `yaml:"seorobot"`
		LaterPostTimePeriod time.Duration `yaml:"laterposttimeperiod"`
		//session
		MailTo         string `yaml:"mailto"`
		MailFrom       string `yaml:"mailfrom"`
		MailHost       string `yaml:"mailhost"`
		MailPort       int    `yaml:"mailport"`
		MailUsername   string `yaml:"mailusername"`
		MailPassword   string `yaml:"mailpassword"`
		CSRFTimeActive int    `yaml:"csrftimeactive"`
		CSRFSalt       string `yaml:"csrfsalt"`
		//Cookie part
		CookieName             string        `yaml:"cookiename"`
		CookieDomain           string        `yaml:"cookiedomain"`
		CookieExpired          int64         `yaml:"cookieexpired"`
		CookieSalt             string        `yaml:"cookiesalt"`
		CookieAnonymous        string        `yaml:"cookieanonymous"`
		CookieRegister         string        `yaml:"cookieregister"`
		RoleDefaultUser        string        `yaml:"roledefaultuser"`
		SessionTime            time.Duration `yaml:"sessiontime"`
		SessionTimeExpired     time.Duration `yaml:"sessiontimeexpired"`
		SessionTimeSave        time.Duration `yaml:"sessiontimesave"`
		SessionPathSave        string        `yaml:"pathsavesession"`
		SessionTimeSleepWorker time.Duration `yaml:"sessiontimesleepworker"`
		TimerTime              time.Duration `yaml:"timertime"`
		SleepTimeCatcher       time.Duration `yaml:"sleeptimecatcher"`
		DeferPostSleepTime     time.Duration `yaml:"deferpostsleeptime"`
		DeferPostTime          time.Duration `yaml:"deferposttime"`
		ContactReview          []string      `yaml:"contactreview"`
		FlashSalt              string        `yaml:"flashsatl"`
		RedisAdress            string        `yaml:"redisadress"`
		RedisDB                int           `yaml:"redisdb"`
		RedisPassword          string        `yaml:"redispassword"`
		DumpConfigFile         string        `yaml:"dumpconfigfile"`
	}
)
