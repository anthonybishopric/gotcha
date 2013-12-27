package gotcha

import "testing"

type FakeTest struct {
	receivedMessage string
}

func (f *FakeTest) Fatalf(format string, args ...interface{}) {
	f.receivedMessage = format
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

	if fake.receivedMessage != "the message" {
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

	Assert(t).IsTrue(fake.receivedMessage == "the message", "the message should not have been empty")
}
