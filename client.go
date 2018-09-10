package delay109

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	endpoint = "https://www.tokyu.co.jp/railway/delay/print.php"
)

// Client 遅延情報取得クライアント
type Client struct {
	httpClient *http.Client
}

// Query 検索クエリ
type Query struct {
	Line  LineID
	Year  int
	Month int
	Day   int

	// DirectionUp | DirectionDown
	Direction int

	// Before10am | After10am
	TimePeriod int
}

// New Clientコンストラクタ
func New(h *http.Client) *Client {
	if h == nil {
		// TODO: デフォルトの設定を追加する
		h = &http.Client{}
	}
	return &Client{httpClient: h}
}

// Search 遅延情報を検索する
func (c *Client) Search(q *Query) (*LineStatus, error) {
	URL, _ := queryToURL(q)
	req, _ := http.NewRequest("GET", URL.String(), nil)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// EUC-JP => UTF-8
	in := bytes.NewBuffer(nil)
	out := bytes.NewBuffer(nil)
	reader := transform.NewReader(in, japanese.EUCJP.NewDecoder())
	res.Write(in)
	if _, err = io.Copy(out, reader); err != nil {
		return nil, err
	}

	// 遅延時間のパース
	htmlBytes := out.Bytes()
	delayDuration, err := parseDelayDuration(htmlBytes)
	if err != nil {
		return nil, err
	}

	status := createLineStatus(q.Line, q.Direction, time.Date(q.Year, time.Month(q.Month), q.Day, 0, 0, 0, 0, time.Local), delayDuration)
	return status, nil
}

// queryToURL Queryからパース対象ページのURLを生成する
func queryToURL(q *Query) (*url.URL, error) {
	URL, _ := url.Parse(endpoint)
	qstring := URL.Query()
	qstring.Set("line", strconv.Itoa(int(q.Line)+q.Direction))
	qstring.Set("d", fmt.Sprintf("%d%02d%02d_%d", q.Year, q.Month, q.Day, q.TimePeriod))
	URL.RawQuery = qstring.Encode()

	return URL, nil
}
