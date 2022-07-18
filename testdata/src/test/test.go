package test

type F struct{}

func (f F) Write(s string) (n int, err error) { // want "ng"
	return n, err
}

type FF struct{}

func (f FF) Write(p []byte) (n int, err error) { // OK
	return n, err
}
