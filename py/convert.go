package py

import (
	"strings"
	"unicode"

	gopinyin "github.com/mozillazg/go-pinyin"
)

// Options 定义转换选项。
type Options struct {
	Upper      bool
	Lower      bool
	KeepNonHan bool
}

// ConvertLine 将一行文本中的汉字转换为拼音首字母。
func ConvertLine(line string, opts Options) string {
	a := gopinyin.NewArgs()
	a.Style = gopinyin.FirstLetter

	lowerEnabled := opts.Lower
	if opts.Upper {
		lowerEnabled = false
	}

	var sb strings.Builder
	for _, r := range line {
		if unicode.Is(unicode.Han, r) {
			p := gopinyin.SinglePinyin(r, a)
			if len(p) > 0 && len(p[0]) > 0 {
				letter := string(p[0][0])
				if opts.Upper {
					letter = strings.ToUpper(letter)
				} else if lowerEnabled {
					letter = strings.ToLower(letter)
				}
				sb.WriteString(letter)
			}
		} else if opts.KeepNonHan {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
