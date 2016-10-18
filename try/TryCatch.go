package try

// TryCatch ...
type TryCatch struct {
	TryFunc   func() error
	CatchFunc func(error)
}

// Catch ...
func (t *TryCatch) Catch(catchf func(error)) *TryCatch {
	t.CatchFunc = catchf

	return t
}

// Go ...
func (t *TryCatch) Go() {
	e := t.TryFunc()

	if e != nil {
		t.CatchFunc(e)
	}
}
