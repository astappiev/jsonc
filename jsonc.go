package jsonc

// ToJSON strips out comments and trailing commas and convert the input to a valid JSON
//
// The resulting JSON will always be the same length as the input, and it will include all the same line breaks
// at matching offsets. This is to ensure the result can be later processed by an external parser and that that
// parser will report messages or errors with the correct offsets.
func ToJSON(src []byte) []byte {
	var dst []byte
	for i := 0; i < len(src); i++ {
		if src[i] == '/' { // start of comment
			if i < len(src)-1 {
				if src[i+1] == '/' { // single line comment
					dst = append(dst, ' ', ' ')
					i += 2
					for ; i < len(src); i++ {
						if src[i] == '\n' {
							dst = append(dst, '\n')
							break
						} else if src[i] == '\t' || src[i] == '\r' {
							dst = append(dst, src[i])
						} else {
							dst = append(dst, ' ')
						}
					}
					continue
				}
				if src[i+1] == '*' { // multi line comment
					dst = append(dst, ' ', ' ')
					i += 2
					for ; i < len(src)-1; i++ {
						if src[i] == '*' && src[i+1] == '/' {
							dst = append(dst, ' ', ' ')
							i++
							break
						} else if src[i] == '\n' || src[i] == '\t' ||
							src[i] == '\r' {
							dst = append(dst, src[i])
						} else {
							dst = append(dst, ' ')
						}
					}
					continue
				}
			}
		}
		dst = append(dst, src[i])
		if src[i] == '"' { // start of string literal
			var nl = 0
			for i = i + 1; i < len(src); i++ {
				if src[i] == '\n' { // escape new lines
					dst = append(dst, '\\', 'n')
					nl++
				} else {
					dst = append(dst, src[i])
				}

				if src[i] == '"' { // possible end of string literal
					j := i - 1
					for ; ; j-- {
						if src[j] != '\\' {
							break
						}
					}
					if (j-i)%2 != 0 { // check if not escaped
						if nl > 0 { // if new lines were escaped, add them back
							if src[i+1] == ',' { // finish line with comma first
								i++
								dst = append(dst, src[i])
							}
							for nl > 0 {
								nl--
								dst = append(dst, '\n')
							}
						}
						break
					}
				}
			}
		} else if src[i] == '}' || src[i] == ']' {
			for j := len(dst) - 2; j >= 0; j-- {
				if dst[j] <= ' ' {
					continue
				}
				if dst[j] == ',' {
					dst[j] = ' '
				}
				break
			}
		}
	}
	return dst
}
