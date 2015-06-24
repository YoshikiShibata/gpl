package concat

func Concat(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str // NB: inefficient!
		sep = " "
	}
	return s
}
