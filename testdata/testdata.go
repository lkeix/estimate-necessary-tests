package testdata

func Hoge() {
	if true {
		if true {

		}
	} else if false {

	} else {

	}

	switch true {
	case true:
		// todo
	case false:
		// todo
	}
}

func Piyo() {
	for {
		if true {

		} else {
		}
	}
}

type Fuga struct {
}

func (f *Fuga) F() {
	if true {
		if true {

		} else if false {

		} else {

		}
	} else {

	}
}
