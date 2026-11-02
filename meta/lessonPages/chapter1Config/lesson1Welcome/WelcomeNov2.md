# Welcome to the Blog Aggregator

We're going to build an [RSS](https://en.wikipedia.org/wiki/RSS) feed aggregator in Go! We'll call it "Gator", you know, because aggreGATOR 🐊. Anyhow, it's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!


## for Windows Users
If you try to complete this course without WSL 2 installed on Windows, you're gonna have a bad time.

## Prerequisites
The project assumes that you're already familiar with the [Go programming language](https://www.boot.dev/courses/learn-golang) and [SQL databases](https://www.boot.dev/courses/learn-sql).

## Learning Goals
- Learn how to integrate a Go application with a PostgreSQL database
- Practice using your SQL skills to query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go)
- Learn how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

## Assignment

Before we start, let's get you setup. You'll need both the go toolchain (version 1.27+) and the Boot.dev CLI installed. If you don't already have them, here are the [installation instructions](https://github.com/bootdotdev/bootdev).

Some of the CLI tests in this course depend on tools and services installed on your local machine, like Go and PostgreSQL.

**Submit** the CLI tests.

Copy-paste this command into your terminal:

Submit:
bootdev run -s 14b7179b-ced3-4141-9fa5-e67dbc3e5242

This submission is safe: failed checks won't cost armor or Sharpshooter progress.

1. **Run your first CLI test**
echo "bootdev cli is ready!"
- Expecting exit code: `0`
- Expecting stdout to contain all of:
	- `bootdev cli is ready!`

2. **Check Go is installed**
go version
- Expecting exit code: `0`

**Using the Bootdev CLI**

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.
