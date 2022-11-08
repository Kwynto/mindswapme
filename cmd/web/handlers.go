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
		Link1:   msmLink{Id: "1", Link: "https://github.com/"},
		Link2:   msmLink{Id: "2", Link: "https://yandex.ru/"},
		Link3:   msmLink{Id: "3", Link: "https://google.com/"},
		Link4:   msmLink{Id: "4", Link: "https://youtube.com/"},
		Link5:   msmLink{Id: "5", Link: "https://newprogress.xyz/"},
		Link6:   msmLink{Id: "6", Link: "https://boosty.to/"},
		Link7:   msmLink{Id: "7", Link: "https://nalog.ru/"},
		Link8:   msmLink{Id: "8", Link: "https://github.com/Kwynto"},
		Link9:   msmLink{Id: "9", Link: "https://github.com/Kwynto"},
		Link10:  msmLink{Id: "10", Link: "https://github.com/Kwynto"},
		Link11:  msmLink{Id: "11", Link: "https://github.com/Kwynto"},
		Link12:  msmLink{Id: "12", Link: "https://github.com/Kwynto"},
		Link13:  msmLink{Id: "13", Link: "https://github.com/Kwynto"},
		Link14:  msmLink{Id: "14", Link: "https://github.com/Kwynto"},
		Link15:  msmLink{Id: "15", Link: "https://github.com/Kwynto"},
		Link16:  msmLink{Id: "16", Link: "https://github.com/Kwynto"},
		Link17:  msmLink{Id: "17", Link: "https://github.com/Kwynto"},
		Link18:  msmLink{Id: "18", Link: "https://github.com/Kwynto"},
		Link19:  msmLink{Id: "19", Link: "https://github.com/Kwynto"},
	}

	_ = gosession.StartSecure(&w, r)

	app.render(w, r, "home.html", tD)
}
