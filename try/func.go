package try

// Call ...
func Call(tryf func() error) *TryCatch {
	tc := &TryCatch{
		TryFunc: tryf,
	}

	return tc
}
