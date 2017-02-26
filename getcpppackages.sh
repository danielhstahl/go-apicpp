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

cloneAndCheckout marketRisk efb7b70c8f1f9b20d57885e98c72a92b9f1b7152 #a12d19f2c8a39f23bbdfa5d300fb8e93c36cae56
cloneAndCheckout creditRisk e06bb84de62e55ac26bcc95dbeb15daa0b2d7545 #fb0782e82d4a68cc8b36d15471958f284d1ef7a0
cloneAndCheckout opsRisk 46ab4ccb7f64dc21b9684ca9e95ad20e292436f3 #ada8f6434e3666e5e4d5773386fbc7930d42e794
cloneAndCheckout HullWhite 0fc3e44b45021c84a3db5cf28511a2903f84f502 #5cee8dec3bf60c93e2099dc4d0add82283231860
cloneAndCheckout GaussNewton 518df40f50be744d766d0a3da0db6af19a476094 #82fcc7407469aecba142286e28d4b41b0669e707
cloneAndCheckout FunctionalUtilities 2543852b038436c141a1e3c29bf9c353df9a02b9 #not needed
cloneAndCheckout BinomialTree adfe12447efa369b51825fb4258db356422ee229 #d40d2ebbc9de0e1b7d30c3ec82d55c6d4346ef9a
cloneAndCheckout MonteCarlo 3d8d757c340bca12a064964eee33d60acd888607 #not needed
cloneAndCheckout CharacteristicFunctions 89da0e667d01b2a6c04d8dc8932e8be884a35d7b #369e1d54b80a2324bd0e2553adf195f01699bb9e
cloneAndCheckout FangOost d74b4e541aa5569ccd99192fe3d256e7d23b3883 #not needed
cloneAndCheckout RungeKutta d5056950c35c5b3854c2dbf554516927bd2a2e22 #b286ce51f5e6957a59a5a49a3d1c60abdd765af4
cloneAndCheckout AutoDiff 78c4688fa51dd8336cea4a3b8e94ae8a12d36ba7 #d9d79e251f4ab030bcf17fc7b756b4ba06e6fdbd
cd ..
git clone https://github.com/miloyip/rapidjson
compile marketRisk
compile opsRisk
compile creditRisk