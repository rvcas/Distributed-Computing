path_append ()  { path_remove $1; export PATH="$PATH:$1"; }
path_prepend () { path_remove $1; export PATH="$1:$PATH"; }
path_remove ()  { export PATH=`echo -n $PATH | awk -v RS=: -v ORS=: '$0 != "'$1'"' | sed 's/:$//'`; }

newgopath=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
oldgopath=$GOPATH

gsbashprofile=~/.bash_profile

path_remove $oldgopath/bin

grep -v "source $GOPATH/goswitch.sh" $gsbashprofile > bashptemp
mv bashptemp $gsbashprofile

export GOPATH=$newgopath
path_append $GOPATH/bin
export GOROOT=/usr/local/opt/go/libexec

echo "  GoSwitch: Active Go Project ($(tput setaf 6)$GOPATH$(tput sgr 0)) "
