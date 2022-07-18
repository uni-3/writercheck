package test

type F struct{}

//func (f F) Write(p []byte) (n int, err error) {
func (f F) Write(s string) (n int, err error) {
	return n, err
}
