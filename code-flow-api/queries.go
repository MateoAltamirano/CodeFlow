package main

import "fmt"

const allProgramsQuery string = `
{
	programs(func: has(flow)) {
		uid
		name
	}
}
`
const programByIdQuery string = `
{
	program(func: uid(%s)) {
	  uid
	  name
	  code
	  flow
	}
}
`

func getProgramByIdQuery(id string) string {
	return fmt.Sprintf(programByIdQuery, id)
}