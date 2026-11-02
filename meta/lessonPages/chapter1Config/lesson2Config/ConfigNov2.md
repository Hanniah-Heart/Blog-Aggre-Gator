# Config

Gator is a multi-user CLI application. There's no server (other than the database), so it's only intended for local use, but just like games in the 90's and early 2000's, that doesn't mean we can't have multiplayer functionality on a single device!

We'll use a single [JSON](https://www.json.org/json-en.html) file to keep track of two things:

1. Who is currently logged in
2. The connection credentials for the PostgreSQL database

The JSON file should have this structure (when prettified):

```
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```

Note: There's no user-based authentication for this app. If someone has the database credentials, they can act as any user. We'll cover auth in other courses. We want to focus on SQL, CLIs and long-running services for this project.

Throughout this project, we'll suggest file names and locations to keep your code organized. If you're unsure, follow our structure. You can organize files differently, but keep the Go packages and paths used by the tools and CLI tests.

## Assignment

1. Manually create a config file in your home directory, `~/.gatorconfig.json`, with the following content:

```
{
  "db_url": "postgres://example"
}
```
Don't worry about adding `current_user_name`, that will be set by the application.

2. In a new project directory, initialize a new Go module using `go mod init`. Create a new file, `main.go`. Add a `main` function. It doesn't need to do anything yet.

3. Create `internal/config/config.go` with a `package config` declaration. This package will be responsible for reading and writing the JSON file.

This package should have the following functionality exported so the `main` package can use it:

- Export a `Config` struct that represents the JSON file structure, including struct tags.
- Export a `Read` function that reads the JSON file found at `~/.gatorconfig.json` and returns a `Config` struct. It should read the file from the HOME directory, then decode the JSON string into a new `Config` struct. I used [os.UserHomeDir](https://pkg.go.dev/os#UserHomeDir) to get the location of HOME.
- Export a `SetUser` method on the `Config` struct that writes the config struct to the JSON file after setting the `current_user_name` field.

I also wrote a few non-exported helper functions and added a constant to hold the _filename_.

-`getConfigFilePath() (string, error)`
-`write(cfg Config) error`
-`const configFileName = ".gatorconfig.json"`

But you can implement the internals of the package however you like.

4. Update the `main` function to:
- Read the config file.
- Set the current user to "lane" (actually, you should use your name instead) and update the config file on disk.
- Read the config file again and print the contents of the config struct to the terminal.

**Run and submit** the CLI tests from the **root of your project**.

Copy-paste one of these commands into your terminal:
Run:
bootdev run 5e53c561-f059-49aa-b06c-5c5535fa6df0
Submit:
bootdev run -s 5e53c561-f059-49aa-b06c-5c5535fa6df0

The CLI will run these checks:
1. **Print the configured database URL**
go run .
- Expecting exit code: `0`
- Expecting stdout to contain all of:
	- `postgres://example`

**Using the Bootdev CLI**

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

