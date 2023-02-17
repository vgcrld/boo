# boo

Some sample `go` code just for me to play with. 

The `bin` folder is a list of commands that can be directly installed with 

`go install`

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

`go install github.com/vgcrld/boo/bin/markdown@HEAD`

or Install the **httpd** command right from this command:

`go install github.com/vgcrld/boo/bin/httpd@HEAD`

What happens is that the go command fetches and compiles
the code on the fly and installs it locally. 

So at the prompt you can now just run:

`markdown`

And the command is available, right in the path ready to go.

**NOTE:** The odd thing that I haven't been able to figure out is why latest 
is never actually the latest. I guess it needs to be tagged? The first 
time it's found it works. But Subsequent pushes to github do not tag it as latest.

So you should be able to use `github.com/vgcrld/boo/bin/markdown@latest` but this 
rarely picks up the changes after the first push. 

One thing that __seems__ to be the case is that go.pkg.dev appears to cache stuff.
and on an install it's passing through here? (Maybe).

But if you point to a specific commit hash, you'll get the latest. Or you can point
to HEAD and get the latest (again, I think).

Either way, it's really cool.

