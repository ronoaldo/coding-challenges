package anagramas

import (
	"reflect"
	"testing"
)

func MapearAnagramas(nums1, nums2 []int) (out []int) {
	for _, n1 := range nums1 {
		for pos, n2 := range nums2 {
			if n1 == n2 {
				out = append(out, pos)
				break
			}
		}
	}
	return
}

func TestMapeamentoAnagramas(t *testing.T) {
	testCases := []struct {
		desc  string
		nums1 []int
		nums2 []int
		out   []int
	}{
		{
			"modelo do desafio",
			[]int{21, 82, 64, 23, 5},
			[]int{5, 21, 23, 64, 82},
			[]int{1, 4, 3, 2, 0},
		},
		{
			"teste adicional, com valores diversos",
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]int{2, 1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run("Test "+tc.desc, func(t *testing.T) {
			out := MapearAnagramas(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("Mapeamento inválido: experado: %v, obtive: %v", tc.out, out)
			}
		})
	}
}
