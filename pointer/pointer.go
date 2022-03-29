package pointer

type myStr struct {
	Data  *string
	Data2 *string
}

func Run() {
	entry()
}

func entry() {
	s := "somedata"
	s2 := "somedata2"
	str := myStr{
		Data:  &s,
		Data2: &s2,
	}
	service(str)
}

func service(data myStr) {
	s := "updated"
	s2 := "updated2"
	data.Data = &s
	data.Data2 = &s2
}

func RunRef() {
	entryRef()
}

func entryRef() {
	s := "somedata"
	s2 := "somedata2"
	str := myStr{
		Data:  &s,
		Data2: &s2,
	}
	serviceP(&str)
}

func serviceP(data *myStr) {
	s := "updated"
	s2 := "updated2"
	data.Data = &s
	data.Data2 = &s2
}

func RunNew() {
	entryNew()
}

func entryNew() {
	str := new(myStr)
	s := "somedata"
	s2 := "somedata2"
	str.Data = &s
	str.Data2 = &s2
	serviceNew(str)
}

func serviceNew(data *myStr) {
	s := new(string)
	*s = "updated"
	s2 := new(string)
	*s2 = "updated2"
	data.Data = s
	data.Data2 = s2
}

func RunP() {
	entryP()
}

func entryP() {
	var str *myStr
	s := "somedata"
	s2 := "somedata2"
	str = &myStr{
		Data:  &s,
		Data2: &s2,
	}
	serviceP(str)
}

type myStrN struct {
	Data  string
	Data2 string
}

func RunN() {
	entryN()
}

func entryN() {
	s := "somedata"
	s2 := "somedata2"
	str := myStrN{
		Data:  s,
		Data2: s2,
	}
	serviceN(str)
}

func serviceN(data myStrN) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = s2
}

func RunNRef() {
	entryNRef()
}

func entryNRef() {
	s := "somedata"
	s2 := "somedata2"
	str := myStrN{
		Data:  s,
		Data2: s2,
	}
	serviceNRef(&str)
}

func serviceNRef(data *myStrN) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = s2
}

func RunNNew() {
	entryNNew()
}

func entryNNew() {
	str := new(myStrN)
	s := "somedata"
	s2 := "somedata2"
	str.Data = s
	str.Data2 = s2
	serviceNNew(str)
}

func serviceNNew(data *myStrN) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = s2
}

func RunNP() {
	entryNP()
}

func entryNP() {
	var str *myStrN
	s := "somedata"
	s2 := "somedata2"
	str = &myStrN{
		Data:  s,
		Data2: s2,
	}
	serviceNP(str)
}

func serviceNP(data *myStrN) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = s2
}

type myStrM struct {
	Data  string
	Data2 *string
}

func RunM() {
	entryM()
}

func entryM() {
	s := "somedata"
	s2 := "somedata2"
	str := myStrM{
		Data:  s,
		Data2: &s2,
	}
	serviceM(str)
}

func serviceM(data myStrM) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = &s2
}

func RunMRef() {
	entryMRef()
}

func entryMRef() {
	s := "somedata"
	s2 := "somedata2"
	str := myStrM{
		Data:  s,
		Data2: &s2,
	}
	serviceMRef(&str)
}

func serviceMRef(data *myStrM) {
	s := "updated"
	s2 := "updated2"
	data.Data = s
	data.Data2 = &s2
}

func RunMNew() {
	entryMNew()
}

func entryMNew() {
	str := new(myStrM)
	s := "somedata"
	s2 := new(string)
	*s2 = "somedata2"
	str.Data = s
	str.Data2 = s2
	serviceMNew(str)
}

func serviceMNew(data *myStrM) {
	s := "updated"
	s2 := new(string)
	*s2 = "updated2"
	data.Data = s
	data.Data2 = s2
}

func RunMP() {
	entryMP()
}

func entryMP() {
	var str *myStrM
	s := "somedata"
	var s2 *string
	s2t := "somedata2"
	s2 = &s2t
	str = &myStrM{
		Data:  s,
		Data2: s2,
	}
	serviceMP(str)
}

func serviceMP(data *myStrM) {
	s := "updated"
	var s2 *string
	s2t := "updated"
	s2 = &s2t
	data.Data = s
	data.Data2 = s2
}
