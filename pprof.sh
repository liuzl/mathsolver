>1000.txt;for ((i=0;i<10000;i++)) do echo "八十的平方与八千八百八十八的和">>1000.txt; done
go build
./mathsolver -i 1000.txt -cpuprofile profile.prof > /dev/null
go tool pprof ./mathsolver profile.prof
