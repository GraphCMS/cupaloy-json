package examples_test

import (
	"testing"

	"github.com/GraphCMS/cupaloy-json/v2"
	mock "github.com/stretchr/testify/mock"
)

// Snapshots are isolated by package so test functions with the same name are fine
func TestStringJSON(t *testing.T) {
	result := []byte(`{"in":{"object":{"data":{"object":{"assets":{"from":{"collection":{"collection":{"paginate":{"intersection":[{"collection":{"match":"asset_mimeType"},"filter":{"expr":{"equals":[{"find":"p","findstr":{"to_string":{"var":"field"}}},0]},"lambda":["field","ref"]}},{"union":[{"union":{"collection":["xxxxx","zzzzzz"],"map":{"expr":{"from":{"var":"val"},"range":{"match":"asset_mimeType"},"to":{"var":"val"}},"lambda":"val"}}},{"collection":{"match":"asset_mimeType"},"filter":{"expr":{"not":{"equals":[{"find":"zero","findstr":{"to_string":{"var":"field"}}},-1]}},"lambda":["field","ref"]}}]}]},"size":999999999},"map":{"expr":{"get":{"var":"ref"}},"lambda":["field","ref"]}},"map":{"expr":{"in":{"object":{"id":{"from":{"var":"Query_assets"},"select":["data","id"]},"mimeType":{"default":null,"from":{"var":"Query_assets"},"select":["data","mimeType"]}}},"let":[{"Query_assets":{"var":"X"}}]},"lambda":"X"}},"select":["data"]}}}}},"let":[{"context":{"object":{}}}]}`)
	err := cupaloy.SnapshotJSON(result)
	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}

func TestStringJSONPretty(t *testing.T) {
	result := []byte(`{
    "in": {
        "object": {
            "data": {
                "object": {
                    "assets": {
                        "from": {
                            "collection": {
                                "collection": {
                                    "paginate": {
                                        "intersection": [{
                                            "collection": {
                                                "match": "asset_mimeType"
                                            },
                                            "filter": {
                                                "expr": {
                                                    "equals": [{
                                                        "find": "p",
                                                        "findstr": {
                                                            "to_string": {
                                                                "var": "field"
                                                            }
                                                        }
                                                    }, 0]
                                                },
                                                "lambda": ["field", "ref"]
                                            }
                                        }, {
                                            "union": [{
                                                "union": {
                                                    "collection": ["xxxxx", "zzzzzz"],
                                                    "map": {
                                                        "expr": {
                                                            "from": {
                                                                "var": "val"
                                                            },
                                                            "range": {
                                                                "match": "asset_mimeType"
                                                            },
                                                            "to": {
                                                                "var": "val"
                                                            }
                                                        },
                                                        "lambda": "val"
                                                    }
                                                }
                                            }, {
                                                "collection": {
                                                    "match": "asset_mimeType"
                                                },
                                                "filter": {
                                                    "expr": {
                                                        "not": {
                                                            "equals": [{
                                                                "find": "zero",
                                                                "findstr": {
                                                                    "to_string": {
                                                                        "var": "field"
                                                                    }
                                                                }
                                                            }, -1]
                                                        }
                                                    },
                                                    "lambda": ["field", "ref"]
                                                }
                                            }]
                                        }]
                                    },
                                    "size": 999999999
                                },
                                "map": {
                                    "expr": {
                                        "get": {
                                            "var": "ref"
                                        }
                                    },
                                    "lambda": ["field", "ref"]
                                }
                            },
                            "map": {
                                "expr": {
                                    "in": {
                                        "object": {
                                            "id": {
                                                "from": {
                                                    "var": "Query_assets"
                                                },
                                                "select": ["data", "id"]
                                            },
                                            "mimeType": {
                                                "default": null,
                                                "from": {
                                                    "var": "Query_assets"
                                                },
                                                "select": ["data", "mimeType"]
                                            }
                                        }
                                    },
                                    "let": [{
                                        "Query_assets": {
                                            "var": "X"
                                        }
                                    }]
                                },
                                "lambda": "X"
                            }
                        },
                        "select": ["data"]
                    }
                }
            }
        }
    },
    "let": [{
        "context": {
            "object": {}
        }
    }]
}`)
	err := cupaloy.SnapshotJSON(result)
	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}
func TestStringJSONPrettyFail(t *testing.T) {
	result := []byte(`{
    "in": {
        "object": {
            "data": {
                "object": {
                    "assets": {
                        "from": {
                            "collection": {
                                "collection": {
                                    "paginate": {
                                        "intersection": [{
                                            "collection": {
                                                "match": "asset_mimeType"
                                            },
                                            "filter": {
                                                "expr": {
                                                    "equals": [{
                                                        "find": "p",
                                                        "findstr": {
                                                            "to_string": {
                                                                "var": "field"
                                                            }
                                                        }
                                                    }, 0]
                                                },
                                                "lambda": ["field", "ref"]
                                            }
                                        }, {
                                            "union": [{
                                                "union": {
                                                    "collection": ["xxxxx", "zzzzzz"],
                                                    "map": {
                                                        "expr": {
                                                            "from": {
                                                                "var": "val"
                                                            },
                                                            "range": {
                                                                "match": "asset_mimeType"
                                                            },
                                                            "to": {
                                                                "var": "val"
                                                            }
                                                        },
                                                        "lambda": "val"
                                                    }
                                                }
                                            }, {
                                                "collection": {
                                                    "match": "asset_mimeType"
                                                },
                                                "filter": {
                                                    "expr": {
                                                        "not": {
                                                            "equals": [{
                                                                "find": "zero",
                                                                "findstr": {
                                                                    "to_string": {
                                                                        "var": "field"
                                                                    }
                                                                }
                                                            }, -1]
                                                        }
                                                    },
                                                    "lambda": ["field", "ref"]
                                                }
                                            }]
                                        }]
                                    },
                                    "size": 999999999
                                },
                                "map": {
                                    "expr": {
                                        "get": {
                                            "var": "ref"
                                        }
                                    },
                                    "lambda": ["field", "ref"]
                                }
                            },
                            "map": {
                                "expr": {
                                    "in": {
                                        "object": {
                                            "id": {
                                                "from": {
                                                    "var": "Query_assets"
                                                },
                                                "select": ["data", "id"]
                                            },
                                            "mimeType": {
                                                "default": null,
                                                "from": {
                                                    "var": "Query_assets"
                                                },
                                                "select": ["data", "mimeType"]
                                            }
                                        }
                                    },
                                    "let": [{
                                        "Query_assets": {
                                            "var": "X"
                                        }
                                    }]
                                },
                                "lambda": "X"
                            }
                        },
                        "select": ["data"]
                    }
                }
            }
        }
    },
    "let": [{
        "context": {
            "object": {}
        }
    }]
}`)

	mockT := &TestingT{}
	mockT.On("Helper").Return()
	mockT.On("Failed").Return(false)
	mockT.On("Name").Return(t.Name())
	mockT.On("Error", mock.Anything).Return()
	mockT.On("Fatal", mock.Anything).Return()

	cupaloy.SnapshotJSONT(mockT, result)
	mockT.AssertCalled(t, "Error", mock.Anything)

}
