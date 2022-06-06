<h1 align="center">GithubTools</h1>
<p align="center">
<font size="4px">
Minimalistic toolkit built around the <a href="https://docs.github.com/en/rest">Github API</a> for repetitive tasks
</font>
</p>
<p align="center">

<a href="http://golang.org">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go">
</a>
<img src="https://img.shields.io/github/go-mod/go-version/Gebes/GithubTools.svg" alt="Go Version">
<a href="https://pkg.go.dev/github.com/Gebes/GithubTools">
    <img src="https://img.shields.io/badge/godoc-reference-blue.svg" alt="GoDoc">
</a>
<a href="https://goreportcard.com/report/github.com/Gebes/GithubTools">
    <img src="https://goreportcard.com/badge/github.com/Gebes/GithubTools" alt="GoReportCard">
</a>
<a href="https://github.com/Gebes/GithubTools/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/Gebes/GithubTools.svg" alt="License">
</a>
<a href="https://GitHub.com/Gebes/GithubTools/releases/">
    <img src="https://img.shields.io/github/release/Gebes/GithubTools" alt="Latest releace">
</a>
<a href="https://gocover.io/github.com/Gebes/GithubTools">
    <img src="https://gocover.io/_badge/github.com/Gebes/GithubTools" alt="Latest releace">
</a>
<a href="https://www.codefactor.io/repository/github/Gebes/GithubTools">
    <img src="https://www.codefactor.io/repository/github/Gebes/GithubTools/badge" alt="Latest releace">
</a>

<br><br>
<a href="https://gitHub.com/Gebes/GithubTools/graphs/commit-activity">
<img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" alt="Maintained?">
</a>
<a href="https://github.com/Gebes">
<img src="https://img.shields.io/badge/Maintainer-Gebes-blue" alt="Maintainer">
</a>
</p>



## 🔥 Get Started

### 👩‍🔬 Clone the repository
1. Ensure you have the git client installed 
2. Run `git clone github.com/Gebes/GithubTools`

### 🔨 Configure the project
1. Create a .env file 
2. Generate a <abbr title="Personal Access Token">PAT</abbr> in your [settings](https://github.com/settings/tokens)
3. Randomly generate a `Key` for [authorization](#authorization) (you can leave it empty, but this is not recommended ⚠️)
4. Set all environment variables
```dotenv
GITHUB_ACCESS_TOKEN=ghp_token
KEY=key
```

### 🛫 Run the project
1. Run `go mod tidy`
2. To launch the backend run `go run ./cmd/backend`
3. Access the project under `localhost:8080`

### ℹ️ Make a request
1. Call the specific route
   * Make sure to include the `Key` in the `Authorization` header for propper [authorization](#authorization)


### 🐳 Deploy the project
This project provides a sample Dockerfile especially made for Linux. You can reuse it, but you may need to adapt it if you are on another operating system.

## 📱 Features

### Authorization
Protect all the routes with the `KEY` from the `.env` file. The router will make sure only process requests that contain a valid `KEY` in the `Authorization` header.

### /no-follow-back/:user
Fetches the first 50 followings of the `:user` and returns all people that do not follow back in a list.

### /contains-file/:owner/:repo?files=files
Checks if the provided repository `:owner/:repo` contains or contained any provided files separated by a comma.

