SEED = $(shell date +%s)
TEST_SIZE = 320

.PHONY: build
build: 
	cd single; go build -o="single"
	mv single/single build/single
	mv out.png last.png
	./build/single -size=$(TEST_SIZE) -longName=false -out=out

.PHONY: buildD
buildD:
	cd single; go build -o="single"
	mv single/single build/single
	mv out.png last.png
	./build/single -size=250 -longName=false -out=out -debug=true

.PHONY: buildR
buildR:
	cd random; go build -o="random"
	mv random/random build/random
	./build/random -n=1

.PHONY: buildAll
buildAll:
	cd single; go build -o="single"
	mv single/single build/single
	cd random; go build -o="random"
	mv random/random build/random
	cd enhance; go build -o="enhance"
	mv enhance/enhance build/enhance
	cd enhance-convert; go build -o="enhance-convert"
	mv enhance-convert/enhance-convert build/enhance-convert

.PHONY: keep
keep:
	cp out.png keep.png

SETNAME = Nets
SETDIR = $(SETNAME)-$(shell date +%s)
SIZE = 2400
SET_A = -c2r=96 -c2g=131 -c2b=186
SET_B = -c2r=255 -c2g=0 -c2b=0
SET_C = -c2r=0 -c2g=34 -c2b=0
SET_D = -c2r=242 -c2g=170 -c2b=255

.PHONY: set
set:
	mkdir $(SETDIR)
	cd $(SETDIR); ../single -size=$(SIZE) -longName=true -out=$(SETNAME) $(SET_A) -seed=$(SEED)
	cd $(SETDIR); ../single -size=$(SIZE) -longName=true -out=$(SETNAME) $(SET_B) -seed=$(SEED)
	cd $(SETDIR); ../single -size=$(SIZE) -longName=true -out=$(SETNAME) $(SET_C) -seed=$(SEED)
	cd $(SETDIR); ../single -size=$(SIZE) -longName=true -out=$(SETNAME) $(SET_D) -seed=$(SEED)

COUNT = 150
BATCHDIR = rand-$(COUNT)-$(SEED)

.PHONY: random
random:
	mkdir $(BATCHDIR)
	cd $(BATCHDIR); ../build/random -n=$(COUNT) -seed=$(SEED)
