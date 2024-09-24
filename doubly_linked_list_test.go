package dl_list

import (
	"testing"
)

// This is how your code will be called.
// Your answer should respect the linked list order.
// You can edit this code to try different testing cases.

func TestDoublyLinkedList1(t *testing.T) {
	testCases := []struct {
    index int
    value string
	}{
	{index: 0, value: "C"},
	}
	dl := &DoublyLinkedList[string]{}
	err := dl.AddElements(testCases)
	forwardPrint:= dl.PrintForward()
	reversePrint:= dl.PrintReverse()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fexpected := "HEAD -> C -> NULL"
	if forwardPrint != fexpected {
		t.Errorf("Expected: %v,Forward print: %v", fexpected, forwardPrint)
	}

	rexpected := "NULL <- C <- HEAD"
	if reversePrint !=  rexpected{
		t.Errorf("Expected: %v,Reverse print: %v", rexpected, reversePrint)
	}
}

func TestDoublyLinkedList2(t *testing.T) {
	testCases := []struct {
    index int
    value string
	}{
		{index: 0, value: "C"},
		{index: 0, value: "A"},
	}
	dl := &DoublyLinkedList[string]{}
	err := dl.AddElements(testCases)
	forwardPrint:= dl.PrintForward()
	reversePrint:= dl.PrintReverse()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fExpected := "HEAD -> A -> NULL"
	if forwardPrint != fExpected {
		t.Errorf("Expected: %v,Forward print: %v", fExpected, forwardPrint)
	}

	rExpected := "NULL <- A <- HEAD"
	if reversePrint !=  rExpected{
		t.Errorf("Expected: %v,Reverse print: %v", rExpected, reversePrint)
	}
}

func TestDoublyLinkedList3(t *testing.T) {
	testCases := []struct {
		index int
		value string
		}{
			{index: 1, value: "This value should not be added"},
		}
		dl := &DoublyLinkedList[string]{}
		err := dl.AddElements(testCases)
		
		eExpected := "index exceeds list size"
		if err == nil {
			t.Errorf("Expected index exceeds list size error, got nil") 
		} else if err.Error() != eExpected {
			t.Errorf("Expected Error: %v, Error: %v", eExpected, err)
		}
}



// testCases := []struct {
//     index int
//     value string
// }{
//    {index: 0, value: "C"},
//    {index: 0, value: "A"},
//    {index: 1, value: "B"},
//    {index: 3, value: "D"},
// }
// dl := &DoublyLinkedList[string]{}
// err := dl.AddElements(testCases)
// forwardPrint:= dl.PrintForward()
// reversePrint:= dl.PrintReverse()