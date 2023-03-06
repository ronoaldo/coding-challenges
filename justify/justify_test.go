package justify

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// AddSpaces adiciona espaços em branco para preencher 'word' até que fique com
// maxWidth, preferindo adicionar espaços à esquerda.
func AddSpaces(word []rune, maxWidth int) []rune {
	pos := 0
	for len(word) < maxWidth {
		if word[pos] == ' ' {
			// Adiciona mais um espaço no meio
			suff := word[pos:]
			word = append(word[:pos], ' ')
			word = append(word, suff...)
			// Go to next word
			for word[pos] == ' ' && pos < len(word) {
				pos += 1
			}
		} else {
			pos += 1
		}

		if pos >= len(word) {
			pos = 0
		}
	}
	return word
}

// Justify justifica o array de palavras 'words' deixando cada linha em 'dst'
// com exatamente 'maxWidth' characteres unicode.
func Justify(words []string, maxWidth int) (dst []string) {
	// Inicializa a linha
	line := []rune(words[0])
	wc := 1
	if len(line) > maxWidth {
		return []string{"word[0] > maxWidth"}
	}

	for i := 1; i < len(words); i++ {
		// Tratando a palavra como um array de caracteres unicode,
		// para não quebrar a acentuação (string == []byte != []rune)
		w := []rune(words[i])

		// Reporta erros, se houver
		if len(w) > maxWidth {
			return []string{fmt.Sprintf("word[%d] > maxWidth", i)}
		}

		if len(line)+1+len(w) > maxWidth {
			// Linha fica muito grande, então faz o padding e começa uma nova
			line = AddSpaces(line, maxWidth)
			dst = append(dst, string(line))
			line = w
			wc = 1
		} else {
			// Ainda cabem mais palavras, concatena os caracteres com um novo espaço
			line = append(line, ' ')
			line = append(line, w...)
			wc++
		}
	}

	// Para a última linha, apenas aplica o padding no final
	if len(line) != 0 {
		if len(line) < maxWidth {
			padding := strings.Repeat(" ", maxWidth-len(line))
			line = append(line, []rune(padding)...)
		}
		dst = append(dst, string(line))
	}
	return
}

func TestJustify(t *testing.T) {
	testCases := []struct {
		desc     string
		words    []string
		maxWidth int
		out      []string
	}{
		{
			"exemplo 1",
			[]string{"Este", "é", "um", "exemplo", "de", "justificação", "de", "texto."},
			15,
			[]string{
				"Este    é    um",
				"exemplo      de",
				"justificação de",
				"texto.         "},
		},
		{
			"exemplo 2",
			[]string{"Este", "challenge", "é", "muito", "difícil.", "Pelo",
				"visto", "vou", "quebrar", "a", "cabeça", "de", "tanto", "pensar!"},
			20,
			[]string{
				"Este   challenge   é",
				"muito  difícil. Pelo",
				"visto  vou quebrar a",
				"cabeça    de   tanto",
				"pensar!             "},
		},
		{
			"exemplo 3",
			[]string{"Este", "é", "o", "meu", "challenge", "favorito", "até", "agora!!!"},
			44,
			[]string{
				"Este é o meu challenge favorito até agora!!!",
			},
		},
		{
			"exemplo 4",
			[]string{"Paranguaripotiromiruaru"},
			23,
			[]string{"Paranguaripotiromiruaru"},
		},
		{
			"primeira palavra maior que maxWidth",
			[]string{"Pneumoultramicroscopicossilicovulcanoconiótico"},
			10,
			[]string{"word[0] > maxWidth"},
		},
		{
			"palavra maior que maxWidth no final",
			[]string{"Não", "sei", "o", "que", "é",
				"Pneumoultramicroscopicossilicovulcanoconiótico"},
			10,
			[]string{"word[5] > maxWidth"},
		},
	}

	for _, tc := range testCases {
		t.Run("Test "+tc.desc, func(t *testing.T) {
			out := Justify(tc.words, tc.maxWidth)

			t.Logf("Justify(%d): '%v'\nOutput:\n|%v|",
				tc.maxWidth, strings.Join(tc.words, " "), strings.Join(out, "|\n|"))

			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("Resultado incorreto:\nesperava: %#v,\n  obtive: %#v", tc.out, out)
			}
		})
	}
}
