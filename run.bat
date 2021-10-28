go clean -testcache -cache
del .\results.txt
del .\WaterSortPuzzleSolver.exe
go build
.\WaterSortPuzzleSolver.exe > results.txt
pause