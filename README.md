# groq-cli-chat
A CLI Wrapper around GROQ AI Chat using an API key

## Pre-requisites

You will need an Api key from GROQ Cloud to use this CLI. You can get one by signing up at [Groq's website](https://groq.com/) and navigating to the groqcloud section and generating an api key.

The API key should be exported as an environment variable called `GROQ_API_KEY`. I'd recommend setting this up inside your rc file.

## Installation

You can download the binary from the releases section of this repo or build it yourself.

Building from cloning the repo requires golang installed and then running `go build -o gchat` inside the directory.


## Usage


You can then run the binary with `./gchat -model <model> -message <message>` and receive a response from Groq AI. Doesn't support chat history. Just simple string prompts.

Available models in can be found using `-help` flag.

I'd also recommend setting up an alias for the binary in your rc file for ease of use.

`alias g="~/path/to/binary/gchat -model <model> -message"`
