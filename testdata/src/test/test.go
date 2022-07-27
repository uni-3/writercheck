package test

// OK
type correct struct{}

func (f correct) Write(p []byte) (n int, err error) { // OK
	return n, err
}

func Write(p []byte) (n int, err error) { // OK
	return n, err
}

// test pattern invalid arguments
type incorrectWithArgLength struct{}

func (f incorrectWithArgLength) Write(s string, p int) (n int, err error) { // want "Write's argument length is '2' must be 1"
	return n, err
}

type incorrectWithFirstArgType struct{}

func (f incorrectWithFirstArgType) Write(p []int) (n int, err error) { // want "p argument is invalid type 'int' must be 'byte'"
	return n, err
}

type incorrectWithFirstArgName struct{}

func (f incorrectWithFirstArgName) Write(pp []byte) (n int, err error) { // want "Write's argument name is 'pp' must be 'p'"
	return n, err
}

// test pattern invalid return values
type incorrectWithResLength struct{}
type FFFF struct{}

func (f incorrectWithResLength) Write(p []byte) (err error) { // want "Write returns length '1' must be 2"
	return err
}

type incorrectWithFirstResType struct{}

func (f incorrectWithFirstResType) Write(p []byte) (n string, err error) { // want "Write first return type is 'string' must be 'int'"
	return n, err
}

type incorrectWithFirstResName struct{}

func (f incorrectWithFirstResName) Write(p []byte) (num int, err error) { // want "Write first return name is 'num' must be 'n'"
	return num, err
}
