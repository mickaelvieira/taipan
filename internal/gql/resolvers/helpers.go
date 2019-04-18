package resolvers

import "log"

// GetBoundariesFromArgs prepare the default offset and limit for the SQL query
// provide a default limit value and get back a closure to prepare the boundaries
// Example:
// 		fromArgs := GetBoundariesFromArgs(10)
// 		offset, limit := fromArgs(args.Offset, args.Limit)
func GetBoundariesFromArgs(defLimit int32) func(*int32, *int32) (int32, int32) {
	if defLimit <= 0 {
		log.Fatal("the default limit must be greater than zero")
	}

	return func(o *int32, l *int32) (int32, int32) {
		var offset int32
		if o != nil {
			offset = *o
		}

		if offset < 0 {
			offset = 0
		}

		var limit int32
		if l != nil {
			limit = *l
		}

		if limit <= 0 {
			limit = defLimit
		}

		return offset, limit
	}
}
