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

func TestMore(t *testing.T) {
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

func TestGrep(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5, 6, 7)
	t.Logf("%#v", liste.elements)
	gerade := liste.Grep(func(i int, v int) bool {
		if v%2 == 0 {
			return true
		}
		return false
	})
	t.Logf("gerade: %#v", gerade.elements)
	if gerade.Len() != 3 {
		t.Errorf("ERROR: Anzahl gerade falsch: %d %#v", gerade.Len(), gerade)
	}
	ungerade := liste.Grep(func(i int, v int) bool {
		if v%2 != 0 {
			return true
		}
		return false
	})
	t.Logf("ungerade: %#v", ungerade.elements)
	if ungerade.Len() != 4 {
		t.Errorf("ERROR: Anzahl ungerade falsch: %d %#v", ungerade.Len(),
			ungerade)
	}
	alle := liste.Grep(func(i int, v int) bool {
		if v <= 100 {
			return true
		}
		return false
	})
	if alle.Len() != 7 {
		t.Errorf("ERROR: Anzahl alle falsch: %d %#v", alle.Len(),
			alle)
	}
	leer := liste.Grep(func(i int, v int) bool {
		if v == 100 {
			return true
		}
		return false
	})
	if leer.Len() != 0 {
		t.Errorf("ERROR: Anzahl leer falsch: %d %#v", leer.Len(),
			ungerade)
	}
}

func TestMap(t *testing.T) {
	liste := NewMYTYPEList(1, 2, 3, 4, 5, 6, 7)
	t.Logf("%#v", liste.elements)
	malzwei := liste.Map(func(i int, v int) int {
		return v * 2
	})
	t.Logf("%#v", malzwei.elements)
	if malzwei.Len() != 7 {
		t.Errorf("ERROR: Anzahl malzwei falsch: %d %#v", malzwei.Len(),
			malzwei)
	}
	if malzwei.ElementAt(1) != 4 {
		t.Error("ERROR: Element falsch:", malzwei.ElementAt(1))
	}
	if malzwei.ElementAt(4) != 10 {
		t.Error("ERROR: Element falsch:", malzwei.ElementAt(4))
	}

}

func TestAny(t *testing.T) {

	liste := NewMYTYPEList(1, 2, 3, 4, 5, 6, 7)
	t.Logf("%#v", liste.elements)

	gt5 := liste.Any(func(i int, v int) bool {
		return v > 5
	})
	if gt5 == false {
		t.Error("ERROR: Any gt5 falsch")
	}

	gt10 := liste.Any(func(i int, v int) bool {
		return v > 10
	})
	if gt10 == true {
		t.Error("ERROR: Any gt10 falsch")
	}

}

func TestAll(t *testing.T) {

	liste := NewMYTYPEList(1, 2, 3, 4, 5, 6, 7)
	t.Logf("%#v", liste.elements)

	gt0 := liste.All(func(i int, v int) bool {
		return v > 0
	})
	if gt0 == false {
		t.Error("ERROR: All gt0 falsch")
	}

	gt5 := liste.All(func(i int, v int) bool {
		return v > 5
	})
	if gt5 == true {
		t.Error("ERROR: All gt5 falsch")
	}

}
