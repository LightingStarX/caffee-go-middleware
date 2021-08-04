package stringx

const defaultMask = '*'

type (
	// TrieFunc 提供了定制trie树的函数
	TrieFunc func(trie *trieNode)

	Trie interface {
		Filter(text string) (string, []string, bool)
		FindKeyWords(text string) []string
	}

	trieNode struct {
		node
		mask rune
	}

	scope struct {
		start int
		end   int
	}

	node struct {
		children map[rune]*node
		end      bool
	}
)

// parseWordToTree 构建一个单词的树形关系，构建字典树
func (nd *node) parseWordToTree(word string) {
	chars := []rune(word)
	if len(chars) == 0 {
		return
	}

	ndcp := nd
	for _, char := range chars {
		if ndcp.children == nil {
			child := new(node)
			ndcp.children = map[rune]*node{
				char: child,
			}
			ndcp = child
		} else if child, ok := ndcp.children[char]; ok {
			ndcp = child
		} else {
			child := new(node)
			ndcp.children[char] = child
			ndcp = child
		}
	}
	ndcp.end = true
}

func NewTrie(words []string, opts ...TrieFunc) *trieNode {
	n := new(trieNode)

	for _, opt := range opts {
		opt(n)
	}

	if n.mask == 0 {
		n.mask = defaultMask
	}

	for _, word := range words {
		n.parseWordToTree(word)
	}
	return n
}
