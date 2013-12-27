/** package gotcha contains a set of common test assertions. It doesn't
attempt to be any sort of fancy test library. Use the Assert() function
to grab an Asserter, which has assertion methods*/
package gotcha

type Failer interface {
	Fatalf(string, ...interface{})
}

type Matcher func(interface{}) bool

type Asserter struct {
	t Failer
}

func Assert(t Failer) *Asserter {
	return &Asserter{t}
}

func (a *Asserter) IsTrue(statement bool, message string) *Asserter {
	if !statement {
		a.t.Fatalf("%s. Was unexpectedly false.", message)
	}
	return a
}

func (a *Asserter) AreEqual(left, right interface{}, message string) *Asserter {
	if left != right {
		a.t.Fatalf("%s. Expected %+v to equal %+v.", message, left, right)
	}
	return a
}

func (a *Asserter) IsNil(subject interface{}, message string) *Asserter {
	if subject != nil {
		a.t.Fatalf("%s. Expected %+v to be nil", message, subject)
	}
	return a
}

func (a *Asserter) IsNotNil(subject interface{}, message string) *Asserter {
	if subject == nil {
		a.t.Fatalf("%s. Was unexepectedly nil", message)
	}
	return a
}

func (a *Asserter) Fail(message string) {
	a.t.Fatalf(message)
}

func (a *Asserter) Matches(subject interface{}, matcher Matcher, message string) {
	if !matcher(subject) {
		a.t.Fatalf(message)
	}
}
