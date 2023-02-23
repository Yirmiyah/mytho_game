package structure

type Team1 struct {
	Name   string
	Avatar string
	Score  int
}

type Team2 struct {
	Name   string
	Avatar string
	Jeton  int
}

type Quizz struct {
	Question []string
	Response []string
	Level    []string
}

type FakeResponse struct {
	Response1 string
	Response2 string
}

type Manche struct {
	Tour int
}
