# bom-studies-go

![Alt text](https://pixabay.com/get/g71268da09876695bccfb19aadbe85ca775a2bde60858db94a2fa7dee9054e2b5ff30067d427545bd3f754f0efa34f66d12c4df3267148de7c868499058a9ab38823efbd774eae69db4b3eb422af36738_1280.png "Title")

go mod init bom-go-studies                          #Init module path
git init                                            #Initialize git ignore.
git add go.mod                                      #Git add go.mod to git ignore.
go run .	/go run bom-studies-go/hello			#Run go
go mod tidy/ go mod init/ go work init[workspace]   #Install dependencies
go test                                             #Run unit test
go test -fuzz=Fuzz                                  #Run test fuzz
go test -fuzz=Fuzz -fuzztime 30s                    #Run test fuzzing for 30 second
go env -w GOBIN=/Users/bomtawep/go/bin              #Set default value for an environment variable
go install bom-studies-go/gowiki                    #CD to src dir befor run and can go to root to run this package
go build                                            #Test that the package compiles with go build