package packages

import "net/http"

//404
func (s *Server) Handler404(w http.ResponseWriter, r *http.Request) {
	_ = s.R.Render("404.html", nil, w)

}

//405
func (s *Server) Handler405(w http.ResponseWriter, r *http.Request) {
	_ = s.R.Render("405.html", nil, w)
}

//root
func (s *Server) HandlerRoot(w http.ResponseWriter, r *http.Request) {

	if err := s.R.Render("index.html", nil, w); err != nil {
		s.log.Println(err)
	}
}
