package edg_test

import (
	"github.com/speedyhoon/cnst/edg"
	"strconv"
	"testing"
)

func TestUnique(t *testing.T) {
	list := map[int][]int{
		3:       {edg.Triangle},
		4:       {edg.Square, edg.Rectangle, edg.Quadrilateral, edg.Diamond},
		5:       {edg.Pentagon},
		6:       {edg.Hexagon, edg.Tristar, edg.Triquetra},
		7:       {edg.Heptagon},
		8:       {edg.Octagon, edg.Tetragram},
		9:       {edg.Nonagon},
		10:      {edg.Decagon, edg.Pentagram},
		11:      {edg.Hendecagon},
		12:      {edg.Dodecagon, edg.Hexagram},
		13:      {edg.Tridecagon},
		14:      {edg.Tetradecagon, edg.Heptagram},
		15:      {edg.Pentadecagon},
		16:      {edg.Hexadecagon, edg.Octagram},
		17:      {edg.Heptadecagon},
		18:      {edg.Octadecagon, edg.Enneagram},
		19:      {edg.Enneadecagon},
		20:      {edg.Icosagon, edg.Decagram},
		21:      {edg.Icosikaihenagon},
		22:      {edg.Icosikaidigon, edg.Hendecagram},
		23:      {edg.Icositrigon},
		24:      {edg.Icositetragon, edg.Dodecagram},
		25:      {edg.Icosikaipentagon},
		26:      {edg.Icosikaihexagon, edg.Tridecagram},
		27:      {edg.Icosikaiheptagon},
		28:      {edg.Icosikaioctagon, edg.Tetradecagram},
		29:      {edg.Icosikaienneagon},
		30:      {edg.Triacontagon, edg.Pentadecagram},
		32:      {edg.Hexadecagram},
		34:      {edg.Heptadecagram},
		36:      {edg.Octadecagram},
		38:      {edg.Enneadecagram},
		40:      {edg.Tetracontagon, edg.Icosagram},
		42:      {edg.Icosikaihenagram},
		44:      {edg.Icosikaidigram},
		46:      {edg.Icositrigram},
		48:      {edg.Icositetragram},
		50:      {edg.Pentacontagon, edg.Icosikaipentagram},
		52:      {edg.Icosikaihexagram},
		54:      {edg.Icosikaiheptagram},
		56:      {edg.Icosikaioctagram},
		58:      {edg.Icosikaienneagram},
		60:      {edg.Hexacontagon, edg.Triacontagram},
		70:      {edg.Heptacontagon},
		80:      {edg.Octacontagon, edg.Tetracontagram},
		90:      {edg.Enneacontagon},
		100:     {edg.Hectogon, edg.Pentacontagram},
		120:     {edg.Hexacontagram},
		140:     {edg.Heptacontagram},
		160:     {edg.Octacontagram},
		180:     {edg.Enneacontagram},
		200:     {edg.Dihectogon, edg.Hectogram},
		300:     {edg.Trihectogon},
		400:     {edg.Tetrahectogon, edg.Dihectogram},
		500:     {edg.Pentahectogon},
		600:     {edg.Hexahectogon, edg.Trihectogram},
		700:     {edg.Heptahectogon},
		800:     {edg.Octahectogon, edg.Tetrahectogram},
		900:     {edg.Enneahectogon},
		1000:    {edg.Chiliagon, edg.Pentahectogram},
		1200:    {edg.Hexahectogram},
		1400:    {edg.Heptahectogram},
		1600:    {edg.Octahectogram},
		1800:    {edg.Enneahectogram},
		2000:    {edg.Chiliagram},
		10000:   {edg.Myriagon},
		20000:   {edg.Myriagram},
		1000000: {edg.Megagon},
		2000000: {edg.Megagram},
	}

	t.Run("expected quantity", func(t *testing.T) {
		const expectedQty, expectedValues = 96, 69
		if qty := len(list); qty != expectedValues {
			t.Errorf("have %d map items, expected %d\n", qty, expectedValues)
		}

		var qty int
		for _, names := range list {
			qty += len(names)
		}
		if qty != expectedQty {
			t.Errorf("have %d total edges, expected %d\n", qty, expectedQty)
		}
	})

	for expected, names := range list {
		t.Run(strconv.Itoa(expected), func(t *testing.T) {
			for _, name := range names {
				if expected != name {
					t.Errorf("expected edg type `%d`, got: `%d`", name, expected)
				}
			}
		})
	}
}
