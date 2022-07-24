package test

type F struct{}

func (f F) Write(p []byte) (n int, err error) { // OK
	return n, err
}

func Write(p []byte) (n int, err error) { // OK
	return n, err
}

type FF struct{}

func (f FF) Write(s string, p int) (n int, err error) { // Write arg length be must 1
	return n, err
}

type FFF struct{}

func (f FFF) Write(p []int) (n int, err error) { // Write arg length be must 1
	return n, err
}
