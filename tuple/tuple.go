package tuple

type Pair[A, B any] struct {
	First  A
	Second B
}

type Triple[A, B, C any] struct {
	First  A
	Second B
	Third  C
}

type Quadruple[A, B, C, D any] struct {
	First  A
	Second B
	Third  C
	Fourth D
}

type Quintuple[A, B, C, D, E any] struct {
	First  A
	Second B
	Third  C
	Fourth D
	Fifth  E
}

type Sextuple[A, B, C, D, E, F any] struct {
	First  A
	Second B
	Third  C
	Fourth D
	Fifth  E
	Sixth  F
}

type Septuple[A, B, C, D, E, F, G any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
}

type Octuple[A, B, C, D, E, F, G, H any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
	Eighth  H
}

type Nonuple[A, B, C, D, E, F, G, H, I any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
	Eighth  H
	Ninth   I
}

type Decuple[A, B, C, D, E, F, G, H, I, J any] struct {
	First   A
	Second  B
	Third   C
	Fourth  D
	Fifth   E
	Sixth   F
	Seventh G
	Eighth  H
	Ninth   I
	Tenth   J
}
