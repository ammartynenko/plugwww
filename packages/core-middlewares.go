package packages

//import (
//	"context"
//	"gitlab.com/Spouk/gotool/session"
//	"net/http"
//	"time"
//)

const (
	keyCountConnect = "user_count_connect"
)

//func (s *Server) MiddlewareSession(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//переменные
//		var (
//			//user      = &User{}
//			ses       = s.Pool.Get().(*session.Session)
//			newCookie = http.Cookie{}
//		)
//		//обработка кукиса
//		if c, err := r.Cookie(s.cfg.CookieName); err != nil {
//			newS := &SessionTable{}
//			//кукис не найден -> создаю новую сессию и передаю ее в контекст (1)
//			newCookie = s.Support.NewCook(s.cfg.CookieName, s.cfg.CookieSalt, r)
//			res, err := ses.TextToJSON() //конвертация сессии в текстовое представление для сохранения в базе данных
//			if err != nil {
//				s.log.Println(err)
//			} else {
//				newS.Session = string(res)                        //сохраняю строковое представление объекта сессиии в строке в базу данных
//				newS.Active = true                                //выставление сессии статуса активная
//				newS.Cook = newCookie.Value                       //присваивание пользовательского кукиса объекту сессии
//				newS.LastConnect = time.Now()                     //время последнего подключение = время текущего подключения
//				if err := s.dbs.db.Save(newS).Error; err != nil { //сохранение сессииВтаблице в базе данных
//					s.log.Println(err)
//				}
//			}
//			http.SetCookie(w, &newCookie)
//		} else {
//			//найден кукис -> приоритет :: пользователь > старая сессия
//			oldS, user := s.findOldSessionAndUser(c.Value)
//
//			//если старая сессию не найдена а пользователь есть -> создаю новую сессию с привязкой к кукису пользователя [2]
//			if oldS == nil && user != nil {
//				newCookie.Value = user.Cook  //присваиваю пользовательский кукис объект кукис
//				res, err := ses.TextToJSON() //конвертация сессии в текстовое представление для сохранения в базе данных
//				if err != nil {
//					s.log.Println(err)
//				} else {
//					newS := &SessionTable{}
//					newS.Session = string(res)                        //сохраняю строковое представление объекта сессиии в строке в базу данных
//					newS.Active = true                                //выставление сессии статуса активная
//					newS.Cook = user.Cook                             //присваивание пользовательского кукиса объекту сессии
//					newS.LastConnect = time.Now()                     //время последнего подключение = время текущего подключения
//					if err := s.dbs.db.Save(newS).Error; err != nil { //сохранение сессииВтаблице в базе данных
//						s.log.Println(err)
//					}
//				}
//			}
//			//если старая сессия не найдена и пользователя нет -> создаю новую сессию [1]
//			if oldS == nil && user == nil {
//				newCookie = s.Support.NewCook(s.cfg.CookieName, s.cfg.CookieSalt, r) //генерирую новый кукис
//				http.SetCookie(w, &newCookie)                                        //устанавливаю кукис пользователю
//				res, err := ses.TextToJSON()                                         //конвертация сессии в текстовое представление для сохранения в базе данных
//				if err != nil {
//					s.log.Println(err)
//				} else {
//					newS := &SessionTable{}
//					newS.Session = string(res)                        //сохраняю строковое представление объекта сессиии в строке в базу данных
//					newS.Active = true                                //выставление сессии статуса активная
//					newS.Cook = newCookie.Value                       //присваивание пользовательского кукиса объекту сессии
//					newS.LastConnect = time.Now()                     //время последнего подключение = время текущего подключения
//					if err := s.dbs.db.Save(newS).Error; err != nil { //сохранение сессииВтаблице в базе данных
//						s.log.Println(err)
//					}
//				}
//			}
//			//если старая сессия найдена и пользователь тоже -> восстанавливаю старую сессию
//			if (oldS != nil && user != nil) || (oldS != nil && user == nil) {
//				newText, err := ses.TextFROMJSON([]byte(oldS.Session)) //извлекаю сессию из базы данные и конвертирую в бинарное представление
//				if err != nil {
//					s.log.Println(err)
//				} else {
//					ses.TEXT = newText //замещаю в новой сессии текстовой сегмент восстановленной сессии
//				}
//				newCookie.Value = c.Value //обновляю кукис
//				oldS.Cook = c.Value
//				oldS.LastConnect = time.Now()                      //время последнего подключения
//				oldS.Active = true                                 //выставляю сессию активной
//				if err := s.dbs.db.Save(&oldS).Error; err != nil { //сохранение сессииВтаблице в базе данных
//					s.log.Println(err)
//				}
//			}
//			//передача в контекст пользователя
//			r = r.WithContext(context.WithValue(r.Context(), "user", user))
//		}
//		//флешер
//		var newFlash = session.FLASH{}
//		cFlasher, exists := s.flash.Stock[newCookie.Value]
//		if exists == false {
//			//не найден кукис в стеке флешеров, создаю новый
//			s.flash.Stock[newCookie.Value] = newFlash
//		} else {
//			newFlash = cFlasher
//		}
//		//добавляю флешер в контекст
//		ses.SetDATA("stock", "flasher", newFlash)
//		//общее для всех ситуаций
//		ses.SetTEXT("info", keyCountConnect, 0)
//		ses.SetTEXT("info", "referer", []string{r.Referer()})
//		ses.SetTEXT("info", "user_requesturi", r.RequestURI)
//		ses.SetTEXT("info", "user_remoteaddr", r.RemoteAddr)
//		ses.SetTEXT("info", "user_headers", r.Header)
//		ses.SetTEXT("info", "user_cookies", r.Cookies())
//		ses.SetTEXT("info", "user_referer", r.Referer())
//		ses.SetTEXT("info", "user_useragent", r.UserAgent())
//		ses.SetTEXT("info", "user_timeconnect", time.Now().String())
//		ses.SetTEXT("info", "user_cook", newCookie.Value)
//		//передача в контекст сессии
//		r = r.WithContext(context.WithValue(r.Context(), "session", ses))
//		//передача контекста далее
//		next.ServeHTTP(w, r)
//	})
//}
