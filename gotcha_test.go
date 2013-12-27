package gotcha

import (
	"fmt"
	"testing"
)

type FakeTest struct {
	receivedMessage string
}

func (f *FakeTest) Fatalf(format string, args ...interface{}) {
	f.receivedMessage = fmt.Sprintf(format, args...)
}

func TestIsTruePassesOnTrue(t *testing.T) {
	fake := FakeTest{}
	Assert(&fake).IsTrue(true, "the message")

	if fake.receivedMessage != "" {
		t.Fatal("Expected test to fatal")
	}
}

func TestIsTrueFailsOnFalse(t *testing.T) {

	fake := FakeTest{}
	Assert(&fake).IsTrue(false, "the message")

	if fake.receivedMessage != "the message. Was unexpectedly false." {
		t.Fatal("Expected Assert to fail")
	}
}

func TestAssertEqualsPassesOnEquality(t *testing.T) {
	fake := FakeTest{}
	Assert(&fake).AreEqual(1, 1, "the message")

	Assert(t).IsTrue(fake.receivedMessage == "", "the message should have been empty")
}

func TestAssertEqualsFailsOnInequality(t *testing.T) {
	fake := FakeTest{}
	Assert(&fake).AreEqual(1, 2, "the message")

	Assert(t).IsTrue(fake.receivedMessage == "the message. Expected 1 to equal 2.", "the message should not have been empty. "+fake.receivedMessage)
}

func TestMatcherWillPassOnTrue(t *testing.T) {
	fake := FakeTest{}
	Assert(&fake).Matches(2, func(v interface{}) bool {
		return v == 2
	}, "the message")

	Assert(t).AreEqual("", fake.receivedMessage, "Expected the message to be empty")
}

func TestMatcherWillFailOnFalse(t *testing.T) {
	fake := FakeTest{}
	Assert(&fake).Matches(3, func(v interface{}) bool {
		return v == 2
	}, "the message")

	Assert(t).AreEqual("the message", fake.receivedMessage, "Expected the message to not be empty")
}
