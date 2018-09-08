>1000.txt;for ((i=0;i<1000;i++)) do echo "十万八千与八千八百八十八的和">>1000.txt; done
./mathsolver -i 1000.txt -cpuprofile profile.prof > /dev/null
go tool pprof ./mathsolver profile.prof
