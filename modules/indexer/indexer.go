// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package indexer

type SearchModeType string

const (
	SearchModeSemantic SearchModeType = "semantic" // AI-powered semantic search (default when enabled)
	SearchModeExact    SearchModeType = "exact"
	SearchModeWords    SearchModeType = "words"
	SearchModeFuzzy    SearchModeType = "fuzzy"
	SearchModeRegexp   SearchModeType = "regexp"
)

type SearchMode struct {
	ModeValue    SearchModeType
	TooltipTrKey string
	TitleTrKey   string
}

func SearchModesExactWords() []SearchMode {
	return []SearchMode{
		{
			ModeValue:    SearchModeExact,
			TooltipTrKey: "search.exact_tooltip",
			TitleTrKey:   "search.exact",
		},
		{
			ModeValue:    SearchModeWords,
			TooltipTrKey: "search.words_tooltip",
			TitleTrKey:   "search.words",
		},
	}
}

func SearchModesExactWordsFuzzy() []SearchMode {
	return append(SearchModesExactWords(), []SearchMode{
		{
			ModeValue:    SearchModeFuzzy,
			TooltipTrKey: "search.fuzzy_tooltip",
			TitleTrKey:   "search.fuzzy",
		},
	}...)
}

func GitGrepSupportedSearchModes() []SearchMode {
	return append(SearchModesExactWords(), []SearchMode{
		{
			ModeValue:    SearchModeRegexp,
			TooltipTrKey: "search.regexp_tooltip",
			TitleTrKey:   "search.regexp",
		},
	}...)
}

// SemanticSearchMode returns the semantic search mode definition
func SemanticSearchMode() SearchMode {
	return SearchMode{
		ModeValue:    SearchModeSemantic,
		TooltipTrKey: "search.semantic_tooltip",
		TitleTrKey:   "search.semantic",
	}
}

// SearchModesWithSemantic prepends semantic mode to existing modes (makes it default)
func SearchModesWithSemantic(modes []SearchMode) []SearchMode {
	return append([]SearchMode{SemanticSearchMode()}, modes...)
}
