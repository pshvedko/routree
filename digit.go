package routree

type Digit uint16

func (d Digit) Split() []byte {
	var b []byte
	var i byte
	for d != 0 {
		if d&1 != 0 {
			switch i {
			case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
				b = append(b, byte('0'+i))
			case 15:
				b = append(b, '*')
			case 14:
				b = append(b, '#')
			default:
				return nil
			}
		}
		d >>= 1
		i++
	}
	return b
}
