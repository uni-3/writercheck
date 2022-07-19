package test

type F struct{} // OK

func (f F) Write(s string) (n int, err error) {
	return n, err
}

func (f F) String() string {
	return "test"
}

type FF struct{} // OK

func (f FF) Write(p []byte) (n int, err error) {
	return n, err
}
