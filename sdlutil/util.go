package sdlutil

import "log"

type Destructor interface {
	Destroy() error
}

func HandleDestroy(d Destructor) {
	if d != nil {
		err := d.Destroy()
		if err != nil {
			log.Println(err)
		}
	}
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

