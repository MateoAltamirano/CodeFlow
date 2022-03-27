package main

import "fmt"

const allProgramsQuery string = `
{
	programs(func: has(code)) {
		uid
	}
}
`
const programByIdQuery string = `
{
	program(func: uid(%s)) {
	  uid
	  code
	}
}
`

func getProgramByIdQuery(id string) string {
	return fmt.Sprintf(programByIdQuery, id)
}