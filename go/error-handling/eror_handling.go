func Use(o ResourceOpener, input string) (e error) {
	result, err := o()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(o, input)
		default:
			return err
		}
	}