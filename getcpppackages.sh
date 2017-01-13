#!/bin/bash
git pull origin master
function cloneAndCheckout {
	cd ..
	if [ -d "$1" ]; then
		cd $1
		git pull origin master
	else
		git clone https://github.com/phillyfan1138/$1
		cd $1
	fi

	git checkout $2
	cd ..
	cd go-apicpp
}

function compile {
	cd $1
	make 
	cp $1 ../../../../bin
	cd ..

}

cloneAndCheckout marketRisk a12d19f2c8a39f23bbdfa5d300fb8e93c36cae56
cloneAndCheckout creditRisk fb0782e82d4a68cc8b36d15471958f284d1ef7a0
cloneAndCheckout opsRisk ada8f6434e3666e5e4d5773386fbc7930d42e794
cloneAndCheckout HullWhite 5cee8dec3bf60c93e2099dc4d0add82283231860
cloneAndCheckout GaussNewton 82fcc7407469aecba142286e28d4b41b0669e707
cloneAndCheckout FunctionalUtilities 2543852b038436c141a1e3c29bf9c353df9a02b9
cloneAndCheckout BinomialTree d40d2ebbc9de0e1b7d30c3ec82d55c6d4346ef9a
cloneAndCheckout MonteCarlo 3d8d757c340bca12a064964eee33d60acd888607
cloneAndCheckout CharacteristicFunctions 369e1d54b80a2324bd0e2553adf195f01699bb9e
cloneAndCheckout FangOost d74b4e541aa5569ccd99192fe3d256e7d23b3883
cloneAndCheckout RungeKutta b286ce51f5e6957a59a5a49a3d1c60abdd765af4
cloneAndCheckout AutoDiff d9d79e251f4ab030bcf17fc7b756b4ba06e6fdbd
cd ..
git clone https://github.com/miloyip/rapidjson
compile marketRisk
compile opsRisk
compile creditRisk