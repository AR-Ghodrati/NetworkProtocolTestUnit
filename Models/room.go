package Models

type room struct {
	uuid        string
	name        string
	activeUsers int8
	maxUser     int8
}
