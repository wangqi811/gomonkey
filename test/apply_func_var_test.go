package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "gitlab.prod.dtstack.cn/dt-insight-ops/gomonkey/v2"
	"gitlab.prod.dtstack.cn/dt-insight-ops/gomonkey/v2/test/fake"
)

func TestApplyFuncVar(t *testing.T) {
	Convey("TestApplyFuncVar", t, func() {

		Convey("for succ", func() {
			str := "hello"
			patches := ApplyFuncVar(&fake.Marshal, func(_ interface{}) ([]byte, error) {
				return []byte(str), nil
			})
			defer patches.Reset()
			bytes, err := fake.Marshal(nil)
			So(err, ShouldEqual, nil)
			So(string(bytes), ShouldEqual, str)
		})

		Convey("for fail", func() {
			patches := ApplyFuncVar(&fake.Marshal, func(_ interface{}) ([]byte, error) {
				return nil, fake.ErrActual
			})
			defer patches.Reset()
			_, err := fake.Marshal(nil)
			So(err, ShouldEqual, fake.ErrActual)
		})
	})
}
