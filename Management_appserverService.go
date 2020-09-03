package marklogic

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	AppServers = "LATEST/servers"
)

type ServerParameters struct {
	Format   string `url:"format,omitempty"` //html, json, or xml
	GroupId  string `url:"group-id,omitempty"`
	View     string `url:"view,omitempty"` //schema, status, metrics, package, or default.
	Fullrefs bool   `url:"fullrefs,omitempty"`
	HostId   string `url:"host-id,omitempty"`
	Modules  bool   `url:"modules,omitempty"`
}

type AppServersDefaultResponse struct {
	ServerDefaultList ServerDefaultList `json:"server-default-list"`
}

type ServerDefaultList struct {
	Meta         Meta         `json:"meta"`
	Relations    Relations    `json:"relations"`
	ListItems    ListItems    `json:"list-items"`
	RelatedViews RelatedViews `json:"related-views"`
}

type AppServersStatusResponse struct {
	ServerStatusList ServerStatusList `json:"server-status-list"`
}

type ServerStatusList struct {
	Meta              Meta      `json:"meta"`
	Relations         Relations `json:"relations"`
	StatusListSummary struct {
		RequestRate struct {
			Units string  `json:"units"`
			Value float64 `json:"value"`
		} `json:"request-rate"`
		ExpandedTreeCacheMissRate struct {
			Units string  `json:"units"`
			Value float64 `json:"value"`
		} `json:"expanded-tree-cache-miss-rate"`
		ExpandedTreeCacheHitRate struct {
			Units string  `json:"units"`
			Value float64 `json:"value"`
		} `json:"expanded-tree-cache-hit-rate"`
		RequestCount struct {
			Units string  `json:"units"`
			Value float64 `json:"value"`
		} `json:"request-count"`
	} `json:"status-list-summary"`
	StatusListItems struct {
		StatusListItem []struct {
			RelationID   []string `json:"relation-id"`
			Groupnameref string   `json:"groupnameref"`
			Uriref       string   `json:"uriref"`
			Kindref      string   `json:"kindref"`
			ContentDb    string   `json:"content-db"`
			Idref        string   `json:"idref"`
			Nameref      string   `json:"nameref"`
			ModulesDb    string   `json:"modules-db,omitempty"`
		} `json:"status-list-item"`
	} `json:"status-list-items"`
	RelatedViews RelatedViews `json:"related-views"`
}

type AppServersMetricResponse struct {
	ServerMetricsList ServerMetricsList `json:"server-metrics-list"`
}

type ServerMetricsList struct {
	Meta struct {
		URI         string    `json:"uri"`
		CurrentTime time.Time `json:"current-time"`
		ElapsedTime string    `json:"elapsed-time"`
	} `json:"meta"`
	MetricProperties struct {
		Start   string `json:"start"`
		Period  string `json:"period"`
		Server  string `json:"server"`
		Summary bool   `json:"summary"`
		Detail  bool   `json:"detail"`
	} `json:"metric-properties"`
	MetricsRelations struct {
		ServerMetricsList struct {
			Uriref  string `json:"uriref"`
			Typeref string `json:"typeref"`
			Metrics []struct {
				DbLibModuleCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-lib-module-cache-hits,omitempty"`
				DbLibModuleCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-lib-module-cache-misses,omitempty"`
				DbMainModuleSeqCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-main-module-seq-cache-hits,omitempty"`
				DbMainModuleSeqCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-main-module-seq-cache-misses,omitempty"`
				DbProgramCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-program-cache-hits,omitempty"`
				DbProgramCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"db-program-cache-misses,omitempty"`
				EnvProgramCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"env-program-cache-hits,omitempty"`
				EnvProgramCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"env-program-cache-misses,omitempty"`
				ExpandedTreeCacheHitRate struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"expanded-tree-cache-hit-rate,omitempty"`
				ExpandedTreeCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"expanded-tree-cache-hits,omitempty"`
				ExpandedTreeCacheMissRate struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"expanded-tree-cache-miss-rate,omitempty"`
				ExpandedTreeCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"expanded-tree-cache-misses,omitempty"`
				FsLibModuleCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-lib-module-cache-hits,omitempty"`
				FsLibModuleCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-lib-module-cache-misses,omitempty"`
				FsMainModuleSeqCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-main-module-seq-cache-hits,omitempty"`
				FsMainModuleSeqCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-main-module-seq-cache-misses,omitempty"`
				FsProgramCacheHits struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-program-cache-hits,omitempty"`
				FsProgramCacheMisses struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"fs-program-cache-misses,omitempty"`
				QueueSize struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"queue-size,omitempty"`
				RequestRate struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"request-rate,omitempty"`
				RequestTime struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"request-time,omitempty"`
				ServerReceiveLoad struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"server-receive-load,omitempty"`
				ServerReceiveRate struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"server-receive-rate,omitempty"`
				ServerSendLoad struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"server-send-load,omitempty"`
				ServerSendRate struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value int       `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"server-send-rate,omitempty"`
				Threads struct {
					Name    string `json:"name"`
					Units   string `json:"units"`
					Desc    string `json:"desc"`
					Summary struct {
						Count int    `json:"count"`
						Agg   string `json:"agg"`
						Data  struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"summary"`
					HTTPServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"http-server"`
					OdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"odbc-server"`
					WebdavServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"webdav-server"`
					XdbcServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []interface{} `json:"entry"`
						} `json:"data"`
					} `json:"xdbc-server"`
					TaskServer struct {
						ServerType string `json:"server-type"`
						Count      int    `json:"count"`
						Data       struct {
							Entry []struct {
								Dt    time.Time `json:"dt"`
								Value float64   `json:"value"`
							} `json:"entry"`
						} `json:"data"`
					} `json:"task-server"`
					Detail interface{} `json:"detail"`
				} `json:"threads,omitempty"`
			} `json:"metrics"`
		} `json:"server-metrics-list"`
	} `json:"metrics-relations"`
	RelatedViews struct {
		RelatedView struct {
			ViewType string `json:"view-type"`
			ViewName string `json:"view-name"`
			ViewURI  string `json:"view-uri"`
		} `json:"related-view"`
	} `json:"related-views"`
}

type AppServersPackageResponse struct {
	Set struct {
		Servers []struct {
			PackageHTTPServer PackageHTTPServer `json:"package-http-server"`
		} `json:"servers"`
	} `json:"set"`
}

type PackageHTTPServer struct {
	Metadata struct {
		PackageVersion string    `json:"package-version"`
		User           string    `json:"user"`
		Group          string    `json:"group"`
		Host           string    `json:"host"`
		Timestamp      time.Time `json:"timestamp"`
		Platform       string    `json:"platform"`
	} `json:"metadata"`
	Config struct {
		ServerType            string `json:"server-type"`
		GroupName             string `json:"group-name"`
		Name                  string `json:"name"`
		PackageHTTPProperties struct {
			Enabled                            bool        `json:"enabled"`
			Root                               string      `json:"root"`
			Authentication                     string      `json:"authentication"`
			Port                               int         `json:"port"`
			WebDAV                             bool        `json:"webDAV"`
			Execute                            bool        `json:"execute"`
			DisplayLastLogin                   bool        `json:"display-last-login"`
			Address                            string      `json:"address"`
			Backlog                            int         `json:"backlog"`
			Threads                            int         `json:"threads"`
			RequestTimeout                     int         `json:"request-timeout"`
			KeepAliveTimeout                   int         `json:"keep-alive-timeout"`
			SessionTimeout                     int         `json:"session-timeout"`
			MaxTimeLimit                       int         `json:"max-time-limit"`
			DefaultTimeLimit                   int         `json:"default-time-limit"`
			MaxInferenceSize                   int         `json:"max-inference-size"`
			DefaultInferenceSize               int         `json:"default-inference-size"`
			StaticExpires                      int         `json:"static-expires"`
			PreCommitTriggerDepth              int         `json:"pre-commit-trigger-depth"`
			PreCommitTriggerLimit              int         `json:"pre-commit-trigger-limit"`
			Collation                          string      `json:"collation"`
			CoordinateSystem                   string      `json:"coordinate-system"`
			InternalSecurity                   bool        `json:"internal-security"`
			ConcurrentRequestLimit             int         `json:"concurrent-request-limit"`
			ComputeContentLength               bool        `json:"compute-content-length"`
			FileLogLevel                       string      `json:"file-log-level"`
			LogErrors                          bool        `json:"log-errors"`
			DebugAllow                         bool        `json:"debug-allow"`
			ProfileAllow                       bool        `json:"profile-allow"`
			DefaultXqueryVersion               string      `json:"default-xquery-version"`
			MultiVersionConcurrencyControl     string      `json:"multi-version-concurrency-control"`
			DistributeTimestamps               string      `json:"distribute-timestamps"`
			OutputSgmlCharacterEntities        string      `json:"output-sgml-character-entities"`
			OutputEncoding                     string      `json:"output-encoding"`
			OutputMethod                       string      `json:"output-method"`
			OutputByteOrderMark                string      `json:"output-byte-order-mark"`
			OutputCdataSectionNamespaceURI     string      `json:"output-cdata-section-namespace-uri"`
			OutputCdataSectionLocalname        interface{} `json:"output-cdata-section-localname"`
			OutputDoctypePublic                string      `json:"output-doctype-public"`
			OutputDoctypeSystem                string      `json:"output-doctype-system"`
			OutputEscapeURIAttributes          string      `json:"output-escape-uri-attributes"`
			OutputIncludeContentType           string      `json:"output-include-content-type"`
			OutputIndent                       string      `json:"output-indent"`
			OutputIndentUntyped                string      `json:"output-indent-untyped"`
			OutputIndentTabs                   string      `json:"output-indent-tabs"`
			OutputMediaType                    string      `json:"output-media-type"`
			OutputNormalizationForm            string      `json:"output-normalization-form"`
			OutputOmitXMLDeclaration           string      `json:"output-omit-xml-declaration"`
			OutputStandalone                   string      `json:"output-standalone"`
			OutputUndeclarePrefixes            string      `json:"output-undeclare-prefixes"`
			OutputVersion                      string      `json:"output-version"`
			OutputIncludeDefaultAttributes     string      `json:"output-include-default-attributes"`
			DefaultErrorFormat                 string      `json:"default-error-format"`
			ErrorHandler                       string      `json:"error-handler"`
			Schemas                            interface{} `json:"schemas"`
			Namespaces                         interface{} `json:"namespaces"`
			ModuleLocations                    interface{} `json:"module-locations"`
			RequestBlackouts                   interface{} `json:"request-blackouts"`
			URLRewriter                        string      `json:"url-rewriter"`
			SslCertificateTemplate             string      `json:"ssl-certificate-template"`
			RewriteResolvesGlobally            bool        `json:"rewrite-resolves-globally"`
			SslAllowSslv3                      bool        `json:"ssl-allow-sslv3"`
			SslAllowTLS                        bool        `json:"ssl-allow-tls"`
			SslDisableSslv3                    bool        `json:"ssl-disable-sslv3"`
			SslDisableTlsv1                    bool        `json:"ssl-disable-tlsv1"`
			SslDisableTlsv11                   bool        `json:"ssl-disable-tlsv1-1"`
			SslDisableTlsv12                   bool        `json:"ssl-disable-tlsv1-2"`
			SslHostname                        string      `json:"ssl-hostname"`
			SslCiphers                         string      `json:"ssl-ciphers"`
			SslRequireClientCertificate        bool        `json:"ssl-require-client-certificate"`
			SslClientCertificateAuthoritiesPem []string    `json:"ssl-client-certificate-pem,omitempty"`
			ExternalSecurity                   []string    `json:"external-security,omitempty"`
		} `json:"package-http-properties"`
		Links struct {
			GroupName       string `json:"group-name"`
			Database        string `json:"database"`
			ModulesDatabase string `json:"modules-database"`
			DefaultUser     string `json:"default-user"`
		} `json:"links"`
	} `json:"config"`
}

type AppServerDefaultResponse struct {
	ServerDefault ServerDefault `json:"server-default"`
}

type ServerDefault struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	ServerKind   string       `json:"server-kind"`
	Meta         Meta         `json:"meta"`
	Relations    Relations    `json:"relations"`
	RelatedViews RelatedViews `json:"related-views"`
}

type AppServerStatusResponse struct {
	ServerStatus ServerStatus `json:"server-status"`
}

type ServerStatus struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	ServerKind       string    `json:"server-kind"`
	Meta             Meta      `json:"meta"`
	Relations        Relations `json:"relations"`
	StatusProperties struct {
		Enabled struct {
			Units string `json:"units"`
			Value bool   `json:"value"`
		} `json:"enabled"`
		Port             int    `json:"port"`
		Root             string `json:"root"`
		DisplayLastLogin struct {
			Units string `json:"units"`
			Value bool   `json:"value"`
		} `json:"display-last-login"`
		Backlog struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"backlog"`
		Threads struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"threads"`
		MaxThreads struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"max-threads"`
		RequestTimeout struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"request-timeout"`
		KeepAliveTimeout struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"keep-alive-timeout"`
		SessionTimeout struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"session-timeout"`
		StaticExpires struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"static-expires"`
		MaxTimeLimit struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"max-time-limit"`
		DefaultTimeLimit struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"default-time-limit"`
		MultiVersionConcurrencyControl struct {
			Units string `json:"units"`
			Value string `json:"value"`
		} `json:"multi-version-concurrency-control"`
		Authentication struct {
			Units string `json:"units"`
			Value string `json:"value"`
		} `json:"authentication"`
		DefaultUser            string `json:"default-user"`
		Privilege              int    `json:"privilege"`
		ConcurrentRequestLimit struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"concurrent-request-limit"`
		DefaultXqueryVersion struct {
			Units string `json:"units"`
			Value string `json:"value"`
		} `json:"default-xquery-version"`
		OutputSgmlCharacterEntities struct {
			Units string `json:"units"`
			Value string `json:"value"`
		} `json:"output-sgml-character-entities"`
		OutputEncoding   string `json:"output-encoding"`
		TotalRequestRate struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"total-request-rate"`
		TotalExpandedTreeCacheMissRate struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"total-expanded-tree-cache-miss-rate"`
		TotalExpandedTreeCacheHitRate struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"total-expanded-tree-cache-hit-rate"`
		TotalRequests struct {
			Units string `json:"units"`
			Value int    `json:"value"`
		} `json:"total-requests"`
		ContentDatabaseEnabled struct {
			Units string `json:"units"`
			Value bool   `json:"value"`
		} `json:"content-database-enabled"`
		HostDetail []struct {
			RelationID    string `json:"relation-id"`
			RequestsCount struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"requests-count"`
			MaxInferenceSize struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"max-inference-size"`
			DefaultInferenceSize struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"default-inference-size"`
			DistributeTimestamps struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"distribute-timestamps"`
			RequestRate struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"request-rate"`
			ExpandedTreeCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"expanded-tree-cache-hits"`
			ExpandedTreeCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"expanded-tree-cache-misses"`
			ExpandedTreeCacheHitRate struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"expanded-tree-cache-hit-rate"`
			ExpandedTreeCacheMissRate struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"expanded-tree-cache-miss-rate"`
			FsProgramCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-program-cache-hits"`
			FsProgramCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-program-cache-misses"`
			DbProgramCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-program-cache-hits"`
			DbProgramCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-program-cache-misses"`
			EnvProgramCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"env-program-cache-hits"`
			EnvProgramCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"env-program-cache-misses"`
			FsMainModuleSeqCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-main-module-seq-cache-hits"`
			FsMainModuleSeqCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-main-module-seq-cache-misses"`
			DbMainModuleSeqCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-main-module-seq-cache-hits"`
			DbMainModuleSeqCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-main-module-seq-cache-misses"`
			FsLibModuleCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-lib-module-cache-hits"`
			FsLibModuleCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"fs-lib-module-cache-misses"`
			DbLibModuleCacheHits struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-lib-module-cache-hits"`
			DbLibModuleCacheMisses struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"db-lib-module-cache-misses"`
			RequestTime struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"request-time"`
			ServerReceiveBytes struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-receive-bytes"`
			ServerReceiveTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"server-receive-time"`
			ServerReceiveRate struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-receive-rate"`
			ServerReceiveLoad struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-receive-load"`
			ServerSendBytes struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-send-bytes"`
			ServerSendTime struct {
				Units string `json:"units"`
				Value string `json:"value"`
			} `json:"server-send-time"`
			ServerSendRate struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-send-rate"`
			ServerSendLoad struct {
				Units string `json:"units"`
				Value int    `json:"value"`
			} `json:"server-send-load"`
		} `json:"host-detail"`
	} `json:"status-properties"`
	RelatedViews RelatedViews `json:"related-views"`
}

type AppServerPackageResponse struct {
	PackageHTTPServer PackageHTTPServer `json:"package-http-server"`
}

type AppServerProperties struct {
	ServerName                         string      `json:"server-name"`
	GroupName                          string      `json:"group-name"`
	ServerType                         string      `json:"server-type"`
	Enabled                            bool        `json:"enabled,omitempty"`
	Root                               string      `json:"root"`
	Authentication                     string      `json:"authentication,omitempty"`
	Port                               int         `json:"port"`
	WebDAV                             bool        `json:"webDAV,omitempty"`
	Execute                            bool        `json:"execute,omitempty"`
	DisplayLastLogin                   bool        `json:"display-last-login,omitempty"`
	Address                            string      `json:"address,omitempty"`
	Backlog                            int         `json:"backlog,omitempty"`
	Threads                            int         `json:"threads,omitempty"`
	RequestTimeout                     int         `json:"request-timeout,omitempty"`
	KeepAliveTimeout                   int         `json:"keep-alive-timeout,omitempty"`
	SessionTimeout                     int         `json:"session-timeout,omitempty"`
	MaxTimeLimit                       int         `json:"max-time-limit,omitempty"`
	DefaultTimeLimit                   int         `json:"default-time-limit,omitempty"`
	MaxInferenceSize                   int         `json:"max-inference-size,omitempty"`
	DefaultInferenceSize               int         `json:"default-inference-size,omitempty"`
	StaticExpires                      int         `json:"static-expires,omitempty"`
	PreCommitTriggerDepth              int         `json:"pre-commit-trigger-depth,omitempty"`
	PreCommitTriggerLimit              int         `json:"pre-commit-trigger-limit,omitempty"`
	Collation                          string      `json:"collation,omitempty"`
	CoordinateSystem                   string      `json:"coordinate-system,omitempty"`
	InternalSecurity                   bool        `json:"internal-security,omitempty"`
	ConcurrentRequestLimit             int         `json:"concurrent-request-limit,omitempty"`
	ComputeContentLength               bool        `json:"compute-content-length,omitempty"`
	FileLogLevel                       string      `json:"file-log-level,omitempty"`
	LogErrors                          bool        `json:"log-errors,omitempty"`
	DebugAllow                         bool        `json:"debug-allow,omitempty"`
	ProfileAllow                       bool        `json:"profile-allow,omitempty"`
	DefaultXqueryVersion               string      `json:"default-xquery-version,omitempty"`
	MultiVersionConcurrencyControl     string      `json:"multi-version-concurrency-control,omitempty"`
	DistributeTimestamps               string      `json:"distribute-timestamps,omitempty"`
	OutputSgmlCharacterEntities        string      `json:"output-sgml-character-entities,omitempty"`
	OutputEncoding                     string      `json:"output-encoding,omitempty"`
	OutputMethod                       string      `json:"output-method,omitempty"`
	OutputByteOrderMark                string      `json:"output-byte-order-mark,omitempty"`
	OutputCdataSectionNamespaceURI     string      `json:"output-cdata-section-namespace-uri,omitempty"`
	OutputCdataSectionLocalname        interface{} `json:"output-cdata-section-localname,omitempty"`
	OutputDoctypePublic                string      `json:"output-doctype-public,omitempty"`
	OutputDoctypeSystem                string      `json:"output-doctype-system,omitempty"`
	OutputEscapeURIAttributes          string      `json:"output-escape-uri-attributes,omitempty"`
	OutputIncludeContentType           string      `json:"output-include-content-type,omitempty"`
	OutputIndent                       string      `json:"output-indent,omitempty"`
	OutputIndentUntyped                string      `json:"output-indent-untyped,omitempty"`
	OutputIndentTabs                   string      `json:"output-indent-tabs,omitempty"`
	OutputMediaType                    string      `json:"output-media-type,omitempty"`
	OutputNormalizationForm            string      `json:"output-normalization-form,omitempty"`
	OutputOmitXMLDeclaration           string      `json:"output-omit-xml-declaration,omitempty"`
	OutputStandalone                   string      `json:"output-standalone,omitempty"`
	OutputUndeclarePrefixes            string      `json:"output-undeclare-prefixes,omitempty"`
	OutputVersion                      string      `json:"output-version,omitempty"`
	OutputIncludeDefaultAttributes     string      `json:"output-include-default-attributes,omitempty"`
	DefaultErrorFormat                 string      `json:"default-error-format,omitempty"`
	ErrorHandler                       string      `json:"error-handler,omitempty"`
	URLRewriter                        string      `json:"url-rewriter,omitempty"`
	RewriteResolvesGlobally            bool        `json:"rewrite-resolves-globally,omitempty"`
	SslCertificateTemplate             string      `json:"ssl-certificate-template"`
	SslAllowSslv3                      bool        `json:"ssl-allow-sslv3,omitempty"`
	SslAllowTLS                        bool        `json:"ssl-allow-tls,omitempty"`
	SslDisableSslv3                    bool        `json:"ssl-disable-sslv3,omitempty"`
	SslDisableTlsv1                    bool        `json:"ssl-disable-tlsv1,omitempty"`
	SslDisableTlsv11                   bool        `json:"ssl-disable-tlsv1-1,omitempty"`
	SslDisableTlsv12                   bool        `json:"ssl-disable-tlsv1-2,omitempty"`
	SslHostname                        string      `json:"ssl-hostname,omitempty"`
	SslCiphers                         string      `json:"ssl-ciphers,omitempty"`
	SslRequireClientCertificate        bool        `json:"ssl-require-client-certificate,omitempty"`
	SslClientCertificateAuthoritiesPEM []string    `json:"ssl-client-certificate-pem,omitempty"`
	ExternalSecurity                   []string    `json:"external-security,omitempty"`
	ContentDatabase                    string      `json:"content-database"`
	DefaultUser                        string      `json:"default-user,omitempty"`
}

type AppServerService struct {
	client Client
	base   string
}

func NewAppServerService(client Client, base string) *AppServerService {

	return &AppServerService{
		client: client,
		base:   base,
	}
}

//https://docs.marklogic.com/REST/GET/manage/v2/servers
func (s *AppServerService) GetAppServers(parms ServerParameters) (interface{}, RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	var appServerResponse interface{}
	if strings.EqualFold(parms.View, "Metrics") {
		appServerResponse = new(AppServersMetricResponse)
	} else if strings.EqualFold(parms.View, "Status") {
		appServerResponse = new(AppServersStatusResponse)
	} else if strings.EqualFold(parms.View, "Package") {
		appServerResponse = new(AppServersPackageResponse)
	} else { // Default
		appServerResponse = new(AppServersDefaultResponse)
	}
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp
}

func (s *AppServerService) GetAppServer(appServer string, parms ServerParameters) (interface{}, RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"/"+appServer+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	var appServerResponse interface{}
	if strings.EqualFold(parms.View, "Status") {
		appServerResponse = new(AppServerStatusResponse)
	} else if strings.EqualFold(parms.View, "Package") {
		appServerResponse = new(AppServerPackageResponse)
	} else { // Default
		appServerResponse = new(AppServerDefaultResponse)
	}
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp

}

func (s *AppServerService) GetAppServerProperties(appServer string, parms ServerParameters) (interface{}, RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("GET", s.base+AppServers+"/"+appServer+"/properties?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	appServerResponse := new(AppServerProperties)

	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, appServerResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return appServerResponse, *errorResponse, *resp

}

func (s *AppServerService) AddAppServer(appServer AppServerProperties) (RestErrorResponse, http.Response) {
	body, err := json.Marshal(appServer)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", s.base+AppServers, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, nil, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *errorResponse, *resp
}

func (s *AppServerService) UpdateAppServer(appServer AppServerProperties) (RestartResponse, RestErrorResponse, http.Response) {

	body, err := json.Marshal(appServer)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest("PUT", s.base+AppServers+"/"+appServer.ServerName+"/properties?group-id="+appServer.GroupName, bytes.NewBuffer(body))

	s.client = Decorate(s.client,
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
	)
	errorResponse := new(RestErrorResponse)
	restartResponse := new(RestartResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp
}

func (s *AppServerService) DeleteAppServer(appServer string, parms ServerParameters) (RestartResponse, RestErrorResponse, http.Response) {
	v, err := query.Values(parms)
	if err != nil {
		log.Fatalln(err)
	}
	urlParms := v.Encode()

	req, _ := http.NewRequest("DELETE", s.base+AppServers+"/"+appServer+"?"+urlParms, nil)

	if strings.EqualFold(parms.Format, "xml") {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/xml"),
		)
	} else {
		s.client = Decorate(s.client,
			AddHeader("Accept", "application/json"),
		)
	}

	restartResponse := new(RestartResponse)
	errorResponse := new(RestErrorResponse)
	resp, err := ExecuteRequest(s.client, req, restartResponse, errorResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return *restartResponse, *errorResponse, *resp

}
