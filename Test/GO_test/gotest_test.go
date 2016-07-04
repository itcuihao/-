package gotest

import "testing"

func Test_Division_1(t *testing.T)  {
    if _, e := Division(6, 0); e == nil {
        t.Error("除法测试不通过。")
    } else {
        t.Log("第一个测试通过。")
    }
}

func Test_Division_2(t *testing.T) {
    t.Error("不通过！")
}

