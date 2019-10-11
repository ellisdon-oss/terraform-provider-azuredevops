// Username and Password 
resource "azuredevops_service_endpoint" "bitbucket" {
  project_id = "<project id>"

  name = "<endpoint name>"

  owner = "Library"
  type = "bitbucket"

  // only accepted value is https://api.bitbucket.org
  url = "<bitbucket url>"

  authorization {
    scheme = "UsernamePassword"
    

    parameters = {
      username = "<username>"
      password = "<password>"
    }
  }

  data = {
    displayName = "<display name>"
    avatarUrl = "<avatar url>"
  }
}

// OAuth2 - Not usable due to auth flow
// resource "azuredevops_service_endpoint" "bitbucket_oauth2" {
//   project_id = "<project id>"
// 
//   name = "<endpoint name>"
// 
//   owner = "Library"
//   type = "bitbucket"
// 
//   url = "<bitbucket url>"
// 
//   authorization {
//     scheme = "OAuth"
//     
//     parameters = {
//       AccessToken = "<guid>"
//       ConfigurationId = "<guid>" //Need more docs
//     }
//   }
// 
//   data = {
//     displayName = "<display name>"
//     avatarUrl = "<avatar url>"
//   }
// }
