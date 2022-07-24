package test

// OK
type F struct{}

func (f F) Write(p []byte) (n int, err error) { // OK
	return n, err
}

func Write(p []byte) (n int, err error) { // OK
	return n, err
}

// 引数
type FF struct{}

func (f FF) Write(s string, p int) (n int, err error) { // Write arg length be must 1
	return n, err
}

type FFF struct{}

func (f FFF) Write(p []int) (n int, err error) { // Write arg length be must 1
	return n, err
}

// 返り値
type FFFF struct{}

func (f FFFF) Write(p []byte) (err error) { // Write arg length be must 1
	return err
}

type FFFFF struct{}

func (f FFFFF) Write(p []byte) (n string, err error) { // Write arg length be must 1
	return n, err
}

type FFFFFF struct{}

func (f FFFFFF) Write(p []byte) (num int, err error) { // Write arg length be must 1
	return num, err
}
