package index1

import "testing"

func TestMethod1(t *testing.T) {
	sum := Method1(10)

	if sum != 55 {
		t.Log("测试数据不符合预期")
		t.FailNow()
	}
	t.Log("测试成功")
}
