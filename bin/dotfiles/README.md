# Update vim with pat

* Make the container
* Install pathogen
* Install nerdtree

## Make the container

Star the container:

`docker run -v ~/code3/boo:/go/boo --rm --name go -ti golang`

Then in the container get vim:

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

