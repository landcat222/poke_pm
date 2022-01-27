package pokemana

type Pokemon struct {
    ID UUID
    Specie string
    Lv int
    Gender string
    Ability string
    Nature string
    HeldItem string
    Moves []string
    IVs []int
    EVs []int
}
