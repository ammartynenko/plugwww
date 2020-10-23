package packages

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gitlab.com/Spouk/gotool/config"
	"gitlab.com/Spouk/gotool/render"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	logPrefixServer  = "[serverplugwww]"
	logServerBitMask = log.LstdFlags | log.Lshortfile
)

type Server struct {
	Mux *chi.Mux
	log *log.Logger
	R   *render.Render
	cfg *ConfigStruct
}

func NewServer(configFile string) *Server {

	//читаю конфигурационный файл
	cf := ConfigStruct{}
	c := config.NewConf("", os.Stdout)
	err := c.ReadConfig(configFile, &cf)
	if err != nil {
		log.Fatal(err)
	}

	//создаю инстанс сервера
	var s = &Server{
		log: log.New(os.Stdout, logPrefixServer, logServerBitMask),
		R:   render.NewRender(cf.TemplatePath, cf.TemplateDebug, nil, cf.TemplateDebugFatal),
		Mux: chi.NewMux(),
		cfg: &cf,
	}

	//добавляю миддлы
	s.Mux.Use(middleware.Logger)
	s.Mux.Use(middleware.Recoverer)
	s.Mux.Use(middleware.StripSlashes)

	//статичные файлы
	s.Mux.Route("/static/", func(root chi.Router) {
		workDir, _ := os.Getwd()
		c.Log.Printf("GETWD: %v\n", workDir)
		filesDir := filepath.Join(workDir, "static")
		//filesDir := "/stock/s.develop/go/src/gitlab.com/rdbaru/static"
		s.FileServer(root, "/static", "/", http.Dir(filesDir))
	})

	//404Error
	s.Mux.NotFound(s.Handler404)

	//405Error
	s.Mux.MethodNotAllowed(s.Handler405)

	//возврат инстанса
	return s
}

//статичные файлы
func (s *Server) FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

//запуск сервера без сертификатов
func (s *Server) Run() {
	s.log.Printf("starting server on `%s`\n", s.cfg.AdressHTTP)
	s.log.Fatal(http.ListenAndServe(s.cfg.AdressHTTP, s.Mux))
}

//запуск сервера с поддержкой сертификата
func (s *Server) RunTLS() {
	s.log.Printf("starting TLS server on `%s`\n", s.cfg.AdressHTTP)
	s.log.Fatal(http.ListenAndServeTLS(s.cfg.AdressHTTP, s.cfg.CertFile, s.cfg.KeyFile, s.Mux))
}
