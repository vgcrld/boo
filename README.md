# boo

Some sample go code just for me.

# Install commands on a Container

The best part of go is the ability to distribute commands in a utility sense like this.

If you have docker destop running, get a go container up and running:

```bash
# This will be temporary
# it will disapear after you exit
docker run -ti --rm golang
```

In the container, run this go command:

Install the **markdown** command
`go install github.com/vgcrld/boo/bin/markdown@03b088cc`

or Install the **httpd** command right from this command:
`go install github.com/vgcrld/boo/bin/httpd@03b088cc`

It so great. What happens is that the go command fetches and compiles
the code on the fly and installs it locally. 

So at the prompt you can now just run:

`root@11d36b6ee56b:/go# markdown`

And the command is available, right in the path ready to go.

