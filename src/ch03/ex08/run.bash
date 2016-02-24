rm *.png
go build -o newton

./newton -type=complex64 -output=newton64_100.png
./newton -type=complex64 -zoom=200 -output=newton64_200.png
./newton -type=complex64 -zoom=300 -output=newton64_300.png
./newton -type=complex64 -zoom=400 -output=newton64_400.png

./newton -type=complex128 -output=newton128_100.png
./newton -type=complex128 -zoom=200 -output=newton128_200.png
./newton -type=complex128 -zoom=300 -output=newton128_300.png
./newton -type=complex128 -zoom=400 -output=newton128_400.png

./newton -type=Float -output=newtonFloat_100.png
./newton -type=Float -zoom=200 -output=newtonFloat_200.png
./newton -type=Float -zoom=300 -output=newtonFloat_300.png
./newton -type=Float -zoom=400 -output=newtonFloat_400.png

./newton -type=Float -precision=256 -output=newtonFloat256_100.png
./newton -type=Float -precision=256 -zoom=200 -output=newtonFloat256_200.png
./newton -type=Float -precision=256 -zoom=300 -output=newtonFloat256_300.png
./newton -type=Float -precision=256 -zoom=400 -output=newtonFloat256_400.png

./newton -type=Rat -output=newtonRat_100.png
./newton -type=Rat -zoom=200 -output=newtonRat_200.png
./newton -type=Rat -zoom=300 -output=newtonRat_300.png
./newton -type=Rat -zoom=400 -output=newtonRat_400.png
