package types

/*
{
            "name": "spotify",
            "revision": {
                "version": "1.1.55.498.gf9a83c60"
            },
            "snap": {
                "description": "Love music? Play your favorite songs and albums free on Linux with Spotify.\r\n\r\nStream the tracks you love instantly, browse the charts or fire up readymade playlists in every genre and mood. Radio plays you great song after great song, based on your music taste. Discover new music too, with awesome playlists built just for you.\r\n\r\nStream Spotify free, with occasional ads, or go Premium.\r\n\r\nFree:\r\n\r\n• Play any song, artist, album or playlist instantly\r\n\r\n• Browse hundreds of readymade playlists in every genre and mood\r\n\r\n• Stay on top of the Charts\r\n\r\n• Stream Radio \r\n\r\n• Enjoy podcasts, audiobooks and videos\r\n\r\n• Discover more music with personalized playlists\r\n\r\nPremium:\r\n\r\n• Download tunes and play offline\r\n\r\n• Listen ad-free\r\n\r\n• Get even better sound quality\r\n\r\n• Try it free for 30 days, no strings attached\r\n\r\nLike us on Facebook: http://www.facebook.com/spotify \r\nFollow us on Twitter: http://twitter.com/spotify\r\n\r\nNote: Spotify for Linux is a labor of love from our engineers that wanted to listen to Spotify on their Linux development machines. They work on it in their spare time and it is currently not a platform that we actively support. The experience may differ from our other Spotify Desktop clients, such as Windows and Mac.",
                "media": [
                    {
                        "height": 256,
                        "type": "icon",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/spotify-linux-256.png",
                        "width": 256
                    },
                    {
                        "height": 787,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/Screenshot_from_2017-12-18_12-07-06.png",
                        "width": 1399
                    },
                    {
                        "height": 787,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/Screenshot_from_2017-12-18_12-09-22.png",
                        "width": 1399
                    },
                    {
                        "height": 787,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/Screenshot_from_2017-12-18_12-18-27.png",
                        "width": 1399
                    },
                    {
                        "height": 787,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/Screenshot_from_2017-12-18_12-20-23.png",
                        "width": 1399
                    },
                    {
                        "height": 240,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/banner_dSwF9EF.png",
                        "width": 1218
                    },
                    {
                        "height": 200,
                        "type": "screenshot",
                        "url": "https://dashboard.snapcraft.io/site_media/appmedia/2017/12/banner-icon_WaLCF17.png",
                        "width": 387
                    }
                ],
                "publisher": {
                    "display-name": "Spotify",
                    "id": "Z7emHVerQzTDQrOcaLJ8wO1ir1NPl3fG",
                    "username": "spotify",
                    "validation": "verified"
                },
                "title": "Spotify"
            },
            "snap-id": "pOBIoZ2LrCB3rDohMxoYGnbN14EHOgD7"
        }
*/

type Revision struct {
	Version string `json:"version"`
}

type SnapAppMedia string

var (
	TypeIcon       SnapAppMedia = "icon"
	TypeScreenshot SnapAppMedia = "screenshot"
	IsVerified                  = "verified"
)

type Media struct {
	Height int          `json:"height"`
	Type   SnapAppMedia `json:"type"`
	URL    string       `json:"url"`
	Width  int          `json:"width"`
}

type SnapPublisher struct {
	DisplayName string `json:"display-name"`
	Id          string `json:"id"`
	Username    string `json:"username"`
	Validation  string `json:"validation"`
}

type SApp struct {
	Description string `json:"description"`
	Media       []Media

	Publisher SnapPublisher `json:"publisher"`
	Title     string        `json:"title"`
}

type Snap struct {
	Name string   `json:"name"`
	Rev  Revision `json:"revision"`

	App SApp `json:"snap"`

	SnapID string `json:"snap-id"`
}

type SnapResults struct {
	Results []Snap `json:"results"`
}

func (s *Snap) Export() ExportedApp {
	return ExportedApp{
		AppType:        SnapApp,
		Name:           s.App.Title,
		Description:    s.App.Title,
		Version:        s.Rev.Version,
		StatsAvailable: false,
		Stats:          ExportedStats{},
		IconURL:        s.Icon(),
		HeaderURL:      "",
		Screenshots:    s.Screenshots(),
		Publisher:      s.ExportPublisher(),
		PackageName:    s.Name,
		Package:        Package{},
	}
}

func (s *Snap) IsPaid() bool {
	return false
}

func (s *Snap) ExportPublisher() ExportedPublisher {
	return ExportedPublisher{}
}

func (s *Snap) Screenshots() []string {
	var res []string

	for _, x := range s.App.Media {
		if x.Type == TypeScreenshot {
			res = append(res, x.URL)
		}
	}

	return res
}

func (s *Snap) Icon() string {
	var res string

	for _, x := range s.App.Media {
		if x.Type == TypeIcon {
			res = x.URL
			break
		}
	}

	return res
}
