package owntracks

const (
	TypeLocation = "location"
)

type Location struct {
	// Message type
	Type string `json:"_type"`

	// [meters]
	Accuracy int `json:"acc"`

	// [meters]
	Altitude int `json:"alt"`

	// [percent]
	Battery int `json:"batt"`

	// Device heading
	// [degrees]
	Course int `json:"cog"`

	// [degrees]
	Latitude float64 `json:"lat"`

	// [degrees]
	Longitude float64 `json:"lon"`

	// Radius of the region when entering/leaving
	// [meters]
	Radius int `json:"rad"` // [meters]

	// Report trigger
	// p (random ping) | c (enter/leave circle) | b (enter/leave beacon) |
	// r (response to report command) | u (manual) | t (timer) | v (os settings change)
	Trigger string `json:"t"`

	// User tracker initials
	TrackerID string `json:"tid"`

	// Unix time of the event as reported by OwnTracks
	// [seconds]
	ReportedTimestamp int `json:"tst"`

	// Unix time of the events reception and processing by the Lambda
	ReceivedTimestamp int64

	// [meters]
	VerticalAccuracy int `json:"vac"`

	// [kilometers / hour]
	Velocity int `json:"vel"`

	// [kilopascals]
	Pressure float64 `json:"p"`

	// Data connectivity status during event
	// w (wifi) | o (offline) | m (mobile)
	ConnectivityStatus string `json:"conn"`

	// MQTT topic
	Topic string `json:"topic"`

	// List of regions event is currently within
	InRegions []string `json:"inregions"`
}
