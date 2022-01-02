package tags

/* Why use macros when you have Emacs?

;; I don't think I could have used `stringer` bacause of the case conversions.

(save-match-data
  (search-forward-regexp "const (\n[ \t]*\\([a-zA-Z]+ +\\)Tag = iota\\([^)]+\\))")
  (let ((tags (split-string
               (concat (match-string-no-properties 1)
                       (match-string-no-properties 2))))
        (lower (lambda (s)
                 (replace-regexp-in-string "\\(musicbrainz\\)\\(.*\\)"
                                           "\\1_\\2"
                                           (downcase s)))))
    (insert
     (concat "\n\nvar FromString = map[string]Tag{\n"
             (string-join
              (mapcar (lambda (s)
                        (format "\"%s\": %s,\n" (funcall lower s) s))
                      tags))
             "}\n\nvar ToString = map[Tag]string{\n"
             (string-join
              (mapcar (lambda (s)
                        (format "%s: \"%s\",\n" s (funcall lower s)))
                      tags))
             "}"))
    (gofmt)))
*/

type Tag int

const (
	Artist Tag = iota
	ArtistSort
	Album
	AlbumSort
	AlbumArtist
	AlbumArtistSort
	Title
	Track
	Name
	Genre
	Date
	OriginalDate
	Composer
	ComposerSort
	Performer
	Conductor
	Work
	Ensemble
	Movement
	MovementNumber
	Location
	Grouping
	Comment
	Disc
	Label
	MusicBrainzArtistId
	MusicBrainzAlbumId
	MusicBrainzAlbumArtistId
	MusicBrainzTrackId
	MusicBrainzReleaseTrackId
	MusicBrainzWorkId
)

var FromString = map[string]Tag{
	"artist":                     Artist,
	"artistsort":                 ArtistSort,
	"album":                      Album,
	"albumsort":                  AlbumSort,
	"albumartist":                AlbumArtist,
	"albumartistsort":            AlbumArtistSort,
	"title":                      Title,
	"track":                      Track,
	"name":                       Name,
	"genre":                      Genre,
	"date":                       Date,
	"originaldate":               OriginalDate,
	"composer":                   Composer,
	"composersort":               ComposerSort,
	"performer":                  Performer,
	"conductor":                  Conductor,
	"work":                       Work,
	"ensemble":                   Ensemble,
	"movement":                   Movement,
	"movementnumber":             MovementNumber,
	"location":                   Location,
	"grouping":                   Grouping,
	"comment":                    Comment,
	"disc":                       Disc,
	"label":                      Label,
	"musicbrainz_artistid":       MusicBrainzArtistId,
	"musicbrainz_albumid":        MusicBrainzAlbumId,
	"musicbrainz_albumartistid":  MusicBrainzAlbumArtistId,
	"musicbrainz_trackid":        MusicBrainzTrackId,
	"musicbrainz_releasetrackid": MusicBrainzReleaseTrackId,
	"musicbrainz_workid":         MusicBrainzWorkId,
}

var ToString = map[Tag]string{
	Artist:                    "artist",
	ArtistSort:                "artistsort",
	Album:                     "album",
	AlbumSort:                 "albumsort",
	AlbumArtist:               "albumartist",
	AlbumArtistSort:           "albumartistsort",
	Title:                     "title",
	Track:                     "track",
	Name:                      "name",
	Genre:                     "genre",
	Date:                      "date",
	OriginalDate:              "originaldate",
	Composer:                  "composer",
	ComposerSort:              "composersort",
	Performer:                 "performer",
	Conductor:                 "conductor",
	Work:                      "work",
	Ensemble:                  "ensemble",
	Movement:                  "movement",
	MovementNumber:            "movementnumber",
	Location:                  "location",
	Grouping:                  "grouping",
	Comment:                   "comment",
	Disc:                      "disc",
	Label:                     "label",
	MusicBrainzArtistId:       "musicbrainz_artistid",
	MusicBrainzAlbumId:        "musicbrainz_albumid",
	MusicBrainzAlbumArtistId:  "musicbrainz_albumartistid",
	MusicBrainzTrackId:        "musicbrainz_trackid",
	MusicBrainzReleaseTrackId: "musicbrainz_releasetrackid",
	MusicBrainzWorkId:         "musicbrainz_workid",
}

func (tag Tag) String() string {
	return ToString[tag]
}
