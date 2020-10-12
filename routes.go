package cyoa

//func (s *server) routes() {
//	s.router.HandleFunc("/", s.HandleStart())
//}


//func (s *server) HandleStart() http.HandlerFunc {
	// the purpose of returning the HandlerFunc instead of being the HandlerFunc
	// is that if there's any prep to do that's specific to this Handler
	// it only has to be done once, not every time that url is accessed
	// not that useful in this example, just playing around with formats
//	startText := make([]string, 1)
//	startText[0] = "Welcome to \"The Little Blue Gopher\", a choose-your-own-adventure story!"

//	return func(w http.ResponseWriter, r *http.Request) {
//		for i := range startText {
//			fmt.Fprintln(w, startText[i])
//		}
//	}
//}