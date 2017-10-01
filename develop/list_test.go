package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	liste := NewMYTYPEList()
	kapazität := cap(liste.elements)
	if kapazität != 0 {
		t.Error("ERROR: cap falsch:", kapazität)
	}
	liste.Push(1)
	kapazität = cap(liste.elements)
	if kapazität != 1 {
		t.Error("ERROR: cap falsch:", kapazität)
	}
	liste.Push(2)
	liste.Push(3)
	kapazität = cap(liste.elements)
	if kapazität != 4 {
		t.Error("ERROR: cap falsch:", kapazität)
	}
	t.Logf("%#v", liste.elements)
	if liste.Len() != 3 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	liste.Push(4)
	liste.Push(5)
	kapazität = cap(liste.elements)
	if kapazität != 8 {
		t.Error("ERROR: cap falsch:", kapazität)
	}

	value := liste.elements[2]
	if value != 3 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	t.Logf("%#v", liste.elements)
}

func TestUnshift(t *testing.T) {
	liste := NewMYTYPEList()
	liste.Unshift(6)
	liste.Unshift(7)
	liste.Unshift(8)
	liste.Unshift(9)
	t.Logf("%#v", liste.elements)
	if liste.Len() != 4 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	value := liste.elements[0]
	if value != 9 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	liste.Unshift(5)
	value = liste.elements[0]
	if value != 5 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	liste.Unshift(4)
	value = liste.elements[0]
	if value != 4 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	liste.Unshift(3)
	value = liste.elements[0]
	if value != 3 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	t.Logf("%#v", liste.elements)
}

func TestShift(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5)
	t.Logf("%#v", liste.elements)
	value := liste.Shift()
	if value != 1 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 4 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value = liste.Shift()
	if value != 2 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 3 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value = liste.Shift() // 3
	value = liste.Shift() // 4
	value = liste.Shift() // 5
	value = liste.Shift() // empty
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value = liste.Shift() // empty
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
}

func TestShiftChecked(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5)
	t.Logf("%#v", liste.elements)
	value, ok := liste.ShiftChecked()
	if !ok {
		t.Errorf("ERROR: flag false: (%#v)", liste.elements)
	}
	if value != 1 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 4 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value, ok = liste.ShiftChecked()
	if !ok {
		t.Errorf("ERROR: flag false: (%#v)", liste.elements)
	}
	if value != 2 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 3 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}

	for i := 0; i < 4; i++ {
		value, ok = liste.ShiftChecked()
	}
	if ok {
		t.Errorf("ERROR: flag false: (%#v)", liste.elements)
	}
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value, ok = liste.ShiftChecked()
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	if ok {
		t.Errorf("ERROR: flag false: (%#v)", liste.elements)
	}
}

func TestPopChecked(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5)
	t.Logf("%#v", liste.elements)
	value, ok := liste.PopChecked()
	if value != 5 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 4 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	if !ok {
		t.Errorf("ERROR: PopChecked false: (%#v)", liste.elements)
	}
	t.Logf("%#v", liste.elements)
	value, ok = liste.PopChecked()
	if value != 4 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 3 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	if !ok {
		t.Errorf("ERROR: PopChecked false: (%#v)", liste.elements)
	}

	for i := 0; i < 4; i++ {
		value, ok = liste.PopChecked()
	}
	if ok {
		t.Errorf("ERROR: PopChecked true: (%#v)", liste.elements)
	}
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value, ok = liste.PopChecked()
	if ok {
		t.Errorf("ERROR: PopChecked true: (%#v)", liste.elements)
	}
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
}
func TestPop(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5)
	t.Logf("%#v", liste.elements)
	value := liste.Pop()
	if value != 5 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 4 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	t.Logf("%#v", liste.elements)
	value = liste.Pop()
	if value != 4 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 3 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value = liste.Pop()
	value = liste.Pop()
	value = liste.Pop()
	value = liste.Pop()
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
	value = liste.Pop()
	if value != 0 {
		t.Errorf("ERROR: value falsch: %d (%#v)", value, liste.elements)
	}
	if liste.Len() != 0 {
		t.Errorf("ERROR: länge falsch: %d (%#v)", liste.Len(), liste.elements)
	}
}

func TestAll(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5)
	t.Logf("%#v", liste.elements)
	if liste.Len() != 5 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	wert := liste.Pop()
	if wert != 5 {
		t.Error("ERROR: Wert 5 falsch:", wert)
	}
	t.Logf("%#v", liste.elements)
	wert = liste.Shift()
	if wert != 1 {
		t.Error("ERROR: Wert 1 falsch:", wert)
	}
	t.Logf("%#v", liste.elements)
	if liste.Len() != 3 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	liste.Unshift(6)
	t.Logf("%#v", liste.elements)
	if liste.Len() != 4 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	for i := 0; i < 6; i++ {
		wert = liste.Pop()
	}
	t.Logf("%#v", liste.elements)
	if wert != 0 {
		t.Error("ERROR: Wert 0 falsch:", wert)
	}
	wert = liste.Shift()
	if wert != 0 {
		t.Error("ERROR: Wert 0 falsch:", wert)
	}
}

func TestForeach(t *testing.T) {
	liste := NewMYTYPEList(10, 22, 44, 33, 55)
	t.Logf("%#v", liste.elements)
	if liste.Len() != 5 {
		t.Error("ERROR: Anzahl falsch:", liste.Len())
	}
	liste.Foreach(func(i int, w int) {
		t.Logf("index: %d, wert: %d\n", i, w)
	})
}

func TestElementAt(t *testing.T) {
	liste := NewMYTYPEList(10, 22, 44, 33, 55)
	t.Logf("%#v", liste.elements)
	value := liste.ElementAt(3)
	if value != 33 {
		t.Error("ERROR: value falsch:", value)
	}
	callChecked := func() {
		defer func() {
			msg := recover()
			if msg != nil {
				t.Log("OK: panic: ", msg)
			} else {
				t.Error("ERROR: keine panic:")
			}
		}()
		value = liste.ElementAt(7)
		if value != 33 {
			t.Error("ERROR: value falsch:", value)
		}
	}
	callChecked()
}
