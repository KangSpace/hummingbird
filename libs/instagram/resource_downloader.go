//Package instagram Download Instagram Shared Image/Videos
package instagram

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hummingbird/libs/proxyhttp"
	"hummingbird/libs/trycatch"
	"hummingbird/libs/webserver"
	"io"
	"net/http"
	"strings"
)

//is or not show console out
var isShowLog = false

//InsResourceDownloader downloader
type InsResourceDownloader struct {
	URL string
	//Proxy Server e.g.: http://127.0.0.1:1880
	Proxy string
}

//InsGetHttpHandle http handler for http server
func (downloader *InsResourceDownloader) InsGetHttpHandle(w http.ResponseWriter, req *http.Request) {
	webserver.SetAccessControlAllowOrgin(w)
	// instagram shared url
	url := req.FormValue("url")
	if len(strings.Trim(url, " ")) < 1 || (strings.Index(url, "http") != 0 && strings.Index(url, "https") != 0) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		str := "<pre>Usage:<p style='color:red'>" + req.Host + "/ins-get?url=https://www.instagram.com/p/xxxxx/</span></p>"
		w.Write([]byte(str))
		return
	}
	trycatch.Trycatch(func() {
		w.Header().Set("Content-Type", "text/plain")
		downloader.URL = url
		data, err := downloader.FetchInsShareDataResources()
		if err == nil {
			dataStr, _ := data.ToJSONString()
			w.Write([]byte(dataStr))
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, err)
		}
	}, func(e interface{}) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error:" + url))
	})
}

// script tag profix
var shareDataScriptProfix = "<script type=\"text/javascript\">window._sharedData = "
var scriptSuffix = ";</script>"

//FetchInsShareDataResources fetch instagram share data resources
func (downloader *InsResourceDownloader) FetchInsShareDataResources() (*FetchedResource, error) {
	//resp, err := http.Get(url)
	resp, err := (&(proxyhttp.ProxyHttp{downloader.Proxy})).Client().Get(downloader.URL)
	if err != nil {
		return &FetchedResource{}, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	readReader := bufio.NewReader(resp.Body)
	var dataStr string
	for {
		//line, err2 := readReader.ReadBytes('\n')
		line, err2 := readReader.ReadString('\n')
		line = strings.TrimLeft(line, " ")
		if err2 != nil || io.EOF == err2 {
			break
		}
		if strings.Index(line, shareDataScriptProfix) != -1 {
			dataStr = line[len(shareDataScriptProfix) : len(line)-len(scriptSuffix)-1]
		}
	}
	if len(dataStr) == 0 {
		return nil, nil
	}
	if isShowLog {
		fmt.Println("line:", string(dataStr))
	}
	var v ShareData
	err1 := json.Unmarshal([]byte(dataStr), &v)
	if err1 != nil {
		fmt.Println("err:", err1)
	}
	var displayNodes []DisplayNode
	//a,_:=json.Marshal(v.EntryData.PostPage[0].Graphql.ShortcodeMedia)
	//fmt.Println("al",string(a))
	shortcodeMedia := v.EntryData.PostPage[0].Graphql.ShortcodeMedia
	shortCode := v.EntryData.PostPage[0].Graphql.ShortcodeMedia.ShortCode
	//image
	if !shortcodeMedia.IsVideo {
		if len(shortcodeMedia.EdgeSidecarToChildren.Edges) > 0 {
			for _, l := range shortcodeMedia.EdgeSidecarToChildren.Edges {
				//fmt.Println("l",l)
				var tempDisplayResources []DisplayResource
				for _, rr := range l.Node.DisplayResource {
					tempDisplayResources = append(tempDisplayResources, DisplayResource{rr.ConfigHeight, rr.ConfigWidth, rr.Src})
				}
				displayNodes = append(displayNodes, DisplayNode{DisplayResources: tempDisplayResources,
					VideoProperties: VideoProperties{IsVideo: l.Node.IsVideo, VideoURL: l.Node.VideoURL}})
			}
		} else {
			var tempDisplayResources []DisplayResource
			for _, r := range shortcodeMedia.DisplayResources {
				tempDisplayResources = append(tempDisplayResources, DisplayResource{r.ConfigHeight, r.ConfigWidth, r.Src})
			}
			displayNodes = append(displayNodes, DisplayNode{DisplayResources: tempDisplayResources,
				VideoProperties: VideoProperties{IsVideo: false}})
		}
		return &FetchedResource{IsVideo: false, Images: displayNodes, NodeProperties: NodeProperties{ShortCode: shortCode}}, nil
	}
	//video
	return &FetchedResource{IsVideo: true, Video: VideoDisplayResource{ThumbnailSrc: shortcodeMedia.ThumbnailSrc,
		VideoDuration: shortcodeMedia.VideoDuration, VideoURL: shortcodeMedia.VideoURL}, NodeProperties: NodeProperties{ShortCode: shortCode}}, nil

	//json.NewDecoder().Decode(v)
	//return &displayResources,nil
}

//FetchedResource :fetched resource object
type FetchedResource struct {
	IsVideo bool                 `json:"is_video"`
	Images  []DisplayNode        `json:"images"`
	Video   VideoDisplayResource `json:"video"`
	NodeProperties
}

//ToJSONString get FetchedResource json string
func (resource *FetchedResource) ToJSONString() (string, error) {
	bytes, err := json.Marshal(resource)
	return string(bytes), err
}

//VideoDisplayResource :video resource
type VideoDisplayResource struct {
	VideoURL      string  `json:"video_url"`
	VideoDuration float32 `json:"video_duration"`
	ThumbnailSrc  string  `json:"thumbnail_src"`
	NodeProperties
}
type DisplayNode struct {
	DisplayResources []DisplayResource `json:"display_resources"`
	VideoProperties
}

//DisplayResource :image resource
type DisplayResource struct {
	ConfigHeight int    `json:"config_height"`
	ConfigWidth  int    `json:"config_width"`
	Src          string `json:"src"`
}

//ShareData shareData
type ShareData struct {
	EntryData EntryData `json:"entry_data"`
	//EntryData interface{} `json:"entry_data"`
}

//EntryData entryData
type EntryData struct {
	PostPage []PostPage `json:"PostPage"`
	//EntryData interface{} `json:"entry_data"`
}

//PostPage postPage
type PostPage struct {
	Graphql Graphql `json:"graphql"`
	//Graphql interface{} `json:"graphql"`
}

//Graphql graphql
type Graphql struct {
	ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
}

//ShortcodeMedia shortcodeMedia
type ShortcodeMedia struct {
	DisplayResources      []DisplayResource     `json:"display_resources"`
	EdgeSidecarToChildren EdgeSidecarToChildren `json:"edge_sidecar_to_children"`
	VideoProperties
}

//EdgeSidecarToChildren edgeSidecarToChildren
type EdgeSidecarToChildren struct {
	Edges []Edges `json:"edges"`
}

//Edges edges
type Edges struct {
	Node Node `json:"node"`
}

//Node node
type Node struct {
	DisplayResource []DisplayResource `json:"display_resources"`
	VideoProperties
}

type NodeProperties struct {
	//Dimensions string `json:"dimensions"`
	ShortCode string `json:"shortcode"`
	TypeName  string `json:"__typename"` //"GraphVideo","GraphImage"
}

// video 属性
type VideoProperties struct {
	IsVideo       bool    `json:"is_video"`
	VideoURL      string  `json:"video_url"`
	ThumbnailSrc  string  `json:"thumbnail_src"`
	VideoDuration float32 `json:"video_duration"`
	NodeProperties
}

const (
	NodeTypeName_GraphVideo = "GraphVideo"
	NodeTypeName_GraphImage = "GraphImage"
)

//Main Test main
func Main() {
	//single image
	//var insShareUrl = "https://www.instagram.com/p/Brx04JWHqfS/"
	// multi-image
	var insShareURL = "https://www.instagram.com/p/BrJhZU4AsEZ/?utm_source=ig_web_copy_link"
	//video
	//var insShareUrl = "https://www.instagram.com/p/BrOTfIeg_1l/?utm_source=ig_web_copy_link"
	var insDownloader = &InsResourceDownloader{URL: insShareURL}
	data, err := insDownloader.FetchInsShareDataResources()
	if err != nil {
		fmt.Println(err)
	}
	returnJSON, _ := json.Marshal(data)
	fmt.Println("data:", string(returnJSON))
	//download file on server

}
