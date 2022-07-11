trie:
	git clone -b empty-remove-bug git@github.com:derekparker/trie.git

.PHONY: trie
test:
	cd trie && go test -v .

.PHONY: debug
debug:
	cd ./trie && dlv debug

.PHONY: dlv-test
dlv-test:
	cd ./trie && dlv test
