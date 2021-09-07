# oneCode_many_bins_fyne
options commands for build(you will need C compiler, I recomend m2sys on windows or other gcc on mac and linux) :
1. go build -ldflags "-w -s -H=windowsgui"
2.  upx --ultra-brute
3. fyne package -os os -icon icon -name ....
4.  fyne package -os os -icon icon -name .... -release -appID c....
5. (for apk must have ndk for example in linux set:) export ANDROID_HOME=$HOME/Android/Sdk; export PATH=$PATH:$ANDROID_HOME/tools
