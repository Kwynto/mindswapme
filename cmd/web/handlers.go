package main

import (
	"net/http"

	"github.com/Kwynto/gosession"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tD := tmplHomeData{
		isLinks: true,
		Link1:   MsmLink{Id: "1", Link: "https://github.com/"},
		Link2:   MsmLink{Id: "2", Link: "https://yandex.ru/"},
		Link3:   MsmLink{Id: "3", Link: "https://google.com/"},
		Link4:   MsmLink{Id: "4", Link: "https://youtube.com/"},
		Link5:   MsmLink{Id: "5", Link: "https://newprogress.xyz/"},
		Link6:   MsmLink{Id: "6", Link: "https://boosty.to/"},
		Link7:   MsmLink{Id: "7", Link: "https://nalog.ru/"},
		Link8:   MsmLink{Id: "8", Link: "https://github.com/Kwynto"},
		Link9:   MsmLink{Id: "9", Link: "https://github.com/Kwynto"},
		Link10:  MsmLink{Id: "10", Link: "https://github.com/Kwynto"},
		Link11:  MsmLink{Id: "11", Link: "https://github.com/Kwynto"},
		Link12:  MsmLink{Id: "12", Link: "https://github.com/Kwynto"},
		Link13:  MsmLink{Id: "13", Link: "https://github.com/Kwynto"},
		Link14:  MsmLink{Id: "14", Link: "https://github.com/Kwynto"},
		Link15:  MsmLink{Id: "15", Link: "https://github.com/Kwynto"},
		Link16:  MsmLink{Id: "16", Link: "https://github.com/Kwynto"},
		Link17:  MsmLink{Id: "17", Link: "https://github.com/Kwynto"},
		Link18:  MsmLink{Id: "18", Link: "https://github.com/Kwynto"},
		Link19:  MsmLink{Id: "19", Link: "https://github.com/Kwynto"},
	}

	_ = gosession.StartSecure(&w, r)

	app.render(w, r, "home.html", tD)
}
