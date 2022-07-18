.PHONY: clone-trie
clone-trie:
	git clone -b empty-remove-bug git@github.com:derekparker/trie.git

.PHONY: test
test:
	GO111MODULE=off go test -v ./trie

.PHONY: debug
debug:
	cd ./trie && dlv debug

.PHONY: dlv-test
dlv-test:
	cd ./trie && dlv test