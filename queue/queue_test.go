package queue

import "testing"

func TestMain(t *testing.M) {
	t.Run()
}

func TestQueueWith2Stack(t *testing.T) {
	obj := Constructor2()
	obj.AppendTail(3)
	param := obj.DeleteHead()
	t.Logf("value: %d", param)
	param = obj.DeleteHead()
	t.Logf("value: %d", param)
	param = obj.DeleteHead()
	t.Logf("value: %d", param)

	//输入：
	//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
	//[[],[3],[],[],[]]
	//输出：
	//[null,null,3,3,3]
	//预期结果：
	//[null,null,3,-1,-1]
}
