.PHONY: clone-trie
clone-trie:
	git clone -b empty-remove-bug git@github.com:derekparker/trie.git

.PHONY: test
test:
	go test -v ./trie/trie.go ./trie/trie_test.go

.PHONY: debug
debug:
	cd ./trie && dlv debug

.PHONY: dlv-test
dlv-test:
	cd ./trie && dlv test