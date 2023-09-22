package excavator

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"golang.org/x/net/html"
)


type Excavator struct{
    sl *slog.Logger
}

func NewExcavator(sl *slog.Logger) *Excavator{
    return &Excavator{
        sl : sl,
    }

}

func(ex *Excavator)MakeRequest(url string) (io.ReadCloser, error){
    resp, err := http.DefaultClient.Get(url)
    if err != nil{
        //ex.sl.Info("Cannot ger respond", err)
        return nil, err
    }
    return resp.Body, nil
}

func(ex *Excavator)StartExcavating() error{ 
    rb, err := ex.MakeRequest("https://en.wikipedia.org/wiki/Lists_of_songs")
    if err != nil{
        //ex.sl.Info("Cannot Get Response from the Provided url",err)
        return err
    }
    defer rb.Close()
    doc, err := html.Parse(rb) 
    if err != nil {
        //ex.sl.Info("We can't start execavating ",err)
        return err
    }

    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "div"{
            outer:
            for _,v := range n.Attr{
                if v.Val == "div-col"{
                    fmt.Println("Key:", v.Key,"Value:", v.Val )
                    if  n.FirstChild.NextSibling.Data == "ul"{
                        
                    }
                }
                break outer
            }
        }
        for c := n.FirstChild; c != nil; c=c.NextSibling{
            f(c)
        }
    }
    
    f(doc)

    return nil
}


func(ex *Excavator) excavateLink(n *html.Node) {}








