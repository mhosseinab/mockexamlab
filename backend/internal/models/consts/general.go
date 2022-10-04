package consts

const (
	SEARCH   = 1  // Filter response with LIKE query "search={search_phrase}"
	FILTER   = 2  // Filter response by column name values "filter={column_name}:{value}"
	PAGINATE = 4  // Paginate response with page and page_size
	ORDER_BY = 8  // Order response by column name
	ALL      = 15 // Equivalent to SEARCH|FILTER|PAGINATE|ORDER_BY
	TagKey   = "filter"
)
