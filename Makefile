build:	
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/pixelscam \
			./
	mv bin/pixelscam-linux-amd64 bin/pixelscam