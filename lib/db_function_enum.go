package lib

type DBFunction int

const (
	FnFindOne DBFunction = iota + 1 //EnumIndex = 1

	FnFindMany //EnumIndex = 2

	FnDeleteOne //EnumIndex = 3

	FnDeleteMany //EnumIndex = 4

	FnInsertOne //EnumIndex = 5

	FnInsertMany //EnumIndex = 6
)

// String - Creating common behavior - give the type a String function
func (f DBFunction) String() string {
	return [...]string{
		"FnFindOne",
		"FnFindMany",
		"FnDeleteOne",
		"FnDeleteMany",
		"FnInsertOne",
		"FnInsertMany"}[f-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (f DBFunction) EnumIndex() int {
	return int(f)
}
