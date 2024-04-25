/**
  @author: decision
  @date: 2024/4/8
  @note:
**/

package bark

type MsgLevel string

const (
	MsgLevelActive        MsgLevel = "active"
	MsgLevelTimeSensitive MsgLevel = "timeSensitive"
	MsgLevelPassive       MsgLevel = "passive"
)

type RequestBody struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Level     MsgLevel `json:"level,omitempty"`
	Badge     int      `json:"badge,omitempty"`
	AutoCopy  string   `json:"autoCopy,omitempty"`
	Copy      string   `json:"copy,omitempty"`
	Icon      string   `json:"icon,omitempty"`
	Group     string   `json:"group,omitempty"`
	IsArchive int      `json:"isArchive,omitempty"`
	Url       string   `json:"url,omitempty"`
	// Sound?
}

type RequestBuilder struct {
	body *RequestBody
}

func NewRequestBuilder() *RequestBuilder {
	body := RequestBody{}
	return &RequestBuilder{
		body: &body,
	}
}

func (b *RequestBuilder) SetTitle(title string) *RequestBuilder {
	b.body.Title = title
	return b
}

func (b *RequestBuilder) SetBody(body string) *RequestBuilder {
	b.body.Body = body
	return b
}

func (b *RequestBuilder) SetLevel(level MsgLevel) *RequestBuilder {
	b.body.Level = level
	return b
}

func (b *RequestBuilder) SetBadge(badge int) *RequestBuilder {
	b.body.Badge = badge
	return b
}

func (b *RequestBuilder) SetAutoCopy(autoCopy string) *RequestBuilder {
	b.body.AutoCopy = autoCopy
	return b
}

func (b *RequestBuilder) SetCopy(copy string) *RequestBuilder {
	b.body.Copy = copy
	return b
}

func (b *RequestBuilder) SetIcon(icon string) *RequestBuilder {
	b.body.Icon = icon
	return b
}

func (b *RequestBuilder) SetGroup(group string) *RequestBuilder {
	b.body.Group = group
	return b
}

func (b *RequestBuilder) SetArchive() *RequestBuilder {
	b.body.IsArchive = 1
	return b
}

func (b *RequestBuilder) SetUrl(url string) *RequestBuilder {
	b.body.Url = url
	return b
}

func (b *RequestBuilder) Result() *RequestBody {
	return b.body
}
