package gql

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
	gpx "github.com/ptrv/go-gpx"
)

var loader = GetActivityLoader()

// RunResolver a run info
type RunResolver struct {
	ID       graphql.ID
	Name     string
	Datetime string
}

// Records list the records
func (r *RunResolver) Records(ctx context.Context) (*[]*RecordResolver, error) {
	var records []*RecordResolver

	thunk := loader.Load(ctx, dataloader.StringKey(r.ID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	data, ok := result.(*gpx.Gpx)

	if !ok {
		return nil, errors.New("Wrong data")
	}

	if len(data.Tracks) > 0 {
		var track = data.Tracks[0]
		if len(track.Segments) > 0 {
			var segment = track.Segments[0]
			for _, wp := range segment.Waypoints {
				bytes := []byte(fmt.Sprintf("%f%f%s", wp.Lat, wp.Lon, wp.Timestamp))
				id := fmt.Sprintf("%x", md5.Sum(bytes))

				var record = RecordResolver{ID: graphql.ID(id), Lat: wp.Lat, Lon: wp.Lon, Datetime: wp.Timestamp}
				records = append(records, &record)
			}
		}
	}

	return &records, nil
}

// RecordResolver a run info
type RecordResolver struct {
	ID        graphql.ID
	Lat       float64
	Lon       float64
	Datetime  string
	Heartrate int32
}

// Resolvers resolvers
type Resolvers struct{}

// ReadRun returns a run info
func (r *Resolvers) ReadRun(ctx context.Context, args struct{ ID string }) (*RunResolver, error) {

	thunk := loader.Load(ctx, dataloader.StringKey(args.ID))
	result, err := thunk()

	if err != nil {
		return nil, err
	}

	data, ok := result.(*gpx.Gpx)

	if !ok {
		return nil, errors.New("Wrong data")
	}

	if len(data.Tracks) > 0 {
		var track = data.Tracks[0]
		return &RunResolver{ID: graphql.ID(args.ID), Name: track.Name, Datetime: data.Metadata.Timestamp}, nil
	}

	return &RunResolver{}, nil
}
