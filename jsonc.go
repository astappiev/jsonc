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
		escaping := false
		if src[i] == '"' { // start of string literal
			var nl = 0
			for i = i + 1; i < len(src); i++ {
				if src[i] == '\n' { // escape new lines
					dst = append(dst, '\\', 'n')
					nl++
					continue
				}

				if src[i] == '\\' && src[i+1] == '\'' && src[i-1] != '\\' {
					continue // skip invalid escaping of single quotes
				}

				if src[i] == '\t' { // escape tabs
					dst = append(dst, '\\', 't')
					continue
				}

				if escaping == true && src[i] == '"' { // end of escaping
					dst = append(dst, '\\', '"')
					escaping = false
					continue
				}

				dst = append(dst, src[i])

				if src[i] == '"' { // possible end of string literal
					j := i - 1
					for ; ; j-- {
						if src[j] != '\\' {
							break
						}
					}
					if (j-i)%2 != 0 { // check if not escaped
						var firstCh byte
						spaces := 0
						for j := i + 1; ; j++ { // calculate spaces after quote and first non-space character
							if src[j] != ' ' && src[j] != '\t' && src[j] != '\n' && src[j] != '\r' {
								firstCh = src[j]
								break
							} else {
								spaces++
							}
						}

						// check if quote is not followed by expected character
						if firstCh != ',' && firstCh != ':' && firstCh != '}' && firstCh != ']' {
							if spaces == 0 { // if no spaces between quote and text, probably we need to escape it
								dst = append(dst[:len(dst)-1], '\\', '"')
								escaping = true
							} else { // if there are spaces, we can just assume comma is missing
								dst = append(dst, ',')
							}
						}

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
						if escaping == false {
							break
						}
					}
				}
			}
		} else if src[i] == '}' || src[i] == ']' {
			for j := len(dst) - 2; j >= 0; j-- { // remove trailing comma
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
