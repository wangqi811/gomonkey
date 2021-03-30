package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "gitlab.prod.dtstack.cn/dt-insight-ops/gomonkey/v2"
)

var num = 10

func TestApplyGlobalVar(t *testing.T) {
	Convey("TestApplyGlobalVar", t, func() {

		Convey("change", func() {
			patches := ApplyGlobalVar(&num, 150)
			defer patches.Reset()
			So(num, ShouldEqual, 150)
		})

		Convey("recover", func() {
			So(num, ShouldEqual, 10)
		})
	})
}
