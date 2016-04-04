package boundedstack

type waiters []*sema

func (w *waiters) get() *sema {
	if len(*w) == 0 {
		return nil
	}

	sema := (*w)[0]

	copy((*w)[0:], (*w)[1:])

	(*w)[len(*w)-1] = nil // or the zero value of T

	*w = (*w)[:len(*w)-1]

	return sema
}

func (w *waiters) put(sema *sema) {
	*w = append(*w, sema)
}
