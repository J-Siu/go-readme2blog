#!bash

# Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

# Geenerate test files for github.com/J-Siu/go-readme2blog/test/

HEADER_BOTTOM="### Bottom"
HEADER_SPLIT="### Split"
HEADER_TOP="### Top"
SPLIT_MARKER="<!--more-->"
TEST_EXT=".md"
TEST_PREFIX="test"
LICENSE='### License

Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.'

# NUM=$1
# LINE="$2"
# FILE=$3
append() {
	NUM=$1
	LINE="$2"
	FILE=$3

	c=0
	while [ ${c} -lt ${NUM} ]; do
		echo "${LINE}" >>$FILE
		# echo "${LINE}"
		((c++))
	done
}

# NUM=$1
# CHAR=$2
build_line() {
	NUM=$1
	CHAR=$2

	c=0
	LINE=""
	while [ ${c} -lt ${NUM} ]; do
		LINE="${LINE}${CHAR}"
		# echo "${c}<${NUM} ${LINE}"
		((c++))
	done
	echo "$LINE"
}

# test-09
# License
# Top header
# 10 x 0
# 10 x 1
# split marker
# Bottom header
# 10 x 2
# ...
# 10 x 9
create09() {
	TEST="09"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	echo "${LICENSE}" >${FILE}
	echo ${HEADER_TOP} >>${FILE}
	for i in {0..1}; do
		LINE=$(build_line 10 ${i})
		append 10 ${LINE} ${FILE}
	done
	echo ${HEADER_SPLIT} >>${FILE}
	echo ${SPLIT_MARKER} >>${FILE}
	echo ${HEADER_BOTTOM} >>${FILE}
	for i in {2..9}; do
		LINE=$(build_line 10 ${i})
		append 10 ${LINE} ${FILE}
	done
}

# test-az
# License
# Top header
# 10 x a
# 10 x b
# split marker
# Bottom header
# 10 x c
# ...
# 10 x z
createAZ() {
	TEST="az"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	echo "${LICENSE}" >${FILE}
	echo ${HEADER_TOP} >>${FILE}
	for i in {a..b}; do
		LINE=$(build_line 10 ${i})
		append 10 ${LINE} ${FILE}
	done
	echo ${HEADER_SPLIT} >>${FILE}
	echo ${SPLIT_MARKER} >>${FILE}
	echo ${HEADER_BOTTOM} >>${FILE}
	for i in {c..z}; do
		LINE=$(build_line 10 ${i})
		append 10 ${LINE} ${FILE}
	done
}

# test-marker-first-ln
createMarkerFirstLn() {
	TEST="marker-first-ln"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	echo ${SPLIT_MARKER} >${FILE}
	for i in {1..6}; do
		LINE=$(build_line 10 ${i})
		append 1 ${LINE} ${FILE}
	done
}

# test-marker-last-ln
createMarkerLastLn() {
	TEST="marker-last-ln"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	>${FILE}
	for i in {1..6}; do
		LINE=$(build_line 10 ${i})
		append 1 ${LINE} ${FILE}
	done
	echo ${SPLIT_MARKER} >>${FILE}
}

# test-marker-none
createMarkerNone() {
	TEST="marker-none"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	>${FILE}
}

# test-maeker-only
createMarkerOnly() {
	TEST="marker-only"
	FILE=${TEST_PREFIX}-${TEST}${TEST_EXT}
	echo ${SPLIT_MARKER} >${FILE}
}

create09
createAZ
createMarkerFirstLn
createMarkerLastLn
createMarkerNone
createMarkerOnly
