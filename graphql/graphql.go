package graphql

import (
	"sync"
)

const schema = `
schema {
  query: Query
  mutation: Mutation
}
type Query {
	users: [user!]
}
type Mutation {
	AddUser(input: userInput!): Boolean!
}

type user {
	name: String!
	age: Int!
}

input userInput {
	name: String!
	age: Int!
}
`

type QueryResolver struct{}

type UserResolver struct {
	UserName string
	UserAge  int32
}

type AddUserInput struct {
	Name string
	Age  int32
}

var (
	mutex         sync.Mutex
	testUsersData = []*UserResolver{}
)

func init() {
	testUsersData = []*UserResolver{
		&UserResolver{
			UserName: "zhaogaolong",
			UserAge:  18,
		},
	}
}

func (q *QueryResolver) AddUser(args struct{ Input AddUserInput }) bool {
	mutex.Lock()
	defer mutex.Unlock()
	testUsersData = append(testUsersData, &UserResolver{
		UserName: args.Input.Name,
		UserAge:  args.Input.Age,
	})

	return true
}

func (q *QueryResolver) Users() *[]*UserResolver {
	return &testUsersData
}

func (user *UserResolver) Name() string {
	return user.UserName
}

func (user *UserResolver) Age() int32 {
	return user.UserAge
}
