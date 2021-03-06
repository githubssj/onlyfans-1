# Go Onlyfans

A go CLI to scrape and archive content from onlyfans.

## Local Install
Assuming you have a proper [go installation](https://golang.org/doc/install) there are few options to running the cli.

You can grab the cli by running:

`go get github.com/januairi/go-of`

or

`git clone https://github.com/januairi/go-of.git`

You can install the binary by running `go install` from the root directory of the project. If you are running the binary you can test if it was installed properly by typing `go-of` into your command line. All commands will need to have a prefix of `go-of` to run them if you are running the commands via the binary.

You can also run `go run main.go download photo username1234` when you are in the projects root directory. 

A list of commands can be found [here](https://github.com/januairi/go-of/blob/main/README.md#commands)

go-of makes use of a config file in the users $HOME directory that can be YAML or JSON format. The file can be called ~/.go-of.json or ~/.go-of.yaml

You can also set environment variables or pass the key-value pairs on the command line.

These are templates that can be used for your config file:

YAML:
```
token: app-token
session: access-token or sess
user_agent: user-agent
auth_id: auth-id
save_dir: abs/path/to/where/content/should/be/saved
```

JSON:
```
{
   "token": "app-token",
   "session": "access-token or sess",
   "user_agent": "user-agent",
   "auth_id": "auth-id",
   "save_dir": "abs/path/to/where/content/should/be/saved"
}
```

These values can be found by logging into onlyfans and inspecting the network api calls that have query params with the value ?app-token=some-value.

In chrome you can right click on the web page, click inspect, and navigate to the network tab.

You can filter the api calls with `?app-token=`.
![network](examples/network.png "config help")



Once you find an appropriate api call, clicking on the headers should look like the following:
![example](examples/example.png "config help")

## Commands
The commands use common verbs and phrases to be called, all start with a prefix of `go-of`

example: `go-of download photo onlyfansuser1234`

| Command      | Args | Output    |
| :---        |    :----:   |          ---: |
| `go-of download content`       | username       | Downloads all media content from a user to the directory provided in config   |
| `go-of download photo`       | username       | Downloads photos to the directory provided in config   |
| `go-of download video`   | username        | Downloads videos to the directory provided in config      |
| `go-of download post`   | username        | Downloads post media to the directory provided in config      |
| `go-of download archived post`   | username        | Downloads post media to the directory provided in config      |
| `go-of download message`   | username        | Downloads message media to the directory provided in config      |
| `go-of download highlight`   | username        | Downloads story highlights media to the directory provided in config      |
| `go-of get user`   | username        | Spits out user info in json format    |


## FAQ

No you can't bypass paywalls with this cli.

No you can't download content you haven't paid for.

No this isn't an onlyfans hack.

No information used by this program is shared.

If nothing is being pulled when running commands, make sure you session/access-token has not changed or expired. If it has you'll need to use that new token isntead.
