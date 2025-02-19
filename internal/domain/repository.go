package domain

import "context"

type SignalStore interface {
	CreateSignal(ctx context.Context, signal *Signal) error
	GetSignal(ctx context.Context, signalID int) (*Signal, error)
	ListSignals(ctx context.Context, limit, page int) (signals []Signal, count int, err error)
	UpdateSignal(ctx context.Context, signal *Signal) error
	DeleteSignal(ctx context.Context, signalID int) error
}

type TrackStore interface {
	CreateTrack(ctx context.Context, track *Track) error
	GetTrack(ctx context.Context, trackID int) (*Track, error)
	ListTracks(ctx context.Context, limit, page int) (tracks []Track, count int, err error)
	UpdateTrack(ctx context.Context, track *Track) error
	DeleteTrack(ctx context.Context, trackID int) error

	ListSignalTracks(ctx context.Context, signalID, limit, page int) (tracks []Track, count int, err error)
}

type MileageStore interface {
	AddMileage(ctx context.Context, mileage *Mileage) error
}
