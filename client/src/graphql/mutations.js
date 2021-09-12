import gql from "graphql-tag"

export const login = gql`
    mutation login($username: String! $password: String!){
        login(input : {username: $username, password: $password})
    }
`
export const register = gql`
    mutation createUser($username: String! $email: String! $password: String! ) {
        createUser( input: {username: $username, email: $email, password: $password })
    }
    `