import gql from 'graphql-tag'

export const get_user = gql`
  query {
    get_profile{
        email,
        username,
        id
    }
  }
`
