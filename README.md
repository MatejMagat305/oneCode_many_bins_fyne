# oneCode_many_bins_fyne
options commands for build :
1. go build -ldflags "-w -s -H=windowsgui"
2.  upx --ultra-brute
3. fyne package -os windows -icon .\coding.png -name fyneCommand
4.  fyne package -os windows -icon .\coding.png -name fyneCommand -release -appID com.my.co
