# oneCode_many_bins_fyne
(if you will add some platform like ios I will confirm pull request.............)
The first of all I recomend https://developer.fyne.io/started/, 
options commands for build(you will need C compiler, I recomend m2sys on windows or other gcc on mac and linux) :
1. go build -ldflags "-w -s -H=windowsgui"
2.  upx --ultra-brute
3. (for apk must have ndk for example in linux set) export ANDROID_HOME=$HOME/Android/Sdk; export PATH=$PATH:$ANDROID_HOME/tools
4. fyne package -os android -icon some_icon_name -name some_name -release -appID some_package_name
5. on termux you can follow: https://github.com/termux-ndk-lab/termux-setup-ndk
