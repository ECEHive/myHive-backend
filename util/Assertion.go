package util

import "github.com/sirupsen/logrus"

type AssertionUtil struct {
	entry *logrus.Entry
}

var Assert *AssertionUtil = nil

func (v *AssertionUtil) assert(assert bool, errorMsg string) {
	if v == nil {
		Assert = &AssertionUtil{GetLogger("Assertion")}
	}
	v = Assert
	if !assert {
		v.entry.Panicf("PANIC: Assertion ERROR: %s", errorMsg)
		panic(errorMsg)
	}
}
