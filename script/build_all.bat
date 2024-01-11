IF NOT EXIST build (
    mkdir build
)

cd build
for /r ../src/service /d %%i in (*) do (
    go build %%i
)