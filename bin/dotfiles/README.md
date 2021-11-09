# Make My Default Dotfiles

This is just an excersise in Go. It's not really that 
useful. 

In go it shows how to:

* Use `embed`
* Exec a command
* Use variadic calls
* Use os to get ENV
* Use filepath.Join 
* and other stuff

The process that is done after making the container 
(which has to be manual) is:

* Update apt-get 
* Install vim
* Install pathogen
* Install nerdtree
* Add a default .vimrc to home from 

## Make the container

Start the container:

```bash 
docker run -v ~/code3/boo:/go/boo --rm --name go -ti golang

# In the container
cd /go/boo/bin/dotfiles
go run dotfiles.go
```


# For Reference (The rest is done by dotfiles.go)

This stuff has been added to dotfiles.go but has been
added here for reference.

Then in the container get vim:

## Update `apt-get` and install `vim`

`apt-get update && apt-get install -y vim`

## Install pathogen

```bash
mkdir -p ~/.vim/autoload ~/.vim/bundle && \
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim
```

## Make Minimal .vimrc

```bash
cat << EOF > ~/.vimrc
execute pathogen#infect()
syntax on
filetype plugin indent on
EOF
```

## Install Nerdtree

Clone nerdtree into the proper bundle location:

`git clone https://github.com/preservim/nerdtree.git ~/.vim/bundle/nerdtree`

Then add this to ~/.vimrc to setup the ctl-t toggle:

```bash
cat <<EOF >> ~/.vimrc
map <C-t> :NERDTreeToggle<CR>
map <C-n> :NERDTreeFocus<CR>
EOF
```

